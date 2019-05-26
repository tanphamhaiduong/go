package modules

import (
	"github.com/tanphamhaiduong/go/delta/generator/gotype"
	"github.com/tanphamhaiduong/go/delta/generator/graphqltype"
	"github.com/tanphamhaiduong/go/delta/generator/querytype"
)

func init() {
	modules = append(modules, roleModule)
}

var (
	roleModule = Module{
		Name: "Role",
		Fields: []Field{
			{
				Name:               "ID",
				GoType:             gotype.Int64,
				QueryType:          querytype.Equal,
				GraphQLType:        graphqltype.Int,
				GraphQLDescription: "This is role id",
				GraphQLRequired:    true,
				Validate:           `validate:"required,min=1"`,
			},
			{
				Name:               "Name",
				GoType:             gotype.String,
				QueryType:          querytype.Equal,
				GraphQLType:        graphqltype.String,
				GraphQLDescription: "This is role url",
				GraphQLRequired:    true,
				Validate:           `validate:"omitempty"`,
			},
			{
				Name:               "CompanyID",
				GoType:             gotype.Int64,
				QueryType:          querytype.Equal,
				GraphQLType:        graphqltype.Int,
				GraphQLDescription: "This is role companyId",
				GraphQLRequired:    true,
				Validate:           `validate:"omitempty,min=1"`,
			},
			{
				Name:               "Status",
				GoType:             gotype.String,
				QueryType:          querytype.Equal,
				GraphQLType:        graphqltype.String,
				GraphQLDescription: "This is role active",
				GraphQLRequired:    true,
				Validate:           `validate:"omitempty,oneof=active inactive`,
			},
			{
				Name:               "CreatedBy",
				GoType:             gotype.String,
				QueryType:          querytype.Equal,
				GraphQLType:        graphqltype.String,
				GraphQLDescription: "This is role createdBy",
				GraphQLRequired:    false,
				Validate:           `validate:"omitempty`,
			},
			{
				Name:               "UpdatedBy",
				GoType:             gotype.String,
				QueryType:          querytype.Equal,
				GraphQLType:        graphqltype.String,
				GraphQLDescription: "This is role updatedBy",
				GraphQLRequired:    false,
				Validate:           `validate:"omitempty`,
			},
		},
	}
)
