package modules

// import (
// 	"github.com/tanphamhaiduong/go/delta/pkg/generator/internal/gotype"
// 	"github.com/tanphamhaiduong/go/delta/pkg/generator/internal/graphqltype"
// 	"github.com/tanphamhaiduong/go/delta/pkg/generator/internal/querytype"
// )

// func init() {
// 	modules = append(modules, rolePermissionModule)
// }

// var (
// 	rolePermissionModule = Module{
// 		Name: "RolePermission",
// 		Fields: []Field{
// 			{
// 				Name:               "ID",
// 				GoType:             gotype.Int64,
// 				QueryType:          querytype.Equal,
// 				GraphQLType:        graphqltype.Int,
// 				GraphQLDescription: "This is rolePermission's id",
// 				GraphQLRequired:    true,
// 				Validate:           `validate:"required,min=1"`,
// 			},
// 			{
// 				Name:               "RoleID",
// 				GoType:             gotype.Int64,
// 				QueryType:          querytype.Equal,
// 				GraphQLType:        graphqltype.Int,
// 				GraphQLDescription: "This is rolePermission's roleId",
// 				GraphQLRequired:    true,
// 				Validate:           `validate:"required,min=1"`,
// 			},
// 			{
// 				Name:               "PermissionID",
// 				GoType:             gotype.Int64,
// 				QueryType:          querytype.Equal,
// 				GraphQLType:        graphqltype.Int,
// 				GraphQLDescription: "This is rolePermission's permissionId",
// 				GraphQLRequired:    true,
// 				Validate:           `validate:"required,min=1"`,
// 			},
// 			{
// 				Name:               "CreatedBy",
// 				GoType:             gotype.String,
// 				QueryType:          querytype.Equal,
// 				GraphQLType:        graphqltype.String,
// 				GraphQLDescription: "This is rolePermission's createdBy",
// 				GraphQLRequired:    false,
// 				Validate:           `validate:"omitempty`,
// 			},
// 			{
// 				Name:               "UpdatedBy",
// 				GoType:             gotype.String,
// 				QueryType:          querytype.Equal,
// 				GraphQLType:        graphqltype.String,
// 				GraphQLDescription: "This is rolePermission's updatedBy",
// 				GraphQLRequired:    false,
// 				Validate:           `validate:"omitempty`,
// 			},
// 		},
// 	}
// )
