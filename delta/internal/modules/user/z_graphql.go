// @generated
package user

import (
	"context"

	"github.com/graphql-go/graphql"
	log "github.com/sirupsen/logrus"
	"github.com/tanphamhaiduong/go/delta/internal/arguments"
	"github.com/tanphamhaiduong/go/delta/internal/customgraphql"
	"github.com/tanphamhaiduong/go/delta/internal/models"
	"github.com/tanphamhaiduong/go/delta/internal/utils"
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
			"username": &graphql.Field{
				Type:        graphql.NewNonNull(graphql.String),
				Description: "This is user's username",
			},
			"password": &graphql.Field{
				Type:        graphql.NewNonNull(graphql.String),
				Description: "This is user's password",
			},
			"name": &graphql.Field{
				Type:        graphql.NewNonNull(graphql.String),
				Description: "This is user's name",
			},
			"dateOfBirth": &graphql.Field{
				Type:        customgraphql.NullDateTime,
				Description: "This is user's name",
			},
			"reference": &graphql.Field{
				Type:        graphql.NewNonNull(graphql.String),
				Description: "This is user's reference",
			},
			"avatarUrl": &graphql.Field{
				Type:        graphql.NewNonNull(graphql.String),
				Description: "This is user's avatar url",
			},
			"licenseNumber": &graphql.Field{
				Type:        graphql.NewNonNull(graphql.String),
				Description: "This is user's license number",
			},
			"phoneNumber": &graphql.Field{
				Type:        graphql.NewNonNull(graphql.String),
				Description: "This is user's name",
			},
			"extension": &graphql.Field{
				Type:        graphql.NewNonNull(graphql.String),
				Description: "This is user's extension",
			},
			"telProvider": &graphql.Field{
				Type:        graphql.NewNonNull(graphql.String),
				Description: "This is user's tel provider",
			},
			"telApi": &graphql.Field{
				Type:        graphql.NewNonNull(graphql.String),
				Description: "This is user's tel api",
			},
			"supervisorId": &graphql.Field{
				Type:        graphql.NewNonNull(graphql.Int),
				Description: "This is user's tel supervisorId",
			},
			"roleId": &graphql.Field{
				Type:        graphql.NewNonNull(graphql.Int),
				Description: "This is user's tel roleId",
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
		"username": &graphql.ArgumentConfig{
			Type:        graphql.String,
			Description: "This is user's username",
		},
		"password": &graphql.ArgumentConfig{
			Type:        graphql.String,
			Description: "This is user's password",
		},
		"name": &graphql.ArgumentConfig{
			Type:        graphql.String,
			Description: "This is user's name",
		},
		"dateOfBirth": &graphql.ArgumentConfig{
			Type:        customgraphql.NullDateTime,
			Description: "This is user's name",
		},
		"reference": &graphql.ArgumentConfig{
			Type:        graphql.String,
			Description: "This is user's reference",
		},
		"avatarUrl": &graphql.ArgumentConfig{
			Type:        graphql.String,
			Description: "This is user's avatar url",
		},
		"licenseNumber": &graphql.ArgumentConfig{
			Type:        graphql.String,
			Description: "This is user's license number",
		},
		"phoneNumber": &graphql.ArgumentConfig{
			Type:        graphql.String,
			Description: "This is user's name",
		},
		"extension": &graphql.ArgumentConfig{
			Type:        graphql.String,
			Description: "This is user's extension",
		},
		"telProvider": &graphql.ArgumentConfig{
			Type:        graphql.String,
			Description: "This is user's tel provider",
		},
		"telApi": &graphql.ArgumentConfig{
			Type:        graphql.String,
			Description: "This is user's tel api",
		},
		"supervisorId": &graphql.ArgumentConfig{
			Type:        graphql.Int,
			Description: "This is user's tel supervisorId",
		},
		"roleId": &graphql.ArgumentConfig{
			Type:        graphql.Int,
			Description: "This is user's tel roleId",
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
		"username": &graphql.ArgumentConfig{
			Type:        graphql.NewNonNull(graphql.String),
			Description: "This is user's username",
		},
		"password": &graphql.ArgumentConfig{
			Type:        graphql.NewNonNull(graphql.String),
			Description: "This is user's password",
		},
		"name": &graphql.ArgumentConfig{
			Type:        graphql.NewNonNull(graphql.String),
			Description: "This is user's name",
		},
		"dateOfBirth": &graphql.ArgumentConfig{
			Type:        customgraphql.NullDateTime,
			Description: "This is user's name",
		},
		"reference": &graphql.ArgumentConfig{
			Type:        graphql.NewNonNull(graphql.String),
			Description: "This is user's reference",
		},
		"avatarUrl": &graphql.ArgumentConfig{
			Type:        graphql.NewNonNull(graphql.String),
			Description: "This is user's avatar url",
		},
		"licenseNumber": &graphql.ArgumentConfig{
			Type:        graphql.NewNonNull(graphql.String),
			Description: "This is user's license number",
		},
		"phoneNumber": &graphql.ArgumentConfig{
			Type:        graphql.NewNonNull(graphql.String),
			Description: "This is user's name",
		},
		"extension": &graphql.ArgumentConfig{
			Type:        graphql.NewNonNull(graphql.String),
			Description: "This is user's extension",
		},
		"telProvider": &graphql.ArgumentConfig{
			Type:        graphql.NewNonNull(graphql.String),
			Description: "This is user's tel provider",
		},
		"telApi": &graphql.ArgumentConfig{
			Type:        graphql.NewNonNull(graphql.String),
			Description: "This is user's tel api",
		},
		"supervisorId": &graphql.ArgumentConfig{
			Type:        graphql.NewNonNull(graphql.Int),
			Description: "This is user's tel supervisorId",
		},
		"roleId": &graphql.ArgumentConfig{
			Type:        graphql.NewNonNull(graphql.Int),
			Description: "This is user's tel roleId",
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
		"username": &graphql.ArgumentConfig{
			Type:        graphql.String,
			Description: "This is user's username",
		},
		"password": &graphql.ArgumentConfig{
			Type:        graphql.String,
			Description: "This is user's password",
		},
		"name": &graphql.ArgumentConfig{
			Type:        graphql.String,
			Description: "This is user's name",
		},
		"dateOfBirth": &graphql.ArgumentConfig{
			Type:        customgraphql.NullDateTime,
			Description: "This is user's name",
		},
		"reference": &graphql.ArgumentConfig{
			Type:        graphql.String,
			Description: "This is user's reference",
		},
		"avatarUrl": &graphql.ArgumentConfig{
			Type:        graphql.String,
			Description: "This is user's avatar url",
		},
		"licenseNumber": &graphql.ArgumentConfig{
			Type:        graphql.String,
			Description: "This is user's license number",
		},
		"phoneNumber": &graphql.ArgumentConfig{
			Type:        graphql.String,
			Description: "This is user's name",
		},
		"extension": &graphql.ArgumentConfig{
			Type:        graphql.String,
			Description: "This is user's extension",
		},
		"telProvider": &graphql.ArgumentConfig{
			Type:        graphql.String,
			Description: "This is user's tel provider",
		},
		"telApi": &graphql.ArgumentConfig{
			Type:        graphql.String,
			Description: "This is user's tel api",
		},
		"supervisorId": &graphql.ArgumentConfig{
			Type:        graphql.Int,
			Description: "This is user's tel supervisorId",
		},
		"roleId": &graphql.ArgumentConfig{
			Type:        graphql.Int,
			Description: "This is user's tel roleId",
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
func (r *ResolverImpl) ForwardParams(params graphql.ResolveParams) (interface{}, error) {
	log.WithField("params", params).Info("Resolver ForwardParams of user")
	return params.Args, nil
}

// GetByID ...
func (r *ResolverImpl) GetByID(params graphql.ResolveParams) (interface{}, error) {
	log.WithField("params", params).Info("Resolver GetByID of user")
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
func (r *ResolverImpl) Count(params graphql.ResolveParams) (interface{}, error) {
	log.WithField("params", params).Info("Resolver Count of user")
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
func (r *ResolverImpl) List(params graphql.ResolveParams) (interface{}, error) {
	log.WithField("params", params).Info("Resolver List of user")
	// parse params
	args := arguments.UserListArgs{}
	err := utils.Parse(params.Source.(map[string]interface{}), &args)
	if err != nil {
		log.WithField("Error", err).Error("Resolver List utils.Parse user")
		return nil, err
	}
	response, err := r.user.List(params.Context, args)
	if err != nil {
		log.WithField("Error", err).Error("Resolver List r.user.List user")
		return nil, err
	}
	return response, nil
}

// Insert ...
func (r *ResolverImpl) Insert(params graphql.ResolveParams) (interface{}, error) {
	log.WithField("params", params).Info("Resolver Insert of user")
	// parse params
	args := arguments.UserInsertArgs{}
	err := utils.Parse(params.Args, &args)
	if err != nil {
		log.WithField("Error", err).Error("Resolver Insert utils.Parse user")
		return nil, err
	}
	response, err := r.user.Insert(params.Context, args)
	if err != nil {
		log.WithField("Error", err).Error("Resolver Insert r.user.Insert user")
		return nil, err
	}
	return response, nil
}

// Update ...
func (r *ResolverImpl) Update(params graphql.ResolveParams) (interface{}, error) {
	log.WithField("params", params).Info("Resolver Update of user")
	// parse params
	args := arguments.UserUpdateArgs{}
	err := utils.Parse(params.Args, &args)
	if err != nil {
		log.WithField("Error", err).Error("Resolver Update utils.Parse user")
		return nil, err
	}
	response, err := r.user.Update(params.Context, args)
	if err != nil {
		log.WithField("Error", err).Error("Resolver Update r.user.Update user")
		return nil, err
	}
	return response, nil
}

// Delete ...
func (r *ResolverImpl) Delete(params graphql.ResolveParams) (interface{}, error) {
	log.WithField("params", params).Info("Resolver Delete of user")
	// parse params
	args := arguments.UserDeleteArgs{}
	err := utils.Parse(params.Args, &args)
	if err != nil {
		log.WithField("Error", err).Error("Resolver Delete utils.Parse user")
		return nil, err
	}
	response, err := r.user.Delete(params.Context, args)
	if err != nil {
		log.WithField("Error", err).Error("Resolver Delete r.user.Delete user")
		return nil, err
	}
	return response, nil
}

//go:generate mockery -name=IHandler -output=mocks -case=underscore
