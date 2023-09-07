package v1

import (
	"context"
	"errors"
	"github.com/sav1nbrave4code/APG3/internal/entity/utils"
	"net/http"
	"strconv"
)

func (h *handler) GetFncTransferredPoints(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	ctx := context.Background()

	transfer, err := h.functionService.FncTransferredPoints(ctx)
	if err != nil {
		WriteError(w, err, http.StatusInternalServerError)
		return
	}

	WriteResponseJson(w, transfer)
}

func (h *handler) GetFncXpTask(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	ctx := context.Background()

	transfer, err := h.functionService.FncXpTask(ctx)
	if err != nil {
		WriteError(w, err, http.StatusInternalServerError)
		return
	}

	WriteResponseJson(w, transfer)
}

func (h *handler) GetFncPeersDontLeave(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	ctx := context.Background()

	if !r.URL.Query().Has("date") {
		WriteError(w, errors.New("query parameter date is missing"), http.StatusBadRequest)
		return
	}

	date, err := utils.ParseDate(r.URL.Query().Get("date"))
	if err != nil {
		WriteError(w, err, http.StatusBadRequest)
		return
	}

	dontLeave, err := h.functionService.FncPeersDontLeave(ctx, date)
	if err != nil {
		WriteError(w, err, http.StatusInternalServerError)
		return
	}

	WriteResponseJson(w, dontLeave)
}

func (h *handler) GetFncSuccessFailureChecks(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	ctx := context.Background()

	successFailure, err := h.functionService.FncSuccessFailureChecks(ctx)
	if err != nil {
		WriteError(w, err, http.StatusInternalServerError)
		return
	}

	WriteResponseJson(w, successFailure)
}

func (h *handler) GetFncPointsChangeV1(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	ctx := context.Background()

	successFailure, err := h.functionService.FncPointsChangeV1(ctx)
	if err != nil {
		WriteError(w, err, http.StatusInternalServerError)
		return
	}

	WriteResponseJson(w, successFailure)
}

func (h *handler) GetFncPointsChangeV2(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	ctx := context.Background()

	successFailure, err := h.functionService.FncPointsChangeV2(ctx)
	if err != nil {
		WriteError(w, err, http.StatusInternalServerError)
		return
	}

	WriteResponseJson(w, successFailure)
}

func (h *handler) GetFncOftenTaskPerDay(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	ctx := context.Background()

	oftenTasks, err := h.functionService.FncOftenTaskPerDay(ctx)
	if err != nil {
		WriteError(w, err, http.StatusInternalServerError)
		return
	}

	WriteResponseJson(w, oftenTasks)
}

func (h *handler) GetFncLastP2PDuration(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	ctx := context.Background()

	duration, err := h.functionService.FncLastP2PDuration(ctx)
	if err != nil {
		WriteError(w, err, http.StatusInternalServerError)
		return
	}

	WriteResponseJson(w, duration)
}

func (h *handler) GetFncListLastExPeer(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	ctx := context.Background()

	ex := r.URL.Query().Get("ex")
	if ex == "" {
		WriteError(w, errors.New("query parameter ex is missing"), http.StatusBadRequest)
		return
	}

	lastEx, err := h.functionService.FncListLastExPeer(ctx, ex)
	if err != nil {
		WriteError(w, err, http.StatusInternalServerError)
		return
	}

	WriteResponseJson(w, lastEx)
}

func (h *handler) GetFncPeersForP2P(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	ctx := context.Background()

	peers, err := h.functionService.FncPeersForP2P(ctx)
	if err != nil {
		WriteError(w, err, http.StatusInternalServerError)
		return
	}

	WriteResponseJson(w, peers)
}

func (h *handler) GetFncStatisticBlock(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	ctx := context.Background()

	block1 := r.URL.Query().Get("block_1")
	if block1 == "" {
		WriteError(w, errors.New("query parameter block_1 is missing"), http.StatusBadRequest)
		return
	}
	block2 := r.URL.Query().Get("block_2")
	if block2 == "" {
		WriteError(w, errors.New("query parameter block_2 is missing"), http.StatusBadRequest)
		return
	}

	statistic, err := h.functionService.FncStatisticBlock(ctx, block1, block2)
	if err != nil {
		WriteError(w, err, http.StatusInternalServerError)
		return
	}

	WriteResponseJson(w, statistic)
}

func (h *handler) GetFncMostFriendly(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	ctx := context.Background()

	if !r.URL.Query().Has("n") {
		WriteError(w, errors.New("query parameter n is missing"), http.StatusBadRequest)
		return
	}

	n, err := strconv.Atoi(r.URL.Query().Get("n"))
	if err != nil {
		WriteError(w, err, http.StatusBadRequest)
		return
	}

	friendly, err := h.functionService.FncMostFriendly(ctx, n)
	if err != nil {
		WriteError(w, err, http.StatusInternalServerError)
		return
	}

	WriteResponseJson(w, friendly)
}

func (h *handler) GetFncSuccessAtBirthDay(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	ctx := context.Background()

	success, err := h.functionService.FncSuccessAtBirthDay(ctx)
	if err != nil {
		WriteError(w, err, http.StatusInternalServerError)
		return
	}

	WriteResponseJson(w, success)
}

func (h *handler) GetFncPeerXpSum(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	ctx := context.Background()

	sum, err := h.functionService.FncPeerXpSum(ctx)
	if err != nil {
		WriteError(w, err, http.StatusInternalServerError)
		return
	}

	WriteResponseJson(w, sum)
}

