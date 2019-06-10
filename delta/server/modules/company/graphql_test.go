package company

import (
	"testing"

	"github.com/stretchr/testify/suite"
	mocksDB "github.com/tanphamhaiduong/go/delta/server/database/mocks"
	"github.com/tanphamhaiduong/go/delta/server/modules/company/mocks"
)

type CompanyResolverTestSuite struct {
	suite.Suite
	MockIDB                *mocksDB.IDB
	MockIStmt              *mocksDB.IStmt
	MockIRow               *mocksDB.IRow
	MockIRows              *mocksDB.IRows
	MockIResult            *mocksDB.IResult
	MockIHandlerCompany    *mocks.IHandler
	MockIRepositoryCompany *mocks.IRepository
	Company                ResolverImpl
}

func TestCompanyResolverTestSuite(t *testing.T) {
	suite.Run(t, new(CompanyResolverTestSuite))
}

func (s *CompanyResolverTestSuite) SetupTest() {
	mockIDB := &mocksDB.IDB{}
	s.MockIDB = mockIDB
	s.MockIStmt = &mocksDB.IStmt{}
	s.MockIRow = &mocksDB.IRow{}
	s.MockIRows = &mocksDB.IRows{}
	s.MockIResult = &mocksDB.IResult{}
	s.MockIHandlerCompany = &mocks.IHandler{}
	s.MockIRepositoryCompany = &mocks.IRepository{}
	s.Company = NewResolver(mockIDB)
}
