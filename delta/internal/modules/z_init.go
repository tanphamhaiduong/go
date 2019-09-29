// @generated
package modules

import (
	"github.com/graphql-go/graphql"
	"github.com/tanphamhaiduong/go/delta/internal/modules/company"
	"github.com/tanphamhaiduong/go/delta/internal/modules/permission"
	"github.com/tanphamhaiduong/go/delta/internal/modules/role"
	"github.com/tanphamhaiduong/go/delta/internal/modules/user"
)

// Resolver ...
type Resolver struct {
	Company        ICompanyResolver
	Permission     IPermissionResolver
	Role           IRoleResolver
	RolePermission IRolePermissionResolver
	User           IUserResolver
}

// Handler ...
type Handler struct {
	Company        ICompanyHandler
	Permission     IPermissionHandler
	Role           IRoleHandler
	RolePermission IRolePermissionHandler
	User           IUserHandler
}

func addToSchema(resolver Resolver) {
	rootQuery.AddFieldConfig("company", &graphql.Field{
		Type:        company.Type,
		Description: "This is GetByID for company",
		Args:        company.GetByIDTypeArgs,
		Resolve:     resolver.Company.GetByID,
	})
	rootQuery.AddFieldConfig("companies", &graphql.Field{
		Type:        graphql.NewNonNull(graphql.NewList(company.Type)),
		Description: "This is get list of company query",
		Args:        company.ListTypeArgs,
		Resolve:     resolver.Company.List,
	})
	rootMutation.AddFieldConfig("insertCompany", &graphql.Field{
		Type:        company.Type,
		Description: "This is insert company query",
		Args:        company.InsertTypeArgs,
		Resolve:     resolver.Company.Insert,
	})
	rootMutation.AddFieldConfig("updateCompany", &graphql.Field{
		Type:        company.Type,
		Description: "This is update company query",
		Args:        company.UpdateTypeArgs,
		Resolve:     resolver.Company.Update,
	})
	rootQuery.AddFieldConfig("permission", &graphql.Field{
		Type:        permission.Type,
		Description: "This is GetByID for permission",
		Args:        permission.GetByIDTypeArgs,
		Resolve:     resolver.Permission.GetByID,
	})
	rootQuery.AddFieldConfig("permissions", &graphql.Field{
		Type:        graphql.NewNonNull(graphql.NewList(permission.Type)),
		Description: "This is get list of permission query",
		Args:        permission.ListTypeArgs,
		Resolve:     resolver.Permission.List,
	})
	rootMutation.AddFieldConfig("insertPermission", &graphql.Field{
		Type:        permission.Type,
		Description: "This is insert permission query",
		Args:        permission.InsertTypeArgs,
		Resolve:     resolver.Permission.Insert,
	})
	rootQuery.AddFieldConfig("role", &graphql.Field{
		Type:        role.Type,
		Description: "This is GetByID for role",
		Args:        role.GetByIDTypeArgs,
		Resolve:     resolver.Role.GetByID,
	})
	rootQuery.AddFieldConfig("roles", &graphql.Field{
		Type:        graphql.NewNonNull(graphql.NewList(role.Type)),
		Description: "This is get list of role query",
		Args:        role.ListTypeArgs,
		Resolve:     resolver.Role.List,
	})
	rootMutation.AddFieldConfig("insertRole", &graphql.Field{
		Type:        role.Type,
		Description: "This is insert role query",
		Args:        role.InsertTypeArgs,
		Resolve:     resolver.Role.Insert,
	})
	rootMutation.AddFieldConfig("updateRole", &graphql.Field{
		Type:        role.Type,
		Description: "This is update role query",
		Args:        role.UpdateTypeArgs,
		Resolve:     resolver.Role.Update,
	})
	rootQuery.AddFieldConfig("user", &graphql.Field{
		Type:        user.Type,
		Description: "This is GetByID for user",
		Args:        user.GetByIDTypeArgs,
		Resolve:     resolver.User.GetByID,
	})
	rootQuery.AddFieldConfig("users", &graphql.Field{
		Type:        graphql.NewNonNull(graphql.NewList(user.Type)),
		Description: "This is get list of user query",
		Args:        user.ListTypeArgs,
		Resolve:     resolver.User.List,
	})
	rootMutation.AddFieldConfig("insertUser", &graphql.Field{
		Type:        user.Type,
		Description: "This is insert user query",
		Args:        user.InsertTypeArgs,
		Resolve:     resolver.User.Insert,
	})
	rootMutation.AddFieldConfig("updateUser", &graphql.Field{
		Type:        user.Type,
		Description: "This is update user query",
		Args:        user.UpdateTypeArgs,
		Resolve:     resolver.User.Update,
	})
}

// NewResolver ...
func NewResolver(company ICompanyResolver, permission IPermissionResolver, role IRoleResolver, rolepermission IRolePermissionResolver, user IUserResolver) Resolver {
	return Resolver{
		Company:        company,
		Permission:     permission,
		Role:           role,
		RolePermission: rolepermission,
		User:           user,
	}
}

// NewHandler ...
func NewHandler(company ICompanyHandler, permission IPermissionHandler, role IRoleHandler, rolepermission IRolePermissionHandler, user IUserHandler) Handler {
	return Handler{
		Company:        company,
		Permission:     permission,
		Role:           role,
		RolePermission: rolepermission,
		User:           user,
	}
}
