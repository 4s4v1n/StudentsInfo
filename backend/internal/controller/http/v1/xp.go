package v1

import (
	"context"
	"encoding/json"
	"github.com/sav1nbrave4code/APG3/internal/entity"
	"net/http"
	"path"
)

func (h *handler) AddXP(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	xp := entity.XP{}
	if err := json.NewDecoder(r.Body).Decode(&xp); err != nil {
		WriteError(w, err, http.StatusBadRequest)
		return
	}

	ctx := context.Background()

	if err := h.tableService.AddXP(ctx, xp); err != nil {
		WriteError(w, err, http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
}

func (h *handler) XPs(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	ctx := context.Background()

	xps, err := h.tableService.GetXPs(ctx)
	if err != nil {
		WriteError(w, err, http.StatusInternalServerError)
		return
	}

	WriteResponseJson(w, xps)
}

func (h *handler) UpdateXP(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	xp := entity.XP{}
	if err := json.NewDecoder(r.Body).Decode(&xp); err != nil {
		WriteError(w, err, http.StatusBadRequest)
		return
	}

	ctx := context.Background()

	if err := h.tableService.UpdateXP(ctx, xp); err != nil {
		WriteError(w, err, http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
}

func (h *handler) DeleteXP(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	ctx := context.Background()

	if err := h.tableService.DeleteXP(ctx, path.Base(r.URL.Path)); err != nil {
		WriteError(w, err, http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
}
