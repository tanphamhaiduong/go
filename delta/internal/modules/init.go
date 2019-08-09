package modules

import (
	"github.com/graphql-go/graphql"
	"github.com/tanphamhaiduong/go/delta/internal/modules/user"
)

func addCustomToSchema(resolver Resolver) {
	rootMutation.AddFieldConfig("login", &graphql.Field{
		Type:        graphql.NewNonNull(graphql.String),
		Description: "This is mutation for login",
		Args:        user.LoginTypeArgs,
		Resolve:     resolver.User.Login,
	})
}
