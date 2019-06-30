package modules

import (
	"github.com/tanphamhaiduong/go/delta/pkg/generator/internal/gotype"
	"github.com/tanphamhaiduong/go/delta/pkg/generator/internal/graphqltype"
	"github.com/tanphamhaiduong/go/delta/pkg/generator/internal/querytype"
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
				GraphQLDescription: "This is feature's id",
				GraphQLRequired:    true,
				Validate:           `validate:"required,min=1"`,
				FakerTag:           `faker:"unix_time"`,
			},
			{
				Name:               "URL",
				GoType:             gotype.String,
				QueryType:          querytype.Equal,
				GraphQLType:        graphqltype.String,
				GraphQLDescription: "This is feature's url",
				GraphQLRequired:    true,
				Validate:           `validate:"omitempty"`,
				FakerTag:           `faker:"name"`,
			},
			{
				Name:               "Meta",
				GoType:             gotype.String,
				QueryType:          querytype.Equal,
				GraphQLType:        graphqltype.String,
				GraphQLDescription: "This is feature's meta",
				GraphQLRequired:    true,
				Validate:           `validate:"omitempty"`,
				FakerTag:           `faker:"name"`,
			},
			{
				Name:               "CompanyID",
				GoType:             gotype.Int64,
				QueryType:          querytype.Equal,
				GraphQLType:        graphqltype.Int,
				GraphQLDescription: "This is feature's companyId",
				GraphQLRequired:    true,
				Validate:           `validate:"omitempty,min=1"`,
				FakerTag:           `faker:"unix_time"`,
			},
			{
				Name:               "Status",
				GoType:             gotype.String,
				QueryType:          querytype.Equal,
				GraphQLType:        graphqltype.String,
				GraphQLDescription: "This is feature's active",
				GraphQLRequired:    true,
				Validate:           `validate:"omitempty,oneof=active inactive"`,
				FakerTag:           `faker:"name"`,
			},
			{
				Name:               "CreatedBy",
				GoType:             gotype.String,
				QueryType:          querytype.Equal,
				GraphQLType:        graphqltype.String,
				GraphQLDescription: "This is feature's createdBy",
				GraphQLRequired:    false,
				Validate:           `validate:"omitempty"`,
				FakerTag:           `faker:"email"`,
			},
			{
				Name:               "UpdatedBy",
				GoType:             gotype.String,
				QueryType:          querytype.Equal,
				GraphQLType:        graphqltype.String,
				GraphQLDescription: "This is feature's updatedBy",
				GraphQLRequired:    false,
				Validate:           `validate:"omitempty"`,
				FakerTag:           `faker:"email"`,
			},
		},
	}
)
