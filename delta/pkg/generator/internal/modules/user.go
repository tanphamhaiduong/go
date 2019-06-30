package modules

import (
	"github.com/tanphamhaiduong/go/delta/pkg/generator/internal/gotype"
	"github.com/tanphamhaiduong/go/delta/pkg/generator/internal/graphqltype"
	"github.com/tanphamhaiduong/go/delta/pkg/generator/internal/querytype"
)

func init() {
	modules = append(modules, userModule)
}

var (
	userModule = Module{
		Name: "User",
		Fields: []Field{
			{
				Name:               "ID",
				GoType:             gotype.Int64,
				QueryType:          querytype.Equal,
				GraphQLType:        graphqltype.Int,
				GraphQLDescription: "This is user's id",
				GraphQLRequired:    true,
				Validate:           `validate:"required,min=1"`,
				FakerTag:           `faker:"unix_time"`,
			},
			{
				Name:               "Email",
				GoType:             gotype.String,
				QueryType:          querytype.Equal,
				GraphQLType:        graphqltype.String,
				GraphQLDescription: "This is user's email",
				GraphQLRequired:    true,
				Validate:           `validate:"omitempty"`,
				FakerTag:           `faker:"email"`,
			},
			{
				Name:               "Name",
				GoType:             gotype.String,
				QueryType:          querytype.Equal,
				GraphQLType:        graphqltype.String,
				GraphQLDescription: "This is user's name",
				GraphQLRequired:    true,
				Validate:           `validate:"omitempty"`,
				FakerTag:           `faker:"name"`,
			},
			{
				Name:               "CompanyID",
				GoType:             gotype.Int64,
				QueryType:          querytype.Equal,
				GraphQLType:        graphqltype.Int,
				GraphQLDescription: "This is user's companyId",
				GraphQLRequired:    true,
				Validate:           `validate:"omitempty,min=1"`,
				FakerTag:           `faker:"unix_time"`,
			},
			{
				Name:               "Status",
				GoType:             gotype.String,
				QueryType:          querytype.Equal,
				GraphQLType:        graphqltype.String,
				GraphQLDescription: "This is user's status",
				GraphQLRequired:    true,
				Validate:           `validate:"omitempty,oneof=active inactive"`,
				FakerTag:           `faker:"name"`,
			},
			{
				Name:               "CreatedBy",
				GoType:             gotype.String,
				QueryType:          querytype.Equal,
				GraphQLType:        graphqltype.String,
				GraphQLDescription: "This is user's createdBy",
				GraphQLRequired:    false,
				Validate:           `validate:"omitempty"`,
				FakerTag:           `faker:"email"`,
			},
			{
				Name:               "UpdatedBy",
				GoType:             gotype.String,
				QueryType:          querytype.Equal,
				GraphQLType:        graphqltype.String,
				GraphQLDescription: "This is user's updatedBy",
				GraphQLRequired:    false,
				Validate:           `validate:"omitempty"`,
				FakerTag:           `faker:"email"`,
			},
		},
	}
)
