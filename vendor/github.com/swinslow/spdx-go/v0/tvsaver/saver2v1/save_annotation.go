// SPDX-License-Identifier: Apache-2.0 OR GPL-2.0-or-later

package saver2v1

import (
	"fmt"
	"io"

	"github.com/swinslow/spdx-go/v0/spdx"
)

func renderAnnotation2_1(ann *spdx.Annotation2_1, w io.Writer) error {
	if ann.Annotator != "" && ann.AnnotatorType != "" {
		fmt.Fprintf(w, "Annotator: %s: %s\n", ann.AnnotatorType, ann.Annotator)
	}
	if ann.AnnotationDate != "" {
		fmt.Fprintf(w, "AnnotationDate: %s\n", ann.AnnotationDate)
	}
	if ann.AnnotationType != "" {
		fmt.Fprintf(w, "AnnotationType: %s\n", ann.AnnotationType)
	}
	if ann.AnnotationSPDXIdentifier != "" {
		fmt.Fprintf(w, "SPDXREF: %s\n", ann.AnnotationSPDXIdentifier)
	}
	if ann.AnnotationComment != "" {
		fmt.Fprintf(w, "AnnotationComment: %s\n", textify(ann.AnnotationComment))
	}

	return nil
}
