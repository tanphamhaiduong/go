package role

import (
	"testing"

	"github.com/stretchr/testify/suite"
	"github.com/tanphamhaiduong/go/delta/internal/modules/role/mocks"
)

type RoleHandlerTestSuite struct {
	suite.Suite
	MockIRole *mocks.IRepository
	Role      HandlerImpl
}

func TestRoleHandlerTestSuite(t *testing.T) {
	suite.Run(t, new(RoleHandlerTestSuite))
}

func (s *RoleHandlerTestSuite) SetupTest() {
	s.MockIRole = &mocks.IRepository{}
	handler := NewHandler(s.MockIRole)
	s.Role = *handler
	s.Role.role = s.MockIRole
}