func (h *handler) GetFncPassOneTwo(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	ctx := context.Background()

	task1 := r.URL.Query().Get("task_1")
	if task1 == "" {
		WriteError(w, errors.New("query parameter task_1 is missing"), http.StatusBadRequest)
		return
	}
	task2 := r.URL.Query().Get("task_2")
	if task2 == "" {
		WriteError(w, errors.New("query parameter task_2 is missing"), http.StatusBadRequest)
		return
	}
	task3 := r.URL.Query().Get("task_3")
	if task3 == "" {
		WriteError(w, errors.New("query parameter task_3 is missing"), http.StatusBadRequest)
		return
	}

	sum, err := h.functionService.FncPassOneTwo(ctx, task1, task2, task3)
	if err != nil {
		WriteError(w, err, http.StatusInternalServerError)
		return
	}

	WriteResponseJson(w, sum)
}

func (h *handler) GetFncPreviousTasks(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	ctx := context.Background()

	tasks, err := h.functionService.FncPreviousTasks(ctx)
	if err != nil {
		WriteError(w, err, http.StatusInternalServerError)
		return
	}

	WriteResponseJson(w, tasks)
}

func (h *handler) GetFncSuccessfulDays(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	ctx := context.Background()

	if !r.URL.Query().Has("n") {
		WriteError(w, errors.New("query parameter n is missing"), http.StatusBadRequest)
	}

	n, err := strconv.Atoi(r.URL.Query().Get("n"))
	if err != nil {
		WriteError(w, err, http.StatusBadRequest)
		return
	}

	days, err := h.functionService.FncSuccessfulDays(ctx, n)
	if err != nil {
		WriteError(w, err, http.StatusInternalServerError)
		return
	}

	WriteResponseJson(w, days)
}

func (h *handler) GetFncPeerMostTasks(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	ctx := context.Background()

	peer, err := h.functionService.FncPeerMostTasks(ctx)
	if err != nil {
		WriteError(w, err, http.StatusInternalServerError)
		return
	}

	WriteResponseJson(w, peer)
}

func (h *handler) GetFncPeerMostXp(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	ctx := context.Background()

	peer, err := h.functionService.FncPeerMostXp(ctx)
	if err != nil {
		WriteError(w, err, http.StatusInternalServerError)
		return
	}

	WriteResponseJson(w, peer)
}

func (h *handler) GetFncMaxTimeDate(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	ctx := context.Background()

	if !r.URL.Query().Has("date") {
		WriteError(w, errors.New("query parameter date is missing"), http.StatusBadRequest)
		return
	}

	date, err := utils.ParseDate(r.URL.Query().Get("date"))
	if err != nil {
		WriteError(w, err, http.StatusBadRequest)
		return
	}

	maxTime, err := h.functionService.FncMaxTimeDate(ctx, date)
	if err != nil {
		WriteError(w, err, http.StatusInternalServerError)
		return
	}

	WriteResponseJson(w, maxTime)
}

func (h *handler) GetFncTimePeerByTime(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	ctx := context.Background()

	time := r.URL.Query().Get("time")
	if time == "" {
		WriteError(w, errors.New("query parameter time is missing"), http.StatusBadRequest)
		return
	}

	if !r.URL.Query().Has("n") {
		WriteError(w, errors.New("query parameter n is missing"), http.StatusBadRequest)
		return
	}

	n, err := strconv.Atoi(r.URL.Query().Get("n"))
	if err != nil {
		WriteError(w, err, http.StatusBadRequest)
		return
	}

	peer, err := h.functionService.FncTimePeerByTime(ctx, time, n)
	if err != nil {
		WriteError(w, err, http.StatusInternalServerError)
		return
	}

	WriteResponseJson(w, peer)
}

func (h *handler) GetFncEnterPeerByDay(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	ctx := context.Background()

	if !r.URL.Query().Has("n") {
		WriteError(w, errors.New("query parameter n is missing"), http.StatusBadRequest)
		return
	}
	n, err := strconv.Atoi(r.URL.Query().Get("n"))
	if err != nil {
		WriteError(w, err, http.StatusBadRequest)
		return
	}

	if !r.URL.Query().Has("m") {
		WriteError(w, errors.New("query parameter m is missing"), http.StatusBadRequest)
		return
	}
	m, err := strconv.Atoi(r.URL.Query().Get("m"))
	if err != nil {
		WriteError(w, err, http.StatusBadRequest)
		return
	}

	peer, err := h.functionService.FncEnterPeerByDay(ctx, n, m)
	if err != nil {
		WriteError(w, err, http.StatusInternalServerError)
		return
	}

	WriteResponseJson(w, peer)
}

func (h *handler) GetFncLastFeastCame(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	ctx := context.Background()

	if !r.URL.Query().Has("date") {
		WriteError(w, errors.New("query parameter date is missing"), http.StatusBadRequest)
		return
	}

	date, err := utils.ParseDate(r.URL.Query().Get("date"))
	if err != nil {
		WriteError(w, err, http.StatusBadRequest)
		return
	}

	feast, err := h.functionService.FncLastFeastCame(ctx, date)
	if err != nil {
		WriteError(w, err, http.StatusInternalServerError)
		return
	}

	WriteResponseJson(w, feast)
}

func (h *handler) GetFncMoreThenTimePeer(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	ctx := context.Background()

	time := r.URL.Query().Get("time")
	if time == "" {
		WriteError(w, errors.New("query parameter time is missing"), http.StatusBadRequest)
		return
	}

	peer, err := h.functionService.FncMoreThenTimePeer(ctx, time)
	if err != nil {
		WriteError(w, err, http.StatusInternalServerError)
		return
	}

	WriteResponseJson(w, peer)
}

func (h *handler) GetFncEarlyEntries(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	ctx := context.Background()

	entries, err := h.functionService.FncEarlyEntries(ctx)
	if err != nil {
		WriteError(w, err, http.StatusInternalServerError)
		return
	}

	WriteResponseJson(w, entries)
}
