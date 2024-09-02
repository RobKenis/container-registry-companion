package catalog_test

import (
	"errors"
	"testing"

	"github.com/jarcoal/httpmock"
	"github.com/robkenis/container-registry-companion/internal/catalog"
	"github.com/stretchr/testify/suite"
)

type CatalogSuite struct {
	suite.Suite
}

func TestCatalogSuite(t *testing.T) {
	suite.Run(t, new(CatalogSuite))
}

func (suite *CatalogSuite) TestList() {
	httpmock.Activate()
	defer httpmock.DeactivateAndReset()

	httpmock.RegisterResponder("GET", "http://mock.registry/v2/_catalog",
		httpmock.NewStringResponder(200, `{"repositories": ["nginx", "prometheus"]}`))

	catalog := catalog.NewCatalog("http://mock.registry")

	repositories, err := catalog.List()

	suite.Require().NoError(err)
	suite.Require().Len(repositories, 2)

	suite.Equal("nginx", repositories[0].Name)
	suite.Equal("prometheus", repositories[1].Name)
}

func (suite *CatalogSuite) TestList_error() {
	httpmock.Activate()
	defer httpmock.DeactivateAndReset()

	httpmock.RegisterResponder("GET", "http://mock.registry/v2/_catalog",
		httpmock.NewErrorResponder(errors.New("error")))

	catalog := catalog.NewCatalog("http://mock.registry")

	repositories, err := catalog.List()

	suite.Error(err)
	suite.Nil(repositories)
}
