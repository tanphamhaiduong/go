package modules

import (
	"github.com/tanphamhaiduong/go/delta/pkg/generator/internal/gotype"
	"github.com/tanphamhaiduong/go/delta/pkg/generator/internal/graphqltype"
	"github.com/tanphamhaiduong/go/delta/pkg/generator/internal/querytype"
)

func init() {
	modules = append(modules, permissionModule)
}

var (
	permissionModule = Module{
		Name: "Permission",
		Fields: []Field{
			{
				Name:               "ID",
				GoType:             gotype.Int64,
				QueryType:          querytype.Equal,
				GraphQLType:        graphqltype.Int,
				GraphQLDescription: "This is permission's id",
				GraphQLRequired:    true,
				Validate:           `validate:"required,min=1"`,
				FakerTag:           `faker:"unix_time"`,
			},
			{
				Name:               "Name",
				GoType:             gotype.String,
				QueryType:          querytype.Equal,
				GraphQLType:        graphqltype.String,
				GraphQLDescription: "This is permission's name",
				GraphQLRequired:    true,
				Validate:           `validate:"omitempty"`,
				FakerTag:           `faker:"name"`,
			},
			{
				Name:               "Status",
				GoType:             gotype.String,
				QueryType:          querytype.Equal,
				GraphQLType:        graphqltype.String,
				GraphQLDescription: "This is permission's active",
				GraphQLRequired:    true,
				Validate:           `validate:"omitempty,oneof=active inactive"`,
				FakerTag:           `faker:"name"`,
			},
			{
				Name:               "CreatedBy",
				GoType:             gotype.String,
				QueryType:          querytype.Equal,
				GraphQLType:        graphqltype.String,
				GraphQLDescription: "This is permission's createdBy",
				GraphQLRequired:    false,
				Validate:           `validate:"omitempty"`,
				FakerTag:           `faker:"email"`,
			},
			{
				Name:               "UpdatedBy",
				GoType:             gotype.String,
				QueryType:          querytype.Equal,
				GraphQLType:        graphqltype.String,
				GraphQLDescription: "This is permission's updatedBy",
				GraphQLRequired:    false,
				Validate:           `validate:"omitempty"`,
				FakerTag:           `faker:"email"`,
			},
		},
	}
)
