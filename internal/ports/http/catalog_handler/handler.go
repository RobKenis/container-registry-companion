package catalog_handler

import (
	"html/template"
	"net/http"

	"github.com/robkenis/container-registry-companion/internal/catalog"
	"github.com/rs/zerolog/log"
)

const templ = `
<table class="table table-hover">
  <thead>
    <tr>
      <th scope="col">#</th>
      <th scope="col">Name</th>
      <th scope="col">Updated</th>
    </tr>
  </thead>
  <tbody>
{{ range $idx, $repository := . }}
    <tr>
      <th scope="row">{{ $idx }}</th>
      <td>{{ .Name }}</td>
      <td>{{ .LastUpdated.Format "Jan 02, 2006 15:04:05 UTC" }}</td>
    </tr>
{{end}}
  </tbody>
</table>
`

type Handler struct {
	Catalog catalog.Catalog
}

func (h Handler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
	repositories, err := h.Catalog.List()
	if err != nil {
		log.Error().Err(err).Msg("Failed to list repositories")
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	t, err := template.New("catalog").Parse(templ)
	if err != nil {
		log.Error().Err(err).Msg("Failed to parse template")
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	_ = t.Execute(w, repositories)
}
