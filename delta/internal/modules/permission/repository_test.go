package permission

import (
	"testing"

	"github.com/stretchr/testify/suite"
	"github.com/tanphamhaiduong/go/common/logger"
	mocksDB "github.com/tanphamhaiduong/go/delta/internal/database/mocks"
	"github.com/tanphamhaiduong/go/delta/internal/modules/permission/mocks"
)

type PermissionRepositoryTestSuite struct {
	suite.Suite
	MockIDB         *mocksDB.IDB
	MockIStmt       *mocksDB.IStmt
	MockIRow        *mocksDB.IRow
	MockIRows       *mocksDB.IRows
	MockIResult     *mocksDB.IResult
	MockIPermission *mocks.IRepository
	Repository      RepositoryImpl
}

func TestPermissionRepositoryTestSuite(t *testing.T) {
	suite.Run(t, new(PermissionRepositoryTestSuite))
}

func (s *PermissionRepositoryTestSuite) SetupTest() {
	mockIDB := &mocksDB.IDB{}
	s.MockIDB = mockIDB
	s.MockIStmt = &mocksDB.IStmt{}
	s.MockIRow = &mocksDB.IRow{}
	s.MockIRows = &mocksDB.IRows{}
	s.MockIResult = &mocksDB.IResult{}
	s.MockIPermission = &mocks.IRepository{}
	repository := NewRepository(mockIDB)
	s.Repository = *repository
	logConfig := logger.Configuration{}
	logger.NewLogger(logConfig, logger.InstanceZapLogger)
}
