package user

import (
	"context"

	"github.com/graphql-go/graphql"
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
	Login(ctx context.Context, param arguments.UserLoginArgs) (string, error)
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
		"TraceID": param.Context.Value("TraceID"),
		"param":   param,
	}).Infof("Resolver Login of user")
	// parse param
	args := arguments.UserLoginArgs{}
	if err := utils.Parse(param.Args, &args); err != nil {
		logger.WithFields(logger.Fields{
			"TraceID": param.Context.Value("TraceID"),
			"param":   param,
			"err":     err,
		}).Infof("utils.Parse of Login Resolver")
		return nil, err
	}
	response, err := r.user.Login(param.Context, args)
	if err != nil {
		logger.WithFields(logger.Fields{
			"TraceID": param.Context.Value("TraceID"),
			"param":   param,
			"err":     err,
		}).Infof("r.user.Login of Login Resolver")
		return nil, err
	}
	return response, nil
}
