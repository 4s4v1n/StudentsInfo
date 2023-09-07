package v1

import (
	"context"
	"encoding/json"
	"io"
	"net/http"
	"path"
)

func (h *handler) Import(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	ctx := context.Background()

	if err := h.dataService.Import(ctx, path.Base(r.URL.Path), r.Body); err != nil {
		WriteError(w, err, http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
}

func (h *handler) Export(w http.ResponseWriter, r *http.Request) {
	ctx := context.Background()

	out, err := h.dataService.Export(ctx, path.Base(r.URL.Path))
	if err != nil {
		WriteError(w, err, http.StatusInternalServerError)
		return
	}

	WriteResponseCsv(w, out)
}

func (h *handler) RawQuery(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	ctx := context.Background()

	data := struct {
		Expression string `json:"expression"`
	}{}
	jsonBody, err := io.ReadAll(r.Body)
	if err != nil {
		WriteError(w, err, http.StatusInternalServerError)
	}

	if err = json.Unmarshal(jsonBody, &data); err != nil {
		WriteError(w, err, http.StatusBadRequest)
	}

	raw, err := h.dataService.RawQuery(ctx, data.Expression)
	if err != nil {
		WriteError(w, err, http.StatusInternalServerError)
		return
	}

	WriteResponseJson(w, raw)
}
