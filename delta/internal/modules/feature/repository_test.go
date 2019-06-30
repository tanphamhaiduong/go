package feature

import (
	"testing"

	"github.com/stretchr/testify/suite"
	mocksDB "github.com/tanphamhaiduong/go/delta/internal/database/mocks"
	"github.com/tanphamhaiduong/go/delta/internal/modules/feature/mocks"
)

type FeatureRepositoryTestSuite struct {
	suite.Suite
	MockIDB      *mocksDB.IDB
	MockIStmt    *mocksDB.IStmt
	MockIRow     *mocksDB.IRow
	MockIRows    *mocksDB.IRows
	MockIResult  *mocksDB.IResult
	MockIFeature *mocks.IRepository
	Repository   RepositoryImpl
}

func TestFeatureRepositoryTestSuite(t *testing.T) {
	suite.Run(t, new(FeatureRepositoryTestSuite))
}

func (s *FeatureRepositoryTestSuite) SetupTest() {
	mockIDB := &mocksDB.IDB{}
	s.MockIDB = mockIDB
	s.MockIStmt = &mocksDB.IStmt{}
	s.MockIRow = &mocksDB.IRow{}
	s.MockIRows = &mocksDB.IRows{}
	s.MockIResult = &mocksDB.IResult{}
	s.MockIFeature = &mocks.IRepository{}
	repository := NewRepository(mockIDB)
	s.Repository = *repository
}
