// @generated
package rolepermission

import (
	"github.com/tanphamhaiduong/go/delta/internal/models"
)

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
