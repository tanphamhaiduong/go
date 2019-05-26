// @generated

package feature

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
		Name:        "Feature",
		Description: "This is type Feature",
		Fields: graphql.Fields{
			"id": &graphql.Field{
				Type:        graphql.NewNonNull(graphql.Int),
				Description: "This is feature id",
			},
			"url": &graphql.Field{
				Type:        graphql.NewNonNull(graphql.String),
				Description: "This is feature url",
			},
			"meta": &graphql.Field{
				Type:        graphql.NewNonNull(graphql.String),
				Description: "This is feature meta",
			},
			"companyId": &graphql.Field{
				Type:        graphql.NewNonNull(graphql.Int),
				Description: "This is feature companyId",
			},
			"status": &graphql.Field{
				Type:        graphql.NewNonNull(graphql.String),
				Description: "This is feature active",
			},
			"createdBy": &graphql.Field{
				Type:        graphql.String,
				Description: "This is feature createdBy",
			},
			"updatedBy": &graphql.Field{
				Type:        graphql.String,
				Description: "This is feature updatedBy",
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
			Description: "This is feature id",
		},
		"url": &graphql.ArgumentConfig{
			Type:        graphql.String,
			Description: "This is feature url",
		},
		"meta": &graphql.ArgumentConfig{
			Type:        graphql.String,
			Description: "This is feature meta",
		},
		"companyId": &graphql.ArgumentConfig{
			Type:        graphql.Int,
			Description: "This is feature companyId",
		},
		"status": &graphql.ArgumentConfig{
			Type:        graphql.String,
			Description: "This is feature active",
		},
		"createdBy": &graphql.ArgumentConfig{
			Type:        graphql.String,
			Description: "This is feature createdBy",
		},
		"updatedBy": &graphql.ArgumentConfig{
			Type:        graphql.String,
			Description: "This is feature updatedBy",
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
		"url": &graphql.ArgumentConfig{
			Type:        graphql.NewNonNull(graphql.String),
			Description: "This is feature url",
		},
		"meta": &graphql.ArgumentConfig{
			Type:        graphql.NewNonNull(graphql.String),
			Description: "This is feature meta",
		},
		"companyId": &graphql.ArgumentConfig{
			Type:        graphql.NewNonNull(graphql.Int),
			Description: "This is feature companyId",
		},
		"status": &graphql.ArgumentConfig{
			Type:        graphql.NewNonNull(graphql.String),
			Description: "This is feature active",
		},
		"createdBy": &graphql.ArgumentConfig{
			Type:        graphql.String,
			Description: "This is feature createdBy",
		},
		"updatedBy": &graphql.ArgumentConfig{
			Type:        graphql.String,
			Description: "This is feature updatedBy",
		},
	}

	// UpdateTypeArgs ...
	UpdateTypeArgs = graphql.FieldConfigArgument{
		"id": &graphql.ArgumentConfig{
			Type: graphql.NewNonNull(graphql.Int),
		},
		"url": &graphql.ArgumentConfig{
			Type:        graphql.String,
			Description: "This is feature url",
		},
		"meta": &graphql.ArgumentConfig{
			Type:        graphql.String,
			Description: "This is feature meta",
		},
		"companyId": &graphql.ArgumentConfig{
			Type:        graphql.Int,
			Description: "This is feature companyId",
		},
		"status": &graphql.ArgumentConfig{
			Type:        graphql.String,
			Description: "This is feature active",
		},
		"createdBy": &graphql.ArgumentConfig{
			Type:        graphql.String,
			Description: "This is feature createdBy",
		},
		"updatedBy": &graphql.ArgumentConfig{
			Type:        graphql.String,
			Description: "This is feature updatedBy",
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
	GetByID(ctx context.Context, params arguments.FeatureGetByIDArgs) (models.Feature, error)
	Count(ctx context.Context, params arguments.FeatureCountArgs) (int64, error)
	List(ctx context.Context, params arguments.FeatureListArgs) ([]models.Feature, error)
	Insert(ctx context.Context, params arguments.FeatureInsertArgs) (models.Feature, error)
	Update(ctx context.Context, params arguments.FeatureUpdateArgs) (models.Feature, error)
	Delete(ctx context.Context, params arguments.FeatureDeleteArgs) (int64, error)
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
	args := arguments.FeatureGetByIDArgs{}
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
	args := arguments.FeatureCountArgs{}
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
	args := arguments.FeatureListArgs{}
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
	args := arguments.FeatureInsertArgs{}
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
	args := arguments.FeatureUpdateArgs{}
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
	args := arguments.FeatureDeleteArgs{}
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
