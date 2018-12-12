// SPDX-License-Identifier: Apache-2.0 OR GPL-2.0-or-later

package saver2v1

import (
	"fmt"
	"io"

	"github.com/swinslow/spdx-go/v0/spdx"
)

func renderSnippet2_1(sn *spdx.Snippet2_1, w io.Writer) error {
	if sn.SnippetSPDXIdentifier != "" {
		fmt.Fprintf(w, "SnippetSPDXIdentifier: %s\n", sn.SnippetSPDXIdentifier)
	}
	if sn.SnippetFromFileSPDXIdentifier != "" {
		fmt.Fprintf(w, "SnippetFromFileSPDXID: %s\n", sn.SnippetFromFileSPDXIdentifier)
	}
	if sn.SnippetByteRangeStart != 0 && sn.SnippetByteRangeEnd != 0 {
		fmt.Fprintf(w, "SnippetByteRange: %d:%d\n", sn.SnippetByteRangeStart, sn.SnippetByteRangeEnd)
	}
	if sn.SnippetLineRangeStart != 0 && sn.SnippetLineRangeEnd != 0 {
		fmt.Fprintf(w, "SnippetLineRange: %d:%d\n", sn.SnippetLineRangeStart, sn.SnippetLineRangeEnd)
	}
	if sn.SnippetLicenseConcluded != "" {
		fmt.Fprintf(w, "SnippetLicenseConcluded: %s\n", sn.SnippetLicenseConcluded)
	}
	for _, s := range sn.LicenseInfoInSnippet {
		fmt.Fprintf(w, "LicenseInfoInSnippet: %s\n", s)
	}
	if sn.SnippetLicenseComments != "" {
		fmt.Fprintf(w, "SnippetLicenseComments: %s\n", textify(sn.SnippetLicenseComments))
	}
	if sn.SnippetCopyrightText != "" {
		fmt.Fprintf(w, "SnippetCopyrightText: %s\n", sn.SnippetCopyrightText)
	}
	if sn.SnippetComment != "" {
		fmt.Fprintf(w, "SnippetComment: %s\n", textify(sn.SnippetComment))
	}
	if sn.SnippetName != "" {
		fmt.Fprintf(w, "SnippetName: %s\n", sn.SnippetName)
	}

	fmt.Fprintf(w, "\n")

	return nil
}
