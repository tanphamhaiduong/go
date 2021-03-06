package company

import (
	"testing"

	"github.com/stretchr/testify/suite"
	"github.com/tanphamhaiduong/go/common/logger"
	"github.com/tanphamhaiduong/go/delta/internal/modules/company/mocks"
)

type CompanyResolverTestSuite struct {
	suite.Suite
	MockICompany *mocks.IHandler
	Company      ResolverImpl
}

func TestCompanyResolverTestSuite(t *testing.T) {
	suite.Run(t, new(CompanyResolverTestSuite))
}

func (s *CompanyResolverTestSuite) SetupTest() {
	s.MockICompany = &mocks.IHandler{}
	resolver := NewResolver(s.MockICompany)
	s.Company = *resolver
	s.Company.company = s.MockICompany
	logConfig := logger.Configuration{}
	logger.NewLogger(logConfig, logger.InstanceZapLogger)
}
