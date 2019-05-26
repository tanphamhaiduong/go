package modules

import (
	"github.com/tanphamhaiduong/go/delta/generator/gotype"
	"github.com/tanphamhaiduong/go/delta/generator/graphqltype"
	"github.com/tanphamhaiduong/go/delta/generator/querytype"
)

func init() {
	modules = append(modules, featureModule)
}

var (
	featureModule = Module{
		Name: "Feature",
		Fields: []Field{
			{
				Name:               "ID",
				GoType:             gotype.Int64,
				QueryType:          querytype.Equal,
				GraphQLType:        graphqltype.Int,
				GraphQLDescription: "This is feature id",
				GraphQLRequired:    true,
				Validate:           `validate:"required,min=1"`,
			},
			{
				Name:               "URL",
				GoType:             gotype.String,
				QueryType:          querytype.Equal,
				GraphQLType:        graphqltype.String,
				GraphQLDescription: "This is feature url",
				GraphQLRequired:    true,
				Validate:           `validate:"omitempty"`,
			},
			{
				Name:               "Meta",
				GoType:             gotype.String,
				QueryType:          querytype.Equal,
				GraphQLType:        graphqltype.String,
				GraphQLDescription: "This is feature meta",
				GraphQLRequired:    true,
				Validate:           `validate:"omitempty"`,
			},
			{
				Name:               "CompanyID",
				GoType:             gotype.Int64,
				QueryType:          querytype.Equal,
				GraphQLType:        graphqltype.Int,
				GraphQLDescription: "This is feature companyId",
				GraphQLRequired:    true,
				Validate:           `validate:"omitempty,min=1"`,
			},
			{
				Name:               "Status",
				GoType:             gotype.String,
				QueryType:          querytype.Equal,
				GraphQLType:        graphqltype.String,
				GraphQLDescription: "This is feature active",
				GraphQLRequired:    true,
				Validate:           `validate:"omitempty,oneof=active inactive"`,
			},
			{
				Name:               "CreatedBy",
				GoType:             gotype.String,
				QueryType:          querytype.Equal,
				GraphQLType:        graphqltype.String,
				GraphQLDescription: "This is feature createdBy",
				GraphQLRequired:    false,
				Validate:           `validate:"omitempty"`,
			},
			{
				Name:               "UpdatedBy",
				GoType:             gotype.String,
				QueryType:          querytype.Equal,
				GraphQLType:        graphqltype.String,
				GraphQLDescription: "This is feature updatedBy",
				GraphQLRequired:    false,
				Validate:           `validate:"omitempty"`,
			},
		},
	}
)
