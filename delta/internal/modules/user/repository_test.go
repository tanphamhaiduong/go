package user

import (
	"context"
	"errors"
	"testing"

	"github.com/stretchr/testify/mock"
	"github.com/stretchr/testify/suite"
	"github.com/tanphamhaiduong/go/common/logger"
	mocksDB "github.com/tanphamhaiduong/go/delta/internal/database/mocks"
	"github.com/tanphamhaiduong/go/delta/internal/models"
	"github.com/tanphamhaiduong/go/delta/internal/modules/user/mocks"
)

type UserRepositoryTestSuite struct {
	suite.Suite
	MockIDB     *mocksDB.IDB
	MockIStmt   *mocksDB.IStmt
	MockIRow    *mocksDB.IRow
	MockIRows   *mocksDB.IRows
	MockIResult *mocksDB.IResult
	MockIUser   *mocks.IRepository
	Repository  RepositoryImpl
}

func TestUserRepositoryTestSuite(t *testing.T) {
	suite.Run(t, new(UserRepositoryTestSuite))
}

func (s *UserRepositoryTestSuite) SetupTest() {
	mockIDB := &mocksDB.IDB{}
	s.MockIDB = mockIDB
	s.MockIStmt = &mocksDB.IStmt{}
	s.MockIRow = &mocksDB.IRow{}
	s.MockIRows = &mocksDB.IRows{}
	s.MockIResult = &mocksDB.IResult{}
	s.MockIUser = &mocks.IRepository{}
	repository := NewRepository(mockIDB)
	s.Repository = *repository
	logConfig := logger.Configuration{}
	logger.NewLogger(logConfig, logger.InstanceZapLogger)
}

func (s *UserRepositoryTestSuite) TestGetByUsername_Success() {
	var (
		ctx      = context.Background()
		username = "developer"
		status   = "active"
		user     models.User
	)
	s.MockIDB.On("PrepareContext", ctx, mock.Anything).Return(s.MockIStmt, nil)
	s.MockIStmt.On("QueryRowContext", ctx, username, status).Return(s.MockIRow)
	s.MockIRow.On("Scan",
		mock.Anything,
		mock.Anything,
		mock.Anything,
		mock.Anything,
		mock.Anything,
		mock.Anything,
		mock.Anything,
		mock.Anything,
		mock.Anything,
		mock.Anything,
		mock.Anything,
		mock.Anything,
		mock.Anything,
		mock.Anything,
		mock.Anything,
		mock.Anything,
		mock.Anything,
		mock.Anything,
	).Return(nil)
	actual, err := s.Repository.GetByUsername(ctx, username)
	s.Nil(err)
	s.Equal(user, actual)
}

func (s *UserRepositoryTestSuite) TestGetByUsername_Fail() {
	var (
		ctx      = context.Background()
		username = "developer"
		status   = "active"
		user     models.User
	)
	s.MockIDB.On("PrepareContext", ctx, mock.Anything).Return(s.MockIStmt, errors.New("some errors"))
	s.MockIStmt.On("QueryRowContext", ctx, username, status).Return(s.MockIRow)
	s.MockIRow.On("Scan",
		mock.Anything,
		mock.Anything,
		mock.Anything,
		mock.Anything,
		mock.Anything,
		mock.Anything,
		mock.Anything,
		mock.Anything,
		mock.Anything,
		mock.Anything,
		mock.Anything,
		mock.Anything,
		mock.Anything,
		mock.Anything,
		mock.Anything,
		mock.Anything,
		mock.Anything,
		mock.Anything,
	).Return(errors.New("some errors"))
	actual, err := s.Repository.GetByUsername(ctx, username)
	s.Equal(user, actual)
	s.NotNil(err)
}

func (s *UserRepositoryTestSuite) TestGetByUsername_Fail1() {
	var (
		ctx      = context.Background()
		username = "developer"
		status   = "active"
		user     models.User
	)
	s.MockIDB.On("PrepareContext", ctx, mock.Anything).Return(s.MockIStmt, nil)
	s.MockIStmt.On("QueryRowContext", ctx, username, status).Return(s.MockIRow)
	s.MockIRow.On("Scan",
		mock.Anything,
		mock.Anything,
		mock.Anything,
		mock.Anything,
		mock.Anything,
		mock.Anything,
		mock.Anything,
		mock.Anything,
		mock.Anything,
		mock.Anything,
		mock.Anything,
		mock.Anything,
		mock.Anything,
		mock.Anything,
		mock.Anything,
		mock.Anything,
		mock.Anything,
		mock.Anything,
	).Return(errors.New("some errors"))
	actual, err := s.Repository.GetByUsername(ctx, username)
	s.Equal(user, actual)
	s.NotNil(err)
}
