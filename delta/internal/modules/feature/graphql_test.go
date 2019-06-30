package feature

import (
	"testing"

	"github.com/stretchr/testify/suite"
	"github.com/tanphamhaiduong/go/delta/internal/modules/feature/mocks"
)

type FeatureResolverTestSuite struct {
	suite.Suite
	MockIFeature *mocks.IHandler
	Feature      ResolverImpl
}

func TestFeatureResolverTestSuite(t *testing.T) {
	suite.Run(t, new(FeatureResolverTestSuite))
}

func (s *FeatureResolverTestSuite) SetupTest() {
	s.MockIFeature = &mocks.IHandler{}
	resolver := NewResolver(s.MockIFeature)
	s.Feature = *resolver
	s.Feature.feature = s.MockIFeature
}
