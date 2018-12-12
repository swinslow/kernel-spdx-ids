// SPDX-License-Identifier: Apache-2.0 OR GPL-2.0-or-later

package builder2v1

import (
	"fmt"
	"time"

	"github.com/swinslow/spdx-go/v0/spdx"
)

// BuildCreationInfoSection2_1 creates an SPDX Package (version 2.1), returning that
// package or error if any is encountered. Arguments:
//   - packageName: name of package / directory
//   - code: verification code from Package
//   - namespacePrefix: prefix for DocumentNamespace (packageName and code will be added)
//   - creatorType: one of Person, Organization or Tool
//   - creator: creator string
//   - testValues: for testing only; call with nil when using in production
func BuildCreationInfoSection2_1(packageName string, code string, namespacePrefix string, creatorType string, creator string, testValues map[string]string) (*spdx.CreationInfo2_1, error) {
	// build creator slices
	cPersons := []string{}
	cOrganizations := []string{}
	cTools := []string{}
	// add builder as a tool
	cTools = append(cTools, "github.com/swinslow/spdx-go/v0/builder")

	switch creatorType {
	case "Person":
		cPersons = append(cPersons, creator)
	case "Organization":
		cOrganizations = append(cOrganizations, creator)
	case "Tool":
		cTools = append(cTools, creator)
	default:
		cPersons = append(cPersons, creator)
	}

	// use test Created time if passing test values
	created := time.Now().Format("2006-01-02T15:04:05Z")
	if testVal := testValues["Created"]; testVal != "" {
		created = testVal
	}

	ci := &spdx.CreationInfo2_1{
		SPDXVersion:          "SPDX-2.1",
		DataLicense:          "CC0-1.0",
		SPDXIdentifier:       "SPDXRef-DOCUMENT",
		DocumentName:         packageName,
		DocumentNamespace:    fmt.Sprintf("%s%s-%s", namespacePrefix, packageName, code),
		CreatorPersons:       cPersons,
		CreatorOrganizations: cOrganizations,
		CreatorTools:         cTools,
		Created:              created,
	}
	return ci, nil
}
