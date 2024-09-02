package index_handler

import (
	"html/template"
	"net/http"

	"github.com/rs/zerolog/log"
)

type Handler struct {
}

func (h Handler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	t, err := template.New("index").ParseFiles("internal/ports/http/index_handler/index.html.tpl")
	if err != nil {
		log.Error().Err(err).Msg("Failed to parse template")
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)

	_ = t.ExecuteTemplate(w, "index.html.tpl", nil)
}
