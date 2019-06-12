// @generated
package modules

import (
	"github.com/graphql-go/graphql"
	"github.com/tanphamhaiduong/go/delta/internal/modules/company"
	"github.com/tanphamhaiduong/go/delta/internal/modules/feature"
	"github.com/tanphamhaiduong/go/delta/internal/modules/user"
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

// Resolver ...
type Resolver struct {
	Company ICompanyResolver
	Feature IFeatureResolver
	User    IUserResolver
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
			Description: "This is type companies",
			Fields: graphql.Fields{
				"records": &graphql.Field{
					Type:        graphql.NewNonNull(graphql.NewList(company.Type)),
					Description: "This is records of companies",
					Resolve:     resolver.Company.List,
				},
				"totalRecords": &graphql.Field{
					Type:        graphql.NewNonNull(graphql.Int),
					Description: "This is totalRecords of companies query",
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
			Description: "This is type features",
			Fields: graphql.Fields{
				"records": &graphql.Field{
					Type:        graphql.NewNonNull(graphql.NewList(feature.Type)),
					Description: "This is records of features",
					Resolve:     resolver.Feature.List,
				},
				"totalRecords": &graphql.Field{
					Type:        graphql.NewNonNull(graphql.Int),
					Description: "This is totalRecords of features query",
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
	rootQuery.AddFieldConfig("user", &graphql.Field{
		Type:        user.Type,
		Description: "This is GetByID for user",
		Args:        user.GetByIDTypeArgs,
		Resolve:     resolver.User.GetByID,
	})
	rootQuery.AddFieldConfig("users", &graphql.Field{
		Type: graphql.NewObject(graphql.ObjectConfig{
			Name:        "Users",
			Description: "This is type users",
			Fields: graphql.Fields{
				"records": &graphql.Field{
					Type:        graphql.NewNonNull(graphql.NewList(user.Type)),
					Description: "This is records of users",
					Resolve:     resolver.User.List,
				},
				"totalRecords": &graphql.Field{
					Type:        graphql.NewNonNull(graphql.Int),
					Description: "This is totalRecords of users query",
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

// NewResolver ...
func NewResolver(company ICompanyResolver, feature IFeatureResolver, user IUserResolver) Resolver {
	return Resolver{
		Company: company,
		Feature: feature,
		User:    user,
	}
}
