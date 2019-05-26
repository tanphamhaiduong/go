package modules

import (
	"github.com/tanphamhaiduong/go/delta/generator/gotype"
	"github.com/tanphamhaiduong/go/delta/generator/graphqltype"
	"github.com/tanphamhaiduong/go/delta/generator/querytype"
)

func init() {
	modules = append(modules, companyModule)
}

var (
	companyModule = Module{
		Name: "Company",
		Fields: []Field{
			{
				Name:               "ID",
				GoType:             gotype.Int64,
				QueryType:          querytype.Equal,
				GraphQLType:        graphqltype.Int,
				GraphQLDescription: "This is company id",
				GraphQLRequired:    true,
				Validate:           `validate:"required,min=1"`,
			},
			{
				Name:               "Name",
				GoType:             gotype.String,
				QueryType:          querytype.Like,
				GraphQLType:        graphqltype.String,
				GraphQLDescription: "This is company name",
				GraphQLRequired:    true,
				Validate:           `validate:"omitempty"`,
			},
			{
				Name:               "Status",
				GoType:             gotype.String,
				QueryType:          querytype.Equal,
				GraphQLType:        graphqltype.String,
				GraphQLDescription: "This is company active",
				GraphQLRequired:    true,
				Validate:           `validate:"omitempty,oneof=active inactive"`,
			},
			{
				Name:               "CreatedBy",
				GoType:             gotype.String,
				QueryType:          querytype.Equal,
				GraphQLType:        graphqltype.String,
				GraphQLDescription: "This is company createdBy",
				GraphQLRequired:    false,
				Validate:           `validate:"omitempty"`,
			},
			{
				Name:               "UpdatedBy",
				GoType:             gotype.String,
				QueryType:          querytype.Equal,
				GraphQLType:        graphqltype.String,
				GraphQLDescription: "This is company updatedBy",
				GraphQLRequired:    false,
				Validate:           `validate:"omitempty"`,
			},
		},
	}
)
