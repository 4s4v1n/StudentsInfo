package v1

import (
	"context"
	"encoding/json"
	"github.com/sav1nbrave4code/APG3/internal/entity"
	"net/http"
	"path"
)

func (h *handler) AddCheck(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	check := entity.Check{}
	if err := json.NewDecoder(r.Body).Decode(&check); err != nil {
		WriteError(w, err, http.StatusBadRequest)
		return
	}

	ctx := context.Background()

	if err := h.tableService.AddCheck(ctx, check); err != nil {
		WriteError(w, err, http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
}

func (h *handler) Checks(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	ctx := context.Background()

	checks, err := h.tableService.GetChecks(ctx)
	if err != nil {
		WriteError(w, err, http.StatusInternalServerError)
		return
	}

	WriteResponseJson(w, checks)
}

func (h *handler) UpdateCheck(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	check := entity.Check{}
	if err := json.NewDecoder(r.Body).Decode(&check); err != nil {
		WriteError(w, err, http.StatusBadRequest)
		return
	}

	ctx := context.Background()

	if err := h.tableService.UpdateCheck(ctx, check); err != nil {
		WriteError(w, err, http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
}

func (h *handler) DeleteCheck(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	ctx := context.Background()

	if err := h.tableService.DeleteCheck(ctx, path.Base(r.URL.Path)); err != nil {
		WriteError(w, err, http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
}
