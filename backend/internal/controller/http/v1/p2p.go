package v1

import (
	"context"
	"encoding/json"
	"github.com/sav1nbrave4code/APG3/internal/entity"
	"github.com/sav1nbrave4code/APG3/internal/entity/service_types"
	"net/http"
	"path"
)

func (h *handler) AddP2P(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	p2p := service_types.P2PInsert{}
	if err := json.NewDecoder(r.Body).Decode(&p2p); err != nil {
		WriteError(w, err, http.StatusBadRequest)
		return
	}

	ctx := context.Background()

	if err := h.tableService.AddP2P(ctx, p2p); err != nil {
		WriteError(w, err, http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
}

func (h *handler) P2Ps(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	ctx := context.Background()

	p2ps, err := h.tableService.GetP2Ps(ctx)
	if err != nil {
		WriteError(w, err, http.StatusInternalServerError)
		return
	}

	WriteResponseJson(w, p2ps)
}

func (h *handler) UpdateP2P(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	p2p := entity.P2P{}
	if err := json.NewDecoder(r.Body).Decode(&p2p); err != nil {
		WriteError(w, err, http.StatusBadRequest)
		return
	}

	ctx := context.Background()

	if err := h.tableService.UpdateP2P(ctx, p2p); err != nil {
		WriteError(w, err, http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
}

func (h *handler) DeleteP2P(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	ctx := context.Background()

	if err := h.tableService.DeleteP2P(ctx, path.Base(r.URL.Path)); err != nil {
		WriteError(w, err, http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
}
