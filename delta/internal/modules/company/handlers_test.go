package company

import (
	"testing"

	"github.com/stretchr/testify/suite"
	"github.com/tanphamhaiduong/go/common/logger"
	"github.com/tanphamhaiduong/go/delta/internal/modules/company/mocks"
)

type CompanyHandlerTestSuite struct {
	suite.Suite
	MockICompany *mocks.IRepository
	Company      HandlerImpl
}

func TestCompanyHandlerTestSuite(t *testing.T) {
	suite.Run(t, new(CompanyHandlerTestSuite))
}

func (s *CompanyHandlerTestSuite) SetupTest() {
	s.MockICompany = &mocks.IRepository{}
	handler := NewHandler(s.MockICompany)
	s.Company = *handler
	s.Company.company = s.MockICompany
	logConfig := logger.Configuration{}
	logger.NewLogger(logConfig, logger.InstanceZapLogger)
}
