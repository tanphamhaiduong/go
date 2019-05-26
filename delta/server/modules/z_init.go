// @generated

package modules

import (
	"database/sql"

	"github.com/graphql-go/graphql"
	"github.com/tanphamhaiduong/go/delta/server/modules/company"
	"github.com/tanphamhaiduong/go/delta/server/modules/feature"
	"github.com/tanphamhaiduong/go/delta/server/modules/role"
	"github.com/tanphamhaiduong/go/delta/server/modules/rolefeature"
	"github.com/tanphamhaiduong/go/delta/server/modules/user"
)

// ICoreResolver ...
type IResolver interface {
	ForwardParams(p graphql.ResolveParams) (interface{}, error)
	GetByID(p graphql.ResolveParams) (interface{}, error)
	Count(p graphql.ResolveParams) (interface{}, error)
	List(p graphql.ResolveParams) (interface{}, error)
	Insert(p graphql.ResolveParams) (interface{}, error)
	Update(p graphql.ResolveParams) (interface{}, error)
	Delete(p graphql.ResolveParams) (interface{}, error)
}

// Handler ...
type Handler struct {
	Company     ICompanyHandler
	Feature     IFeatureHandler
	Role        IRoleHandler
	RoleFeature IRoleFeatureHandler
	User        IUserHandler
}

// Resolver ...
type Resolver struct {
	Company     ICompanyResolver
	Feature     IFeatureResolver
	Role        IRoleResolver
	RoleFeature IRoleFeatureResolver
	User        IUserResolver
}

