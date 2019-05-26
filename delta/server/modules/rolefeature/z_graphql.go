// @generated

package rolefeature

import (
	"context"
	"database/sql"

	"github.com/graphql-go/graphql"
	"github.com/tanphamhaiduong/go/delta/server/arguments"
	"github.com/tanphamhaiduong/go/delta/server/models"
	"github.com/tanphamhaiduong/go/delta/server/utils"
)

var (
	// Type ...
	Type = graphql.NewObject(graphql.ObjectConfig{
		Name:        "RoleFeature",
		Description: "This is type Feature",
		Fields: graphql.Fields{
			"id": &graphql.Field{
				Type:        graphql.NewNonNull(graphql.Int),
				Description: "This is roleFeature id",
			},
			"roleId": &graphql.Field{
				Type:        graphql.NewNonNull(graphql.Int),
				Description: "This is roleFeature id",
			},
			"featureId": &graphql.Field{
				Type:        graphql.NewNonNull(graphql.Int),
				Description: "This is roleFeature id",
			},
			"createdBy": &graphql.Field{
				Type:        graphql.String,
				Description: "This is roleFeature createdBy",
			},
			"updatedBy": &graphql.Field{
				Type:        graphql.String,
				Description: "This is roleFeature updatedBy",
			},
		},
	})

	// GetByIDTypeArgs ...
	GetByIDTypeArgs = graphql.FieldConfigArgument{
		"id": &graphql.ArgumentConfig{
			Type: graphql.NewNonNull(graphql.Int),
		},
	}

	// ListTypeArgs ...
	ListTypeArgs = graphql.FieldConfigArgument{
		"id": &graphql.ArgumentConfig{
			Type:        graphql.Int,
			Description: "This is roleFeature id",
		},
		"roleId": &graphql.ArgumentConfig{
			Type:        graphql.Int,
			Description: "This is roleFeature id",
		},
		"featureId": &graphql.ArgumentConfig{
			Type:        graphql.Int,
			Description: "This is roleFeature id",
		},
		"createdBy": &graphql.ArgumentConfig{
			Type:        graphql.String,
			Description: "This is roleFeature createdBy",
		},
		"updatedBy": &graphql.ArgumentConfig{
			Type:        graphql.String,
			Description: "This is roleFeature updatedBy",
		},
		"page": &graphql.ArgumentConfig{
			Type:        graphql.NewNonNull(graphql.Int),
			Description: "This is feature page",
		},
		"pageSize": &graphql.ArgumentConfig{
			Type:         graphql.Int,
			Description:  "This is feature pageSize",
			DefaultValue: 10,
		},
	}

	// InsertTypeArgs ...
	InsertTypeArgs = graphql.FieldConfigArgument{
		"roleId": &graphql.ArgumentConfig{
			Type:        graphql.NewNonNull(graphql.Int),
			Description: "This is roleFeature id",
		},
		"featureId": &graphql.ArgumentConfig{
			Type:        graphql.NewNonNull(graphql.Int),
			Description: "This is roleFeature id",
		},
		"createdBy": &graphql.ArgumentConfig{
			Type:        graphql.String,
			Description: "This is roleFeature createdBy",
		},
		"updatedBy": &graphql.ArgumentConfig{
			Type:        graphql.String,
			Description: "This is roleFeature updatedBy",
		},
	}

	// UpdateTypeArgs ...
	UpdateTypeArgs = graphql.FieldConfigArgument{
		"id": &graphql.ArgumentConfig{
			Type: graphql.NewNonNull(graphql.Int),
		},
		"roleId": &graphql.ArgumentConfig{
			Type:        graphql.Int,
			Description: "This is roleFeature id",
		},
		"featureId": &graphql.ArgumentConfig{
			Type:        graphql.Int,
			Description: "This is roleFeature id",
		},
		"createdBy": &graphql.ArgumentConfig{
			Type:        graphql.String,
			Description: "This is roleFeature createdBy",
		},
		"updatedBy": &graphql.ArgumentConfig{
			Type:        graphql.String,
			Description: "This is roleFeature updatedBy",
		},
	}

	// DeleteTypeArgs ...
	DeleteTypeArgs = graphql.FieldConfigArgument{
		"id": &graphql.ArgumentConfig{
			Type: graphql.NewNonNull(graphql.Int),
		},
	}
)

