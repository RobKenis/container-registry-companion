package index_handler

import (
	"html/template"
	"net/http"

	"github.com/rs/zerolog/log"
)

const templ = `<!DOCTYPE html>
<html>

<head>
    <title>Container Registry Companion</title>

    <link rel="stylesheet" href="/static/css/main.css">

    <script src="https://unpkg.com/htmx.org@2.0.2"></script>

    <link href="https://cdn.jsdelivr.net/npm/bootstrap@5.3.3/dist/css/bootstrap.min.css" rel="stylesheet" integrity="sha384-QWTKZyjpPEjISv5WaRU9OFeRpok6YctnYmDr5pNlyT2bRjXh0JMhjY6hW+ALEwIH" crossorigin="anonymous">
    <script src="https://cdn.jsdelivr.net/npm/bootstrap@5.3.3/dist/js/bootstrap.bundle.min.js" integrity="sha384-YvpcrYf0tY3lHB60NNkmXc5s9fDVZLESaAA55NDzOxhy9GkcIdslK1eN7N6jIeHz" crossorigin="anonymous"></script>
</head>

<body>
    <div class="container">
    <h1>Container Registry Companion</h1>

    <button hx-get="/repositories" hx-target="#repositories" class="btn btn-outline-primary">
        Load repositories
    </button>
    <div id="repositories" hx-get="/repositories" hx-trigger="load" />
    </div>
</body>

</html>
`

type Handler struct {
}

func (h Handler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	t, err := template.New("index").Parse(templ)
	if err != nil {
		log.Error().Err(err).Msg("Failed to parse template")
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)

	_ = t.Execute(w, nil)
}
