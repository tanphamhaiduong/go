package rolepermission

import (
	"testing"

	"github.com/stretchr/testify/suite"
	"github.com/tanphamhaiduong/go/delta/internal/modules/rolepermission/mocks"
)

type RolePermissionHandlerTestSuite struct {
	suite.Suite
	MockIRolePermission *mocks.IRepository
	RolePermission      HandlerImpl
}

func TestRolePermissionHandlerTestSuite(t *testing.T) {
	suite.Run(t, new(RolePermissionHandlerTestSuite))
}

func (s *RolePermissionHandlerTestSuite) SetupTest() {
	s.MockIRolePermission = &mocks.IRepository{}
	handler := NewHandler(s.MockIRolePermission)
	s.RolePermission = *handler
	s.RolePermission.rolepermission = s.MockIRolePermission
}
