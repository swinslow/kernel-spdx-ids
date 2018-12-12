// SPDX-License-Identifier: Apache-2.0 OR GPL-2.0-or-later

package spdx

// CreationInfo2_1 is a Document Creation Information section of an
// SPDX Document for version 2.1 of the spec.
type CreationInfo2_1 struct {

	// 2.1: SPDX Version; should be in the format "SPDX-2.1"
	// Cardinality: mandatory, one
	SPDXVersion string

	// 2.2: Data License; should be "CC0-1.0"
	// Cardinality: mandatory, one
	DataLicense string

	// 2.3: SPDX Identifier; should be "SPDXRef-DOCUMENT"
	// Cardinality: mandatory, one
	SPDXIdentifier string

	// 2.4: Document Name
	// Cardinality: mandatory, one
	DocumentName string

	// 2.5: Document Namespace
	// Cardinality: mandatory, one
	DocumentNamespace string

	// 2.6: External Document References
	// Cardinality: optional, one or many
	ExternalDocumentReferences []string

	// 2.7: License List Version
	// Cardinality: optional, one
	LicenseListVersion string

	// 2.8: Creators: may have multiple keys for Person, Organization
	//      and/or Tool
	// Cardinality: mandatory, one or many
	CreatorPersons       []string
	CreatorOrganizations []string
	CreatorTools         []string

	// 2.9: Created: data format YYYY-MM-DDThh:mm:ssZ
	// Cardinality: mandatory, one
	Created string

	// 2.10: Creator Comment
	// Cardinality: optional, one
	CreatorComment string

	// 2.11: Document Comment
	// Cardinality: optional, one
	DocumentComment string
}
