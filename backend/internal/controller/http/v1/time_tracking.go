package v1

import (
	"context"
	"encoding/json"
	"net/http"
	"path"

	"github.com/sav1nbrave4code/APG3/internal/entity"
)

func (h *handler) AddTimeTracking(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	track := entity.TimeTracking{}
	if err := json.NewDecoder(r.Body).Decode(&track); err != nil {
		WriteError(w, err, http.StatusBadRequest)
		return
	}

	ctx := context.Background()

	if err := h.tableService.AddTimeTracking(ctx, track); err != nil {
		WriteError(w, err, http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
}

func (h *handler) TimeTracking(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	ctx := context.Background()

	tracks, err := h.tableService.GetTimeTracking(ctx)
	if err != nil {
		WriteError(w, err, http.StatusInternalServerError)
		return
	}

	WriteResponseJson(w, tracks)
}

func (h *handler) UpdateTimeTracking(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	track := entity.TimeTracking{}
	if err := json.NewDecoder(r.Body).Decode(&track); err != nil {
		WriteError(w, err, http.StatusBadRequest)
		return
	}

	ctx := context.Background()

	if err := h.tableService.UpdateTimeTracking(ctx, track); err != nil {
		WriteError(w, err, http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
}

func (h *handler) DeleteTimeTracking(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	ctx := context.Background()

	if err := h.tableService.DeleteTimeTracking(ctx, path.Base(r.URL.Path)); err != nil {
		WriteError(w, err, http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
}
