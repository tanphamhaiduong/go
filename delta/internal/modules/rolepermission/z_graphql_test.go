// @generated
package rolepermission

import (
	"context"

	"github.com/graphql-go/graphql"

	"github.com/tanphamhaiduong/go/delta/internal/models"
)

func (s *RolePermissionResolverTestSuite) TestForwardParams_Success() {
	var (
		sampleID int64 = 1
		params         = graphql.ResolveParams{Context: context.Background(), Source: map[string]interface{}{}, Args: map[string]interface{}{"id": sampleID}}
		expected       = map[string]interface{}(map[string]interface{}{"id": sampleID})
	)
	actual, err := s.RolePermission.ForwardParams(params)
	s.Nil(err)
	s.Equal(expected, actual)
}

func (s *RolePermissionResolverTestSuite) TestCheckPermission_True() {
	var (
		claims = &models.Claims{
			Permissions: []models.Permission{
				{
					ID:   1,
					Name: "RolePermissionGetByID",
				},
			},
		}
		expected = true
	)
	actual := s.RolePermission.checkPermission(*claims, "RolePermissionGetByID")
	s.Equal(expected, actual)
}

func (s *RolePermissionResolverTestSuite) TestCheckPermission_False() {
	var (
		claims = &models.Claims{
			Permissions: []models.Permission{
				{
					ID:   1,
					Name: "RolePermissionList",
				},
			},
		}
		expected = false
	)
	actual := s.RolePermission.checkPermission(*claims, "RolePermissionGetByID")
	s.Equal(expected, actual)
}
