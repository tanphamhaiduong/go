package modules

import (
	"github.com/tanphamhaiduong/go/delta/pkg/generator/internal/gotype"
	"github.com/tanphamhaiduong/go/delta/pkg/generator/internal/graphqltype"
	"github.com/tanphamhaiduong/go/delta/pkg/generator/internal/querytype"
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
				GraphQLDescription: "This is role's id",
				GraphQLRequired:    true,
				Validate:           `validate:"required,min=1"`,
				FakerTag:           `faker:"unix_time"`,
			},
			{
				Name:               "Name",
				GoType:             gotype.String,
				QueryType:          querytype.Equal,
				GraphQLType:        graphqltype.String,
				GraphQLDescription: "This is role's name",
				GraphQLRequired:    true,
				Validate:           `validate:"omitempty"`,
				FakerTag:           `faker:"name"`,
			},
			{
				Name:               "CompanyID",
				GoType:             gotype.Int64,
				QueryType:          querytype.Equal,
				GraphQLType:        graphqltype.Int,
				GraphQLDescription: "This is role's companyId",
				GraphQLRequired:    true,
				Validate:           `validate:"omitempty,min=1"`,
				FakerTag:           `faker:"unix_time"`,
			},
			{
				Name:               "Status",
				GoType:             gotype.String,
				QueryType:          querytype.Equal,
				GraphQLType:        graphqltype.String,
				GraphQLDescription: "This is role's status",
				GraphQLRequired:    true,
				Validate:           `validate:"omitempty,oneof=active inactive"`,
				FakerTag:           `faker:"word"`,
			},
			{
				Name:               "CreatedBy",
				GoType:             gotype.String,
				QueryType:          querytype.Equal,
				GraphQLType:        graphqltype.String,
				GraphQLDescription: "This is role's createdBy",
				GraphQLRequired:    false,
				Validate:           `validate:"omitempty"`,
				FakerTag:           `faker:"email"`,
			},
			{
				Name:               "UpdatedBy",
				GoType:             gotype.String,
				QueryType:          querytype.Equal,
				GraphQLType:        graphqltype.String,
				GraphQLDescription: "This is role's updatedBy",
				GraphQLRequired:    false,
				Validate:           `validate:"omitempty"`,
				FakerTag:           `faker:"email"`,
			},
		},
	}
)
