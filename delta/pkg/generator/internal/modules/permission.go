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
		Name:          "Permission",
		IsHaveGetByID: true,
		IsHaveCount:   true,
		IsHaveList:    true,
		IsHaveInsert:  true,
		IsHaveUpdate:  false,
		Fields: []Field{
			{
				Name:               "ID",
				GoType:             gotype.Int64,
				QueryType:          querytype.Equal,
				GraphQLType:        graphqltype.Int,
				GraphQLDescription: "This is permission's id",
				GraphQLRequired:    true,
				Validate:           `validate:"required,min=1"`,
			},
			{
				Name:               "Name",
				GoType:             gotype.String,
				QueryType:          querytype.Equal,
				GraphQLType:        graphqltype.String,
				GraphQLDescription: "This is permission's name",
				GraphQLRequired:    true,
				Validate:           `validate:"omitempty"`,
			},
			{
				Name:               "Description",
				GoType:             gotype.String,
				QueryType:          querytype.Equal,
				GraphQLType:        graphqltype.String,
				GraphQLDescription: "This is permission's description",
				GraphQLRequired:    true,
				Validate:           `validate:"omitempty"`,
			},
			{
				Name:               "Status",
				GoType:             gotype.String,
				QueryType:          querytype.Equal,
				GraphQLType:        graphqltype.String,
				GraphQLDescription: "This is permission's active",
				GraphQLRequired:    true,
				Validate:           `validate:"omitempty,oneof=active inactive"`,
			},
			{
				Name:               "CreatedBy",
				GoType:             gotype.String,
				QueryType:          querytype.Equal,
				GraphQLType:        graphqltype.String,
				GraphQLDescription: "This is permission's createdBy",
				GraphQLRequired:    false,
				Validate:           `validate:"omitempty"`,
			},
			{
				Name:               "UpdatedBy",
				GoType:             gotype.String,
				QueryType:          querytype.Equal,
				GraphQLType:        graphqltype.String,
				GraphQLDescription: "This is permission's updatedBy",
				GraphQLRequired:    false,
				Validate:           `validate:"omitempty"`,
			},
		},
	}
)
