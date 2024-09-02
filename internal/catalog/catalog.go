package catalog

import (
	"encoding/json"
	"net/http"

	"github.com/robkenis/container-registry-companion/internal/repository"
	"github.com/rs/zerolog/log"
)

type Catalog interface {
	List() ([]repository.Repository, error)
}

type instantLoadingCatalog struct {
	RegistryUrl string
}

func NewCatalog(registry string) Catalog {
	return &instantLoadingCatalog{
		RegistryUrl: registry,
	}
}

type catalogResponse struct {
	Repositories []string `json:"repositories"`
}

func (c *instantLoadingCatalog) List() ([]repository.Repository, error) {
	catalogEndpoint := c.RegistryUrl + "/v2/_catalog"
	log.Debug().Msgf("Fetching catalog from %s", catalogEndpoint)
	resp, err := http.Get(catalogEndpoint)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	var catalogResponse catalogResponse
	json.NewDecoder(resp.Body).Decode(&catalogResponse)
	repositories := make([]repository.Repository, len(catalogResponse.Repositories))
	for i, name := range catalogResponse.Repositories {
		repositories[i] = repository.Repository{Name: name}
	}
	return repositories, nil
}
