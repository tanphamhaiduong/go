package modules

import (
	"github.com/tanphamhaiduong/go/delta/pkg/generator/internal/gotype"
	"github.com/tanphamhaiduong/go/delta/pkg/generator/internal/graphqltype"
	"github.com/tanphamhaiduong/go/delta/pkg/generator/internal/querytype"
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
				GraphQLDescription: "This is company's id",
				GraphQLRequired:    true,
				Validate:           `validate:"required,min=1"`,
				FakerTag:           `faker:"unix_time"`,
			},
			{
				Name:               "Name",
				GoType:             gotype.String,
				QueryType:          querytype.Like,
				GraphQLType:        graphqltype.String,
				GraphQLDescription: "This is company's name",
				GraphQLRequired:    true,
				Validate:           `validate:"omitempty"`,
				FakerTag:           `faker:"name"`,
			},
			{
				Name:               "CompanyCode",
				GoType:             gotype.String,
				QueryType:          querytype.Equal,
				GraphQLType:        graphqltype.String,
				GraphQLDescription: "This is company's code",
				GraphQLRequired:    true,
				Validate:           `validate:"omitempty"`,
				FakerTag:           `faker:"name"`,
			},
			{
				Name:               "Status",
				GoType:             gotype.String,
				QueryType:          querytype.Equal,
				GraphQLType:        graphqltype.String,
				GraphQLDescription: "This is company's status",
				GraphQLRequired:    true,
				Validate:           `validate:"omitempty,oneof=active inactive"`,
				FakerTag:           `faker:"word"`,
			},
			{
				Name:               "CreatedBy",
				GoType:             gotype.String,
				QueryType:          querytype.Equal,
				GraphQLType:        graphqltype.String,
				GraphQLDescription: "This is company's createdBy",
				GraphQLRequired:    false,
				Validate:           `validate:"omitempty"`,
				FakerTag:           `faker:"email"`,
			},
			{
				Name:               "UpdatedBy",
				GoType:             gotype.String,
				QueryType:          querytype.Equal,
				GraphQLType:        graphqltype.String,
				GraphQLDescription: "This is company's updatedBy",
				GraphQLRequired:    false,
				Validate:           `validate:"omitempty"`,
				FakerTag:           `faker:"email"`,
			},
			{
				Name:               "ApiSecretKey",
				GoType:             gotype.String,
				QueryType:          querytype.Equal,
				GraphQLType:        graphqltype.String,
				GraphQLDescription: "This is company's updatedBy",
				GraphQLRequired:    false,
				Validate:           `validate:"omitempty"`,
				FakerTag:           `faker:"name"`,
			},
		},
	}
)
