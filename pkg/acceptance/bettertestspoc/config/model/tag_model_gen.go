// Code generated by config model builder generator; DO NOT EDIT.

package model

import (
	tfconfig "github.com/hashicorp/terraform-plugin-testing/config"

	"github.com/Snowflake-Labs/terraform-provider-snowflake/pkg/acceptance/bettertestspoc/config"
	"github.com/Snowflake-Labs/terraform-provider-snowflake/pkg/provider/resources"
)

type TagModel struct {
	AllowedValues      tfconfig.Variable `json:"allowed_values,omitempty"`
	Comment            tfconfig.Variable `json:"comment,omitempty"`
	Database           tfconfig.Variable `json:"database,omitempty"`
	FullyQualifiedName tfconfig.Variable `json:"fully_qualified_name,omitempty"`
	MaskingPolicies    tfconfig.Variable `json:"masking_policies,omitempty"`
	Name               tfconfig.Variable `json:"name,omitempty"`
	Schema             tfconfig.Variable `json:"schema,omitempty"`

	*config.ResourceModelMeta
}

/////////////////////////////////////////////////
// Basic builders (resource name and required) //
/////////////////////////////////////////////////

func Tag(
	resourceName string,
	database string,
	name string,
	schema string,
) *TagModel {
	t := &TagModel{ResourceModelMeta: config.Meta(resourceName, resources.Tag)}
	t.WithDatabase(database)
	t.WithName(name)
	t.WithSchema(schema)
	return t
}

func TagWithDefaultMeta(
	database string,
	name string,
	schema string,
) *TagModel {
	t := &TagModel{ResourceModelMeta: config.DefaultMeta(resources.Tag)}
	t.WithDatabase(database)
	t.WithName(name)
	t.WithSchema(schema)
	return t
}

/////////////////////////////////
// below all the proper values //
/////////////////////////////////

// allowed_values attribute type is not yet supported, so WithAllowedValues can't be generated

func (t *TagModel) WithComment(comment string) *TagModel {
	t.Comment = tfconfig.StringVariable(comment)
	return t
}

func (t *TagModel) WithDatabase(database string) *TagModel {
	t.Database = tfconfig.StringVariable(database)
	return t
}

func (t *TagModel) WithFullyQualifiedName(fullyQualifiedName string) *TagModel {
	t.FullyQualifiedName = tfconfig.StringVariable(fullyQualifiedName)
	return t
}

// masking_policy attribute type is not yet supported, so WithMaskingPolicy can't be generated

func (t *TagModel) WithName(name string) *TagModel {
	t.Name = tfconfig.StringVariable(name)
	return t
}

func (t *TagModel) WithSchema(schema string) *TagModel {
	t.Schema = tfconfig.StringVariable(schema)
	return t
}

//////////////////////////////////////////
// below it's possible to set any value //
//////////////////////////////////////////

func (t *TagModel) WithAllowedValuesValue(value tfconfig.Variable) *TagModel {
	t.AllowedValues = value
	return t
}

func (t *TagModel) WithCommentValue(value tfconfig.Variable) *TagModel {
	t.Comment = value
	return t
}

func (t *TagModel) WithDatabaseValue(value tfconfig.Variable) *TagModel {
	t.Database = value
	return t
}

func (t *TagModel) WithFullyQualifiedNameValue(value tfconfig.Variable) *TagModel {
	t.FullyQualifiedName = value
	return t
}

func (t *TagModel) WithMaskingPoliciesValue(value tfconfig.Variable) *TagModel {
	t.MaskingPolicies = value
	return t
}

func (t *TagModel) WithNameValue(value tfconfig.Variable) *TagModel {
	t.Name = value
	return t
}

func (t *TagModel) WithSchemaValue(value tfconfig.Variable) *TagModel {
	t.Schema = value
	return t
}