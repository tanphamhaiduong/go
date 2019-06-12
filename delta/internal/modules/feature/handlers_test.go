package feature

import (
	"testing"

	"github.com/stretchr/testify/suite"
	"github.com/tanphamhaiduong/go/delta/internal/modules/feature/mocks"
)

type FeatureHandlerTestSuite struct {
	suite.Suite
	MockIFeature *mocks.IRepository
	Feature      HandlerImpl
}

func TestFeatureHandlerTestSuite(t *testing.T) {
	suite.Run(t, new(FeatureHandlerTestSuite))
}

func (s *FeatureHandlerTestSuite) SetupTest() {
	s.MockIFeature = &mocks.IRepository{}
	s.Feature = NewHandler(s.MockIFeature)
	s.Feature.feature = s.MockIFeature
}
