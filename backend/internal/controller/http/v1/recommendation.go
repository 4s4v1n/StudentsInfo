package v1

import (
	"context"
	"encoding/json"
	"github.com/sav1nbrave4code/APG3/internal/entity"
	"net/http"
	"path"
)

func (h *handler) AddRecommendation(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	recommendation := entity.Recommendation{}
	if err := json.NewDecoder(r.Body).Decode(&recommendation); err != nil {
		WriteError(w, err, http.StatusBadRequest)
		return
	}

	ctx := context.Background()

	if err := h.tableService.AddRecommendation(ctx, recommendation); err != nil {
		WriteError(w, err, http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
}

func (h *handler) Recommendations(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	ctx := context.Background()

	recommendations, err := h.tableService.GetRecommendations(ctx)
	if err != nil {
		WriteError(w, err, http.StatusInternalServerError)
		return
	}

	WriteResponseJson(w, recommendations)
}

func (h *handler) UpdateRecommendation(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	recommendation := entity.Recommendation{}
	if err := json.NewDecoder(r.Body).Decode(&recommendation); err != nil {
		WriteError(w, err, http.StatusBadRequest)
		return
	}

	ctx := context.Background()

	if err := h.tableService.UpdateRecommendation(ctx, recommendation); err != nil {
		WriteError(w, err, http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
}

func (h *handler) DeleteRecommendation(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	ctx := context.Background()

	if err := h.tableService.DeleteRecommendation(ctx, path.Base(r.URL.Path)); err != nil {
		WriteError(w, err, http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
}
