// Package saver2v1 contains functions to render and write a tag-value
// formatted version of an in-memory SPDX document and its sections
// (version 2.1).
// SPDX-License-Identifier: Apache-2.0 OR GPL-2.0-or-later
package saver2v1

import (
	"fmt"
	"io"

	"github.com/swinslow/spdx-go/v0/spdx"
)

// RenderDocument2_1 is the main entry point to take an SPDX in-memory
// Document (version 2.1), and render it to the received io.Writer.
// It is only exported in order to be available to the tvsaver package,
// and typically does not need to be called by client code.
func RenderDocument2_1(doc *spdx.Document2_1, w io.Writer) error {
	if doc.CreationInfo == nil {
		return fmt.Errorf("Document had nil CreationInfo section")
	}

	renderCreationInfo2_1(doc.CreationInfo, w)

	for _, pkg := range doc.Packages {
		if pkg.IsUnpackaged == true {
			fmt.Fprintf(w, "##### Unpackaged files\n\n")
		} else {
			fmt.Fprintf(w, "##### Package: %s\n\n", pkg.PackageName)
		}
		renderPackage2_1(pkg, w)
	}

	if len(doc.OtherLicenses) > 0 {
		fmt.Fprintf(w, "##### Other Licenses\n\n")
		for _, ol := range doc.OtherLicenses {
			renderOtherLicense2_1(ol, w)
		}
	}

	if len(doc.Relationships) > 0 {
		fmt.Fprintf(w, "##### Relationships\n\n")
		for _, rln := range doc.Relationships {
			renderRelationship2_1(rln, w)
		}
		fmt.Fprintf(w, "\n")
	}

	if len(doc.Annotations) > 0 {
		fmt.Fprintf(w, "##### Annotations\n\n")
		for _, ann := range doc.Annotations {
			renderAnnotation2_1(ann, w)
			fmt.Fprintf(w, "\n")
		}
	}

	if len(doc.Reviews) > 0 {
		fmt.Fprintf(w, "##### Reviews\n\n")
		for _, rev := range doc.Reviews {
			renderReview2_1(rev, w)
		}
	}

	return nil
}
