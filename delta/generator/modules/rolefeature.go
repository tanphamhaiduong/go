package modules

import (
	"github.com/tanphamhaiduong/go/delta/generator/gotype"
	"github.com/tanphamhaiduong/go/delta/generator/graphqltype"
	"github.com/tanphamhaiduong/go/delta/generator/querytype"
)

func init() {
	modules = append(modules, roleFeatureModule)
}

var (
	roleFeatureModule = Module{
		Name: "RoleFeature",
		Fields: []Field{
			{
				Name:               "ID",
				GoType:             gotype.Int64,
				QueryType:          querytype.Equal,
				GraphQLType:        graphqltype.Int,
				GraphQLDescription: "This is roleFeature id",
				GraphQLRequired:    true,
				Validate:           `validate:"required,min=1"`,
			},
			{
				Name:               "RoleID",
				GoType:             gotype.Int64,
				QueryType:          querytype.Equal,
				GraphQLType:        graphqltype.Int,
				GraphQLDescription: "This is roleFeature id",
				GraphQLRequired:    true,
				Validate:           `validate:"required,min=1"`,
			},
			{
				Name:               "FeatureID",
				GoType:             gotype.Int64,
				QueryType:          querytype.Equal,
				GraphQLType:        graphqltype.Int,
				GraphQLDescription: "This is roleFeature id",
				GraphQLRequired:    true,
				Validate:           `validate:"required,min=1"`,
			},
			{
				Name:               "CreatedBy",
				GoType:             gotype.String,
				QueryType:          querytype.Equal,
				GraphQLType:        graphqltype.String,
				GraphQLDescription: "This is roleFeature createdBy",
				GraphQLRequired:    false,
				Validate:           `validate:"omitempty`,
			},
			{
				Name:               "UpdatedBy",
				GoType:             gotype.String,
				QueryType:          querytype.Equal,
				GraphQLType:        graphqltype.String,
				GraphQLDescription: "This is roleFeature updatedBy",
				GraphQLRequired:    false,
				Validate:           `validate:"omitempty`,
			},
		},
	}
)
