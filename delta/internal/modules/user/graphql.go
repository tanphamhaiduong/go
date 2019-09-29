package user

import (
	"context"
	"database/sql"

	"github.com/graphql-go/graphql"
	"github.com/tanphamhaiduong/go/common/goerrors"
	"github.com/tanphamhaiduong/go/common/logger"
	"github.com/tanphamhaiduong/go/delta/internal/arguments"
	"github.com/tanphamhaiduong/go/delta/internal/utils"
)

var (
	// LoginTypeArgs ...
	LoginTypeArgs = graphql.FieldConfigArgument{
		"username": &graphql.ArgumentConfig{
			Type: graphql.NewNonNull(graphql.String),
		},
		"password": &graphql.ArgumentConfig{
			Type: graphql.NewNonNull(graphql.String),
		},
	}
)

// IHandler ...
type IHandler interface {
	ICoreHandler
	Login(ctx context.Context, param arguments.UserLogin) (string, error)
}

// ResolverImpl ...
type ResolverImpl struct {
	user IHandler
}

// NewResolver ...
func NewResolver(user IHandler) *ResolverImpl {
	return &ResolverImpl{
		user: user,
	}
}

// Login ...
func (r *ResolverImpl) Login(param graphql.ResolveParams) (interface{}, error) {
	logger.WithFields(logger.Fields{
		"traceId": param.Context.Value(utils.TraceIDKey),
		"param":   param,
	}).Infof("Resolver Login of user")
	// parse param
	args := arguments.UserLogin{}
	if err := utils.Parse(param.Args, &args); err != nil {
		logger.WithFields(logger.Fields{
			"traceId": param.Context.Value(utils.TraceIDKey),
			"param":   param,
			"err":     err,
		}).Infof("utils.Parse of Login Resolver")
		return nil, err
	}
	response, err := r.user.Login(param.Context, args)
	if err != nil {
		logger.WithFields(logger.Fields{
			"traceId": param.Context.Value(utils.TraceIDKey),
			"param":   param,
			"err":     err,
		}).Infof("r.user.Login of Login Resolver")
		if err == sql.ErrNoRows {
			return nil, goerrors.ErrNotFound(param.Context.Value(utils.TraceIDKey))
		}
		return nil, goerrors.ErrInternalServerError(param.Context.Value(utils.TraceIDKey))
	}
	return response, nil
}
