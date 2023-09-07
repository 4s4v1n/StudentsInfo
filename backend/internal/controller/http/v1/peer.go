package v1

import (
	"context"
	"encoding/json"
	"net/http"
	"path"

	"github.com/sav1nbrave4code/APG3/internal/entity"
)

func (h *handler) AddPeer(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	peer := entity.Peer{}
	if err := json.NewDecoder(r.Body).Decode(&peer); err != nil {
		WriteError(w, err, http.StatusBadRequest)
		return
	}

	ctx := context.Background()

	if err := h.tableService.AddPeer(ctx, peer); err != nil {
		WriteError(w, err, http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
}

func (h *handler) Peers(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	ctx := context.Background()

	peers, err := h.tableService.GetPeers(ctx)
	if err != nil {
		WriteError(w, err, http.StatusInternalServerError)
		return
	}

	WriteResponseJson(w, peers)
}

func (h *handler) UpdatePeer(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	peer := entity.Peer{}
	if err := json.NewDecoder(r.Body).Decode(&peer); err != nil {
		WriteError(w, err, http.StatusBadRequest)
		return
	}

	ctx := context.Background()

	if err := h.tableService.UpdatePeer(ctx, peer); err != nil {
		WriteError(w, err, http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
}

func (h *handler) DeletePeer(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	ctx := context.Background()

	if err := h.tableService.DeletePeer(ctx, path.Base(r.URL.Path)); err != nil {
		WriteError(w, err, http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
}
