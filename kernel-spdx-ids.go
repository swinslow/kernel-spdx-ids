// SPDX-License-Identifier: Apache-2.0 OR GPL-2.0-or-later

package main

import (
	"flag"
	"fmt"
	"io"
	"net/http"
	"os"
	"path/filepath"

	"github.com/mholt/archiver"
	"github.com/swinslow/spdx-go/v0/idsearcher"
	"github.com/swinslow/spdx-go/v0/reporter"
	"github.com/swinslow/spdx-go/v0/tvsaver"
)

func main() {

	// set up command-line flags

	// what kernel version are we looking at?
	// required flag if retrieving kernel, and/or saving an SPDX file.
	versionPtr := flag.String("v", "", "kernel version to scan")

	// path for scanning. required.
	// if -d is included, the kernel will be unpacked within this directory.
	// if -d is not included, we assume the kernel has already been
	// unpacked here.
	pathPtr := flag.String("p", "", "directory path to scan kernel (and unpack it, if also downloading)")

	// want to download the kernel source tarball?
	downloadPtr := flag.Bool("d", false, "download this kernel version from kernel.org")

	// path to save SPDX file to disk. If empty string, will not save.
	spdxPathPtr := flag.String("s", "", "path to output SPDX file, if desired")

	// parse flags and check that we've received a valid combo of arguments
	flag.Parse()

	if *downloadPtr && *versionPtr == "" {
		fmt.Println("Must include flag for kernel version if downloading kernel (e.g., -v=4.19)")
		return
	}
	if *spdxPathPtr != "" && *versionPtr == "" {
		fmt.Println("Must include flag for kernel version if saving SPDX file (e.g., -v=4.19)")
		return
	}
	if *pathPtr == "" {
		if *downloadPtr {
			fmt.Println("Must include flag for path to unpack downloaded kernel (e.g., -p=/tmp/...)")
			return
		}
		fmt.Println("Must include flag for path to search unpacked kernel (e.g., -p=/tmp/...)")
		return
	}

	// now, download and unpack the kernel if asked to
	var dstPath string
	if *downloadPtr {
		err := getKernelTarball(*versionPtr, *pathPtr)
		if err != nil {
			fmt.Println(err)
			return
		}

		tarballPath := filepath.Join(*pathPtr, getKernelFilename(*versionPtr))
		dstPath = filepath.Join(*pathPtr, getKernelPackageName(*versionPtr))
		// create dst unpacking directory if it doesn't already exist
		err = os.MkdirAll(dstPath, 0755)
		if err != nil {
			fmt.Printf("Couldn't create directory %s for unpacking kernel: %v\n", dstPath, err)
			return
		}
		// actually unpack it
		err = archiver.Unarchive(tarballPath, dstPath)
		if err != nil {
			fmt.Printf("Couldn't unpack kernel at %s into %s: %v\n", tarballPath, dstPath, err)
			return
		}
	} else {
		// not unpacking; set dstPath appropriately for scanning
		dstPath = *pathPtr
	}

	// next, define the idsearcher configuration struct
	// NOTE that the user will likely want to modify the PackageNamespace
	// in the generated SPDX file, if they will be saving and reusing it
	config := &idsearcher.Config{

		NamespacePrefix:     "https://github.com/swinslow/kernel-spdx-ids/idsearcher-results-for-",
		BuilderPathsIgnored: []string{},
		SearcherPathsIgnored: []string{
			// ignore the documentation file which explains how to use SPDX
			// short-form IDs (and therefore has a bunch of "SPDX-License-Identifier:"
			// tags that we wouldn't want to pick up).
			"/Documentation/process/license-rules.rst",

			// also ignore all files in the /LICENSES/ directory.
			"/LICENSES/",
		},
	}

	packageName := getKernelPackageName(*versionPtr)
	doc, err := idsearcher.BuildIDsDocument(packageName, dstPath, config)
	if err != nil {
		fmt.Printf("Error while building document and searching for IDs: %v\n", err)
		return
	}

	// and add kernel-spdx-ids as another Creator tool
	doc.CreationInfo.CreatorTools = append(doc.CreationInfo.CreatorTools, "github.com/swinslow/kernel-spdx-ids")

	fmt.Printf("Successfully searched for IDs for kernel version %s\n", *versionPtr)

	// NOTE that BuildIDsDocument does NOT do any validation of the license
	// identifiers, to confirm that they are e.g. on the SPDX License List
	// or in other appropriate format (e.g., LicenseRef-...)

	// generate and print a report
	// pick the first package because that's the only one we've got
	pkg := doc.Packages[0]
	err = reporter.Generate(pkg, os.Stdout)
	if err != nil {
		fmt.Printf("Error while generating report: %v\n", err)
	}

	// we can now save the SPDX document to disk, using tvsaver, if requested.

	if *spdxPathPtr != "" {
		// create a new file for writing
		w, err := os.Create(*spdxPathPtr)
		if err != nil {
			fmt.Printf("Error while opening %v for writing: %v\n", *spdxPathPtr, err)
			return
		}
		defer w.Close()

		err = tvsaver.Save2_1(doc, w)
		if err != nil {
			fmt.Printf("Error while saving %v: %v", *spdxPathPtr, err)
			return
		}

		fmt.Printf("Successfully saved %v\n", *spdxPathPtr)
	}
}

func getKernelPackageName(ver string) string {
	return fmt.Sprintf("linux-%s", ver)
}

func getKernelFilename(ver string) string {
	return fmt.Sprintf("linux-%s.tar.xz", ver)
}

func getKernelTarball(ver string, dst string) error {
	filename := getKernelFilename(ver)
	dstpath := filepath.Join(dst, filename)
	kernelURL := fmt.Sprintf("https://cdn.kernel.org/pub/linux/kernel/v4.x/%s", filename)

	fmt.Printf("Attempting to download kernel version %s\n", ver)
	fmt.Printf("  from %s...\n", kernelURL)

	f, err := os.Create(dstpath)
	if err != nil {
		return fmt.Errorf("Couldn't create file for writing at %s: %v", dstpath, err)
	}
	defer f.Close()

	r, err := http.Get(kernelURL)
	if err != nil {
		return fmt.Errorf("Couldn't download from %s: %v", kernelURL, err)
	}
	defer r.Body.Close()

	_, err = io.Copy(f, r.Body)
	if err != nil {
		return fmt.Errorf("Couldn't save download contents: %v", err)
	}

	return nil
}