// ICoreHandler ...
type ICoreHandler interface {
	GetByID(ctx context.Context, params arguments.RoleFeatureGetByIDArgs) (models.RoleFeature, error)
	Count(ctx context.Context, params arguments.RoleFeatureCountArgs) (int64, error)
	List(ctx context.Context, params arguments.RoleFeatureListArgs) ([]models.RoleFeature, error)
	Insert(ctx context.Context, params arguments.RoleFeatureInsertArgs) (models.RoleFeature, error)
	Update(ctx context.Context, params arguments.RoleFeatureUpdateArgs) (models.RoleFeature, error)
	Delete(ctx context.Context, params arguments.RoleFeatureDeleteArgs) (int64, error)
}

// ResolverImpl ...
type ResolverImpl struct {
	db      *sql.DB
	handler IHandler
}

// NewResolver ...
func NewResolver(db *sql.DB) ResolverImpl {
	return ResolverImpl{
		db:      db,
		handler: NewHandler(db),
	}
}

// ForwardParams ...
func (r ResolverImpl) ForwardParams(params graphql.ResolveParams) (interface{}, error) {
	return params.Args, nil
}

// GetByID ...
func (r ResolverImpl) GetByID(params graphql.ResolveParams) (interface{}, error) {
	// parse params
	args := arguments.RoleFeatureGetByIDArgs{}
	if err := utils.Parse(params.Args, &args); err != nil {
		return nil, err
	}
	response, err := r.handler.GetByID(params.Context, args)
	if err != nil {
		return nil, err
	}
	return response, nil
}

// Count ...
func (r ResolverImpl) Count(params graphql.ResolveParams) (interface{}, error) {
	// parse params
	args := arguments.RoleFeatureCountArgs{}
	err := utils.Parse(params.Source.(map[string]interface{}), &args)
	if err != nil {
		return nil, err
	}
	response, err := r.handler.Count(params.Context, args)
	if err != nil {
		return nil, err
	}
	return response, nil
}

// List ...
func (r ResolverImpl) List(params graphql.ResolveParams) (interface{}, error) {
	// parse params
	args := arguments.RoleFeatureListArgs{}
	err := utils.Parse(params.Source.(map[string]interface{}), &args)
	if err != nil {
		return nil, err
	}
	response, err := r.handler.List(params.Context, args)
	if err != nil {
		return nil, err
	}
	return response, nil
}

// Insert ...
func (r ResolverImpl) Insert(params graphql.ResolveParams) (interface{}, error) {
	// parse params
	args := arguments.RoleFeatureInsertArgs{}
	err := utils.Parse(params.Args, &args)
	if err != nil {
		return nil, err
	}
	response, err := r.handler.Insert(params.Context, args)
	if err != nil {
		return nil, err
	}
	return response, nil
}

// Update ...
func (r ResolverImpl) Update(params graphql.ResolveParams) (interface{}, error) {
	// parse params
	args := arguments.RoleFeatureUpdateArgs{}
	err := utils.Parse(params.Args, &args)
	if err != nil {
		return nil, err
	}
	response, err := r.handler.Update(params.Context, args)
	if err != nil {
		return nil, err
	}
	return response, nil
}

// Delete ...
func (r ResolverImpl) Delete(params graphql.ResolveParams) (interface{}, error) {
	// parse params
	args := arguments.RoleFeatureDeleteArgs{}
	err := utils.Parse(params.Args, &args)
	if err != nil {
		return nil, err
	}
	response, err := r.handler.Delete(params.Context, args)
	if err != nil {
		return nil, err
	}
	return response, nil
}

//go:generate mockery -name=IHandler -output=mocks -case=underscore
