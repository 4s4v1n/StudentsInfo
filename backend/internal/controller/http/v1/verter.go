package v1

import (
	"context"
	"encoding/json"
	"github.com/sav1nbrave4code/APG3/internal/entity"
	"github.com/sav1nbrave4code/APG3/internal/entity/service_types"
	"net/http"
	"path"
)

func (h *handler) AddVerter(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	verter := service_types.VerterInsert{}
	if err := json.NewDecoder(r.Body).Decode(&verter); err != nil {
		WriteError(w, err, http.StatusBadRequest)
		return
	}

	ctx := context.Background()

	if err := h.tableService.AddVerter(ctx, verter); err != nil {
		WriteError(w, err, http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
}

func (h *handler) Verters(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	ctx := context.Background()

	verters, err := h.tableService.GetVerters(ctx)
	if err != nil {
		WriteError(w, err, http.StatusInternalServerError)
		return
	}

	WriteResponseJson(w, verters)
}

func (h *handler) UpdateVerter(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	verter := entity.Verter{}
	if err := json.NewDecoder(r.Body).Decode(&verter); err != nil {
		WriteError(w, err, http.StatusBadRequest)
		return
	}

	ctx := context.Background()

	if err := h.tableService.UpdateVerter(ctx, verter); err != nil {
		WriteError(w, err, http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
}

func (h *handler) DeleteVerter(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	ctx := context.Background()

	if err := h.tableService.DeleteVerter(ctx, path.Base(r.URL.Path)); err != nil {
		WriteError(w, err, http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
}
