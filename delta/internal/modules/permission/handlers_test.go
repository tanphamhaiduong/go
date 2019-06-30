package permission

import (
	"testing"

	"github.com/stretchr/testify/suite"
	"github.com/tanphamhaiduong/go/delta/internal/modules/permission/mocks"
)

type PermissionHandlerTestSuite struct {
	suite.Suite
	MockIPermission *mocks.IRepository
	Permission      HandlerImpl
}

func TestPermissionHandlerTestSuite(t *testing.T) {
	suite.Run(t, new(PermissionHandlerTestSuite))
}

func (s *PermissionHandlerTestSuite) SetupTest() {
	s.MockIPermission = &mocks.IRepository{}
	handler := NewHandler(s.MockIPermission)
	s.Permission = *handler
	s.Permission.permission = s.MockIPermission
}
