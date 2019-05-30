// @generated
package user

import (
	"context"

	"github.com/graphql-go/graphql"
	"github.com/tanphamhaiduong/go/delta/server/arguments"
	"github.com/tanphamhaiduong/go/delta/server/models"
	"github.com/tanphamhaiduong/go/delta/server/utils"
)

var (
	// Type ...
	Type = graphql.NewObject(graphql.ObjectConfig{
		Name:        "User",
		Description: "This is type Feature",
		Fields: graphql.Fields{
			"id": &graphql.Field{
				Type:        graphql.NewNonNull(graphql.Int),
				Description: "This is user's id",
			},
			"email": &graphql.Field{
				Type:        graphql.NewNonNull(graphql.String),
				Description: "This is user's email",
			},
			"name": &graphql.Field{
				Type:        graphql.NewNonNull(graphql.String),
				Description: "This is user's name",
			},
			"companyId": &graphql.Field{
				Type:        graphql.NewNonNull(graphql.Int),
				Description: "This is user's companyId",
			},
			"status": &graphql.Field{
				Type:        graphql.NewNonNull(graphql.String),
				Description: "This is user's status",
			},
			"createdBy": &graphql.Field{
				Type:        graphql.String,
				Description: "This is user's createdBy",
			},
			"updatedBy": &graphql.Field{
				Type:        graphql.String,
				Description: "This is user's updatedBy",
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
			Description: "This is user's id",
		},
		"email": &graphql.ArgumentConfig{
			Type:        graphql.String,
			Description: "This is user's email",
		},
		"name": &graphql.ArgumentConfig{
			Type:        graphql.String,
			Description: "This is user's name",
		},
		"companyId": &graphql.ArgumentConfig{
			Type:        graphql.Int,
			Description: "This is user's companyId",
		},
		"status": &graphql.ArgumentConfig{
			Type:        graphql.String,
			Description: "This is user's status",
		},
		"createdBy": &graphql.ArgumentConfig{
			Type:        graphql.String,
			Description: "This is user's createdBy",
		},
		"updatedBy": &graphql.ArgumentConfig{
			Type:        graphql.String,
			Description: "This is user's updatedBy",
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
		"email": &graphql.ArgumentConfig{
			Type:        graphql.NewNonNull(graphql.String),
			Description: "This is user's email",
		},
		"name": &graphql.ArgumentConfig{
			Type:        graphql.NewNonNull(graphql.String),
			Description: "This is user's name",
		},
		"companyId": &graphql.ArgumentConfig{
			Type:        graphql.NewNonNull(graphql.Int),
			Description: "This is user's companyId",
		},
		"status": &graphql.ArgumentConfig{
			Type:        graphql.NewNonNull(graphql.String),
			Description: "This is user's status",
		},
		"createdBy": &graphql.ArgumentConfig{
			Type:        graphql.String,
			Description: "This is user's createdBy",
		},
		"updatedBy": &graphql.ArgumentConfig{
			Type:        graphql.String,
			Description: "This is user's updatedBy",
		},
	}

	// UpdateTypeArgs ...
	UpdateTypeArgs = graphql.FieldConfigArgument{
		"id": &graphql.ArgumentConfig{
			Type: graphql.NewNonNull(graphql.Int),
		},
		"email": &graphql.ArgumentConfig{
			Type:        graphql.String,
			Description: "This is user's email",
		},
		"name": &graphql.ArgumentConfig{
			Type:        graphql.String,
			Description: "This is user's name",
		},
		"companyId": &graphql.ArgumentConfig{
			Type:        graphql.Int,
			Description: "This is user's companyId",
		},
		"status": &graphql.ArgumentConfig{
			Type:        graphql.String,
			Description: "This is user's status",
		},
		"createdBy": &graphql.ArgumentConfig{
			Type:        graphql.String,
			Description: "This is user's createdBy",
		},
		"updatedBy": &graphql.ArgumentConfig{
			Type:        graphql.String,
			Description: "This is user's updatedBy",
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
	GetByID(ctx context.Context, params arguments.UserGetByIDArgs) (models.User, error)
	GetByIDs(ctx context.Context, params arguments.UserGetByIDsArgs) ([]models.User, error)
	Count(ctx context.Context, params arguments.UserCountArgs) (int64, error)
	List(ctx context.Context, params arguments.UserListArgs) ([]models.User, error)
	Insert(ctx context.Context, params arguments.UserInsertArgs) (models.User, error)
	Update(ctx context.Context, params arguments.UserUpdateArgs) (models.User, error)
	Delete(ctx context.Context, params arguments.UserDeleteArgs) (int64, error)
}

// ForwardParams ...
func (r ResolverImpl) ForwardParams(params graphql.ResolveParams) (interface{}, error) {
	return params.Args, nil
}

// GetByID ...
func (r ResolverImpl) GetByID(params graphql.ResolveParams) (interface{}, error) {
	// parse params
	args := arguments.UserGetByIDArgs{}
	if err := utils.Parse(params.Args, &args); err != nil {
		return nil, err
	}
	response, err := r.user.GetByID(params.Context, args)
	if err != nil {
		return nil, err
	}
	return response, nil
}

// Count ...
func (r ResolverImpl) Count(params graphql.ResolveParams) (interface{}, error) {
	// parse params
	args := arguments.UserCountArgs{}
	err := utils.Parse(params.Source.(map[string]interface{}), &args)
	if err != nil {
		return nil, err
	}
	response, err := r.user.Count(params.Context, args)
	if err != nil {
		return nil, err
	}
	return response, nil
}

// List ...
func (r ResolverImpl) List(params graphql.ResolveParams) (interface{}, error) {
	// parse params
	args := arguments.UserListArgs{}
	err := utils.Parse(params.Source.(map[string]interface{}), &args)
	if err != nil {
		return nil, err
	}
	response, err := r.user.List(params.Context, args)
	if err != nil {
		return nil, err
	}
	return response, nil
}

// Insert ...
func (r ResolverImpl) Insert(params graphql.ResolveParams) (interface{}, error) {
	// parse params
	args := arguments.UserInsertArgs{}
	err := utils.Parse(params.Args, &args)
	if err != nil {
		return nil, err
	}
	response, err := r.user.Insert(params.Context, args)
	if err != nil {
		return nil, err
	}
	return response, nil
}

// Update ...
func (r ResolverImpl) Update(params graphql.ResolveParams) (interface{}, error) {
	// parse params
	args := arguments.UserUpdateArgs{}
	err := utils.Parse(params.Args, &args)
	if err != nil {
		return nil, err
	}
	response, err := r.user.Update(params.Context, args)
	if err != nil {
		return nil, err
	}
	return response, nil
}

// Delete ...
func (r ResolverImpl) Delete(params graphql.ResolveParams) (interface{}, error) {
	// parse params
	args := arguments.UserDeleteArgs{}
	err := utils.Parse(params.Args, &args)
	if err != nil {
		return nil, err
	}
	response, err := r.user.Delete(params.Context, args)
	if err != nil {
		return nil, err
	}
	return response, nil
}

//go:generate mockery -name=IHandler -output=mocks -case=underscore
