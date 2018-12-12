// SPDX-License-Identifier: Apache-2.0 OR GPL-2.0-or-later

package saver2v1

import (
	"fmt"
	"io"

	"github.com/swinslow/spdx-go/v0/spdx"
)

func renderPackage2_1(pkg *spdx.Package2_1, w io.Writer) error {
	if pkg.IsUnpackaged == false {
		if pkg.PackageName != "" {
			fmt.Fprintf(w, "PackageName: %s\n", pkg.PackageName)
		}
		if pkg.PackageSPDXIdentifier != "" {
			fmt.Fprintf(w, "SPDXID: %s\n", pkg.PackageSPDXIdentifier)
		}
		if pkg.PackageVersion != "" {
			fmt.Fprintf(w, "PackageVersion: %s\n", pkg.PackageVersion)
		}
		if pkg.PackageFileName != "" {
			fmt.Fprintf(w, "PackageFileName: %s\n", pkg.PackageFileName)
		}
		if pkg.PackageSupplierPerson != "" {
			fmt.Fprintf(w, "PackageSupplier: Person: %s\n", pkg.PackageSupplierPerson)
		}
		if pkg.PackageSupplierOrganization != "" {
			fmt.Fprintf(w, "PackageSupplier: Organization: %s\n", pkg.PackageSupplierOrganization)
		}
		if pkg.PackageSupplierNOASSERTION == true {
			fmt.Fprintf(w, "PackageSupplier: NOASSERTION\n")
		}
		if pkg.PackageOriginatorPerson != "" {
			fmt.Fprintf(w, "PackageOriginator: Person: %s\n", pkg.PackageOriginatorPerson)
		}
		if pkg.PackageOriginatorOrganization != "" {
			fmt.Fprintf(w, "PackageOriginator: Organization: %s\n", pkg.PackageOriginatorOrganization)
		}
		if pkg.PackageOriginatorNOASSERTION == true {
			fmt.Fprintf(w, "PackageOriginator: NOASSERTION\n")
		}
		if pkg.PackageDownloadLocation != "" {
			fmt.Fprintf(w, "PackageDownloadLocation: %s\n", pkg.PackageDownloadLocation)
		}
		if pkg.FilesAnalyzed == true {
			if pkg.IsFilesAnalyzedTagPresent == true {
				fmt.Fprintf(w, "FilesAnalyzed: true\n")
			}
		} else {
			fmt.Fprintf(w, "FilesAnalyzed: false\n")
		}
		if pkg.PackageVerificationCode != "" && pkg.FilesAnalyzed == true {
			if pkg.PackageVerificationCodeExcludedFile == "" {
				fmt.Fprintf(w, "PackageVerificationCode: %s\n", pkg.PackageVerificationCode)
			} else {
				fmt.Fprintf(w, "PackageVerificationCode: %s (excludes %s)\n", pkg.PackageVerificationCode, pkg.PackageVerificationCodeExcludedFile)
			}
		}
		if pkg.PackageChecksumSHA1 != "" {
			fmt.Fprintf(w, "PackageChecksum: SHA1: %s\n", pkg.PackageChecksumSHA1)
		}
		if pkg.PackageChecksumSHA256 != "" {
			fmt.Fprintf(w, "PackageChecksum: SHA256: %s\n", pkg.PackageChecksumSHA256)
		}
		if pkg.PackageChecksumMD5 != "" {
			fmt.Fprintf(w, "PackageChecksum: MD5: %s\n", pkg.PackageChecksumMD5)
		}
		if pkg.PackageHomePage != "" {
			fmt.Fprintf(w, "PackageHomePage: %s\n", pkg.PackageHomePage)
		}
		if pkg.PackageSourceInfo != "" {
			fmt.Fprintf(w, "PackageSourceInfo: %s\n", textify(pkg.PackageSourceInfo))
		}
		if pkg.PackageLicenseConcluded != "" {
			fmt.Fprintf(w, "PackageLicenseConcluded: %s\n", pkg.PackageLicenseConcluded)
		}
		if pkg.FilesAnalyzed == true {
			for _, s := range pkg.PackageLicenseInfoFromFiles {
				fmt.Fprintf(w, "PackageLicenseInfoFromFiles: %s\n", s)
			}
		}
		if pkg.PackageLicenseDeclared != "" {
			fmt.Fprintf(w, "PackageLicenseDeclared: %s\n", pkg.PackageLicenseDeclared)
		}
		if pkg.PackageLicenseComments != "" {
			fmt.Fprintf(w, "PackageLicenseComments: %s\n", textify(pkg.PackageLicenseComments))
		}
		if pkg.PackageCopyrightText != "" {
			fmt.Fprintf(w, "PackageCopyrightText: %s\n", pkg.PackageCopyrightText)
		}
		if pkg.PackageSummary != "" {
			fmt.Fprintf(w, "PackageSummary: %s\n", textify(pkg.PackageSummary))
		}
		if pkg.PackageDescription != "" {
			fmt.Fprintf(w, "PackageDescription: %s\n", textify(pkg.PackageDescription))
		}
		if pkg.PackageComment != "" {
			fmt.Fprintf(w, "PackageComment: %s\n", textify(pkg.PackageComment))
		}
		for _, s := range pkg.PackageExternalReferences {
			fmt.Fprintf(w, "ExternalRef: %s %s %s\n", s.Category, s.RefType, s.Locator)
			if s.ExternalRefComment != "" {
				fmt.Fprintf(w, "ExternalRefComment: %s\n", s.ExternalRefComment)
			}
		}

		fmt.Fprintf(w, "\n")
	}

	// also render any files for this package, even if unpackaged
	for _, f := range pkg.Files {
		renderFile2_1(f, w)
	}

	return nil
}
