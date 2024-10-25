// Code generated by assertions generator; DO NOT EDIT.

package resourceassert

import (
	"testing"

	"github.com/Snowflake-Labs/terraform-provider-snowflake/pkg/acceptance/bettertestspoc/assert"
)

type StreamOnDirectoryTableResourceAssert struct {
	*assert.ResourceAssert
}

func StreamOnDirectoryTableResource(t *testing.T, name string) *StreamOnDirectoryTableResourceAssert {
	t.Helper()

	return &StreamOnDirectoryTableResourceAssert{
		ResourceAssert: assert.NewResourceAssert(name, "resource"),
	}
}

func ImportedStreamOnDirectoryTableResource(t *testing.T, id string) *StreamOnDirectoryTableResourceAssert {
	t.Helper()

	return &StreamOnDirectoryTableResourceAssert{
		ResourceAssert: assert.NewImportedResourceAssert(id, "imported resource"),
	}
}

///////////////////////////////////
// Attribute value string checks //
///////////////////////////////////

func (s *StreamOnDirectoryTableResourceAssert) HasCommentString(expected string) *StreamOnDirectoryTableResourceAssert {
	s.AddAssertion(assert.ValueSet("comment", expected))
	return s
}

func (s *StreamOnDirectoryTableResourceAssert) HasCopyGrantsString(expected string) *StreamOnDirectoryTableResourceAssert {
	s.AddAssertion(assert.ValueSet("copy_grants", expected))
	return s
}

func (s *StreamOnDirectoryTableResourceAssert) HasDatabaseString(expected string) *StreamOnDirectoryTableResourceAssert {
	s.AddAssertion(assert.ValueSet("database", expected))
	return s
}

func (s *StreamOnDirectoryTableResourceAssert) HasFullyQualifiedNameString(expected string) *StreamOnDirectoryTableResourceAssert {
	s.AddAssertion(assert.ValueSet("fully_qualified_name", expected))
	return s
}

func (s *StreamOnDirectoryTableResourceAssert) HasNameString(expected string) *StreamOnDirectoryTableResourceAssert {
	s.AddAssertion(assert.ValueSet("name", expected))
	return s
}

func (s *StreamOnDirectoryTableResourceAssert) HasSchemaString(expected string) *StreamOnDirectoryTableResourceAssert {
	s.AddAssertion(assert.ValueSet("schema", expected))
	return s
}

func (s *StreamOnDirectoryTableResourceAssert) HasStageString(expected string) *StreamOnDirectoryTableResourceAssert {
	s.AddAssertion(assert.ValueSet("stage", expected))
	return s
}

func (s *StreamOnDirectoryTableResourceAssert) HasStaleString(expected string) *StreamOnDirectoryTableResourceAssert {
	s.AddAssertion(assert.ValueSet("stale", expected))
	return s
}

////////////////////////////
// Attribute empty checks //
////////////////////////////

func (s *StreamOnDirectoryTableResourceAssert) HasNoComment() *StreamOnDirectoryTableResourceAssert {
	s.AddAssertion(assert.ValueNotSet("comment"))
	return s
}

func (s *StreamOnDirectoryTableResourceAssert) HasNoCopyGrants() *StreamOnDirectoryTableResourceAssert {
	s.AddAssertion(assert.ValueNotSet("copy_grants"))
	return s
}

func (s *StreamOnDirectoryTableResourceAssert) HasNoDatabase() *StreamOnDirectoryTableResourceAssert {
	s.AddAssertion(assert.ValueNotSet("database"))
	return s
}

func (s *StreamOnDirectoryTableResourceAssert) HasNoFullyQualifiedName() *StreamOnDirectoryTableResourceAssert {
	s.AddAssertion(assert.ValueNotSet("fully_qualified_name"))
	return s
}

func (s *StreamOnDirectoryTableResourceAssert) HasNoName() *StreamOnDirectoryTableResourceAssert {
	s.AddAssertion(assert.ValueNotSet("name"))
	return s
}

func (s *StreamOnDirectoryTableResourceAssert) HasNoSchema() *StreamOnDirectoryTableResourceAssert {
	s.AddAssertion(assert.ValueNotSet("schema"))
	return s
}

func (s *StreamOnDirectoryTableResourceAssert) HasNoStage() *StreamOnDirectoryTableResourceAssert {
	s.AddAssertion(assert.ValueNotSet("stage"))
	return s
}

func (s *StreamOnDirectoryTableResourceAssert) HasNoStale() *StreamOnDirectoryTableResourceAssert {
	s.AddAssertion(assert.ValueNotSet("stale"))
	return s
}