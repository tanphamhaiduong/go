package role

import (
	"testing"

	"github.com/stretchr/testify/suite"
	mocksDB "github.com/tanphamhaiduong/go/delta/internal/database/mocks"
	"github.com/tanphamhaiduong/go/delta/internal/modules/role/mocks"
)

type RoleRepositoryTestSuite struct {
	suite.Suite
	MockIDB     *mocksDB.IDB
	MockIStmt   *mocksDB.IStmt
	MockIRow    *mocksDB.IRow
	MockIRows   *mocksDB.IRows
	MockIResult *mocksDB.IResult
	MockIRole   *mocks.IRepository
	Repository  RepositoryImpl
}

func TestRoleRepositoryTestSuite(t *testing.T) {
	suite.Run(t, new(RoleRepositoryTestSuite))
}

func (s *RoleRepositoryTestSuite) SetupTest() {
	mockIDB := &mocksDB.IDB{}
	s.MockIDB = mockIDB
	s.MockIStmt = &mocksDB.IStmt{}
	s.MockIRow = &mocksDB.IRow{}
	s.MockIRows = &mocksDB.IRows{}
	s.MockIResult = &mocksDB.IResult{}
	s.MockIRole = &mocks.IRepository{}
	repository := NewRepository(mockIDB)
	s.Repository = *repository
}
