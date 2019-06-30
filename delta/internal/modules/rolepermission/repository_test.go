package rolepermission

import (
	"testing"

	"github.com/stretchr/testify/suite"
	mocksDB "github.com/tanphamhaiduong/go/delta/internal/database/mocks"
	"github.com/tanphamhaiduong/go/delta/internal/modules/rolepermission/mocks"
)

type RolePermissionRepositoryTestSuite struct {
	suite.Suite
	MockIDB             *mocksDB.IDB
	MockIStmt           *mocksDB.IStmt
	MockIRow            *mocksDB.IRow
	MockIRows           *mocksDB.IRows
	MockIResult         *mocksDB.IResult
	MockIRolePermission *mocks.IRepository
	Repository          RepositoryImpl
}

func TestRolePermissionRepositoryTestSuite(t *testing.T) {
	suite.Run(t, new(RolePermissionRepositoryTestSuite))
}

func (s *RolePermissionRepositoryTestSuite) SetupTest() {
	mockIDB := &mocksDB.IDB{}
	s.MockIDB = mockIDB
	s.MockIStmt = &mocksDB.IStmt{}
	s.MockIRow = &mocksDB.IRow{}
	s.MockIRows = &mocksDB.IRows{}
	s.MockIResult = &mocksDB.IResult{}
	s.MockIRolePermission = &mocks.IRepository{}
	repository := NewRepository(mockIDB)
	s.Repository = *repository
}