func addToSchema(resolver Resolver) {
	rootQuery.AddFieldConfig("company", &graphql.Field{
		Type:        company.Type,
		Description: "This is GetByID for company",
		Args:        company.GetByIDTypeArgs,
		Resolve:     resolver.Company.GetByID,
	})
	rootQuery.AddFieldConfig("companies", &graphql.Field{
		Type: graphql.NewObject(graphql.ObjectConfig{
			Name:        "Companies",
			Description: "This is typecompanies",
			Fields: graphql.Fields{
				"records": &graphql.Field{
					Type:        graphql.NewNonNull(graphql.NewList(company.Type)),
					Description: "This is records ofcompanies",
					Resolve:     resolver.Company.List,
				},
				"totalRecords": &graphql.Field{
					Type:        graphql.NewNonNull(graphql.Int),
					Description: "This is totalRecords ofcompaniesquery",
					Resolve:     resolver.Company.Count,
				},
			},
		}),
		Description: "This is get list of company query",
		Args:        company.ListTypeArgs,
		Resolve:     resolver.Company.ForwardParams,
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
	rootMutation.AddFieldConfig("deleteCompany", &graphql.Field{
		Type:        graphql.NewNonNull(graphql.Int),
		Description: "This is delete company query",
		Args:        company.DeleteTypeArgs,
		Resolve:     resolver.Company.Delete,
	})
	rootQuery.AddFieldConfig("feature", &graphql.Field{
		Type:        feature.Type,
		Description: "This is GetByID for feature",
		Args:        feature.GetByIDTypeArgs,
		Resolve:     resolver.Feature.GetByID,
	})
	rootQuery.AddFieldConfig("features", &graphql.Field{
		Type: graphql.NewObject(graphql.ObjectConfig{
			Name:        "Features",
			Description: "This is typefeatures",
			Fields: graphql.Fields{
				"records": &graphql.Field{
					Type:        graphql.NewNonNull(graphql.NewList(feature.Type)),
					Description: "This is records offeatures",
					Resolve:     resolver.Feature.List,
				},
				"totalRecords": &graphql.Field{
					Type:        graphql.NewNonNull(graphql.Int),
					Description: "This is totalRecords offeaturesquery",
					Resolve:     resolver.Feature.Count,
				},
			},
		}),
		Description: "This is get list of feature query",
		Args:        feature.ListTypeArgs,
		Resolve:     resolver.Feature.ForwardParams,
	})
	rootMutation.AddFieldConfig("insertFeature", &graphql.Field{
		Type:        feature.Type,
		Description: "This is insert feature query",
		Args:        feature.InsertTypeArgs,
		Resolve:     resolver.Feature.Insert,
	})
	rootMutation.AddFieldConfig("updateFeature", &graphql.Field{
		Type:        feature.Type,
		Description: "This is update feature query",
		Args:        feature.UpdateTypeArgs,
		Resolve:     resolver.Feature.Update,
	})
	rootMutation.AddFieldConfig("deleteFeature", &graphql.Field{
		Type:        graphql.NewNonNull(graphql.Int),
		Description: "This is delete feature query",
		Args:        feature.DeleteTypeArgs,
		Resolve:     resolver.Feature.Delete,
	})
	rootQuery.AddFieldConfig("role", &graphql.Field{
		Type:        role.Type,
		Description: "This is GetByID for role",
		Args:        role.GetByIDTypeArgs,
		Resolve:     resolver.Role.GetByID,
	})
	rootQuery.AddFieldConfig("roles", &graphql.Field{
		Type: graphql.NewObject(graphql.ObjectConfig{
			Name:        "Roles",
			Description: "This is typeroles",
			Fields: graphql.Fields{
				"records": &graphql.Field{
					Type:        graphql.NewNonNull(graphql.NewList(role.Type)),
					Description: "This is records ofroles",
					Resolve:     resolver.Role.List,
				},
				"totalRecords": &graphql.Field{
					Type:        graphql.NewNonNull(graphql.Int),
					Description: "This is totalRecords ofrolesquery",
					Resolve:     resolver.Role.Count,
				},
			},
		}),
		Description: "This is get list of role query",
		Args:        role.ListTypeArgs,
		Resolve:     resolver.Role.ForwardParams,
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
	rootMutation.AddFieldConfig("deleteRole", &graphql.Field{
		Type:        graphql.NewNonNull(graphql.Int),
		Description: "This is delete role query",
		Args:        role.DeleteTypeArgs,
		Resolve:     resolver.Role.Delete,
	})
	rootQuery.AddFieldConfig("roleFeature", &graphql.Field{
		Type:        rolefeature.Type,
		Description: "This is GetByID for roleFeature",
		Args:        rolefeature.GetByIDTypeArgs,
		Resolve:     resolver.RoleFeature.GetByID,
	})
	rootQuery.AddFieldConfig("roleFeatures", &graphql.Field{
		Type: graphql.NewObject(graphql.ObjectConfig{
			Name:        "RoleFeatures",
			Description: "This is typeroleFeatures",
			Fields: graphql.Fields{
				"records": &graphql.Field{
					Type:        graphql.NewNonNull(graphql.NewList(rolefeature.Type)),
					Description: "This is records ofroleFeatures",
					Resolve:     resolver.RoleFeature.List,
				},
				"totalRecords": &graphql.Field{
					Type:        graphql.NewNonNull(graphql.Int),
					Description: "This is totalRecords ofroleFeaturesquery",
					Resolve:     resolver.RoleFeature.Count,
				},
			},
		}),
		Description: "This is get list of roleFeature query",
		Args:        rolefeature.ListTypeArgs,
		Resolve:     resolver.RoleFeature.ForwardParams,
	})
	rootMutation.AddFieldConfig("insertRoleFeature", &graphql.Field{
		Type:        rolefeature.Type,
		Description: "This is insert roleFeature query",
		Args:        rolefeature.InsertTypeArgs,
		Resolve:     resolver.RoleFeature.Insert,
	})
	rootMutation.AddFieldConfig("updateRoleFeature", &graphql.Field{
		Type:        rolefeature.Type,
		Description: "This is update roleFeature query",
		Args:        rolefeature.UpdateTypeArgs,
		Resolve:     resolver.RoleFeature.Update,
	})
	rootMutation.AddFieldConfig("deleteRoleFeature", &graphql.Field{
		Type:        graphql.NewNonNull(graphql.Int),
		Description: "This is delete roleFeature query",
		Args:        rolefeature.DeleteTypeArgs,
		Resolve:     resolver.RoleFeature.Delete,
	})
	rootQuery.AddFieldConfig("user", &graphql.Field{
		Type:        user.Type,
		Description: "This is GetByID for user",
		Args:        user.GetByIDTypeArgs,
		Resolve:     resolver.User.GetByID,
	})
	rootQuery.AddFieldConfig("users", &graphql.Field{
		Type: graphql.NewObject(graphql.ObjectConfig{
			Name:        "Users",
			Description: "This is typeusers",
			Fields: graphql.Fields{
				"records": &graphql.Field{
					Type:        graphql.NewNonNull(graphql.NewList(user.Type)),
					Description: "This is records ofusers",
					Resolve:     resolver.User.List,
				},
				"totalRecords": &graphql.Field{
					Type:        graphql.NewNonNull(graphql.Int),
					Description: "This is totalRecords ofusersquery",
					Resolve:     resolver.User.Count,
				},
			},
		}),
		Description: "This is get list of user query",
		Args:        user.ListTypeArgs,
		Resolve:     resolver.User.ForwardParams,
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
	rootMutation.AddFieldConfig("deleteUser", &graphql.Field{
		Type:        graphql.NewNonNull(graphql.Int),
		Description: "This is delete user query",
		Args:        user.DeleteTypeArgs,
		Resolve:     resolver.User.Delete,
	})
}

// NewHandler ...
func NewHandler(db *sql.DB) Handler {
	return Handler{
		Company:     company.NewHandler(db),
		Feature:     feature.NewHandler(db),
		Role:        role.NewHandler(db),
		RoleFeature: rolefeature.NewHandler(db),
		User:        user.NewHandler(db),
	}
}

// NewResolver ...
func NewResolver(db *sql.DB) Resolver {
	return Resolver{
		Company:     company.NewResolver(db),
		Feature:     feature.NewResolver(db),
		Role:        role.NewResolver(db),
		RoleFeature: rolefeature.NewResolver(db),
		User:        user.NewResolver(db),
	}
}
