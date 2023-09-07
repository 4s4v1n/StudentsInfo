package v1

import (
	"github.com/go-chi/chi"
	"github.com/go-chi/cors"
	"github.com/ironstar-io/chizerolog"
	"github.com/rs/zerolog/log"
	"github.com/sav1nbrave4code/APG3/internal/service/data_manager"
	"github.com/sav1nbrave4code/APG3/internal/service/function_manager"
	"github.com/sav1nbrave4code/APG3/internal/service/table_manager"
)

type handler struct {
	mux             *chi.Mux
	tableService    table_manager.Service
	dataService     data_manager.Service
	functionService function_manager.Service
}

func New(mux *chi.Mux, tableService table_manager.Service, dataService data_manager.Service,
	functionService function_manager.Service) *handler {
	return &handler{
		mux:             mux,
		tableService:    tableService,
		dataService:     dataService,
		functionService: functionService,
	}
}

func (h *handler) Run() {
	h.mux.Use(chizerolog.LoggerMiddleware(&log.Logger))
	h.mux.Use(cors.Handler(cors.Options{
		AllowedOrigins: []string{"http://localhost:3000", "http://localhost:8080", "http://localhost:4500"},
		AllowedMethods: []string{"GET", "POST", "PUT", "DELETE", "OPTIONS", "PATCH"},
		AllowedHeaders: []string{"Accept", "Authorization", "Content-Type", "X-CSRF-Token"},
	}))

	h.mux.Route("/api/v1", func(r chi.Router) {
		r.Route("/peer", func(r chi.Router) {
			r.Post("/", h.AddPeer)
			r.Get("/", h.Peers)
			r.Patch("/", h.UpdatePeer)
			r.Delete("/{nickname}", h.DeletePeer)
		})

		r.Route("/task", func(r chi.Router) {
			r.Post("/", h.AddTask)
			r.Get("/", h.Tasks)
			r.Patch("/", h.UpdateTask)
			r.Delete("/{title}", h.DeleteTask)
		})

		r.Route("/check", func(r chi.Router) {
			r.Post("/", h.AddCheck)
			r.Get("/", h.Checks)
			r.Patch("/", h.UpdateCheck)
			r.Delete("/{id}", h.DeleteCheck)
		})

		r.Route("/p2p", func(r chi.Router) {
			r.Post("/", h.AddP2P)
			r.Get("/", h.P2Ps)
			r.Patch("/", h.UpdateP2P)
			r.Delete("/{id}", h.DeleteP2P)
		})

		r.Route("/verter", func(r chi.Router) {
			r.Post("/", h.AddVerter)
			r.Get("/", h.Verters)
			r.Patch("/", h.UpdateVerter)
			r.Delete("/{id}", h.DeleteVerter)
		})

		r.Route("/transferred_points", func(r chi.Router) {
			r.Post("/", h.AddTransferredPoints)
			r.Get("/", h.TransferredPoints)
			r.Patch("/", h.UpdateTransferredPoints)
			r.Delete("/{id}", h.DeleteTransferredPoints)
		})

		r.Route("/friends", func(r chi.Router) {
			r.Post("/", h.AddFriends)
			r.Get("/", h.Friends)
			r.Patch("/", h.UpdateFriends)
			r.Delete("/{id}", h.DeleteFriends)
		})

		r.Route("/recommendation", func(r chi.Router) {
			r.Post("/", h.AddRecommendation)
			r.Get("/", h.Recommendations)
			r.Patch("/", h.UpdateRecommendation)
			r.Delete("/{id}", h.DeleteRecommendation)
		})

		r.Route("/xp", func(r chi.Router) {
			r.Post("/", h.AddXP)
			r.Get("/", h.XPs)
			r.Patch("/", h.UpdateXP)
			r.Delete("/{id}", h.DeleteXP)
		})

		r.Route("/time_tracking", func(r chi.Router) {
			r.Post("/", h.AddTimeTracking)
			r.Get("/", h.TimeTracking)
			r.Patch("/", h.UpdateTimeTracking)
			r.Delete("/{id}", h.DeleteTimeTracking)
		})

		r.Post("/import/{table}", h.Import)
		r.Get("/export/{table}", h.Export)
		r.Post("/raw_query", h.RawQuery)

		r.Route("/fnc", func(r chi.Router) {
			r.Get("/transferred_points", h.GetFncTransferredPoints)
			r.Get("/xp_task", h.GetFncXpTask)
			r.Get("/peers_dont_leave", h.GetFncPeersDontLeave)
			r.Get("/success_failure_checks", h.GetFncSuccessFailureChecks)
			r.Get("/points_change_v1", h.GetFncPointsChangeV1)
			r.Get("/points_change_v2", h.GetFncPointsChangeV2)
			r.Get("/often_task_per_day", h.GetFncOftenTaskPerDay)
			r.Get("/last_p2p_duration", h.GetFncLastP2PDuration)
			r.Get("/list_last_ex_peer", h.GetFncListLastExPeer)
			r.Get("/peers_for_p2p", h.GetFncPeersForP2P)
			r.Get("/statistic_block", h.GetFncStatisticBlock)
			r.Get("/most_friendly", h.GetFncMostFriendly)
			r.Get("/success_at_birthday", h.GetFncSuccessAtBirthDay)
			r.Get("/peer_xp_sum", h.GetFncPeerXpSum)
			r.Get("/pass_one_two", h.GetFncPassOneTwo)
			r.Get("/previous_tasks", h.GetFncPreviousTasks)
			r.Get("/successful_days", h.GetFncSuccessfulDays)
			r.Get("/peer_most_tasks", h.GetFncPeerMostTasks)
			r.Get("/peer_most_xp", h.GetFncPeerMostXp)
			r.Get("/max_time_date", h.GetFncMaxTimeDate)
			r.Get("/time_peer_by_time", h.GetFncTimePeerByTime)
			r.Get("/enter_peer_by_day", h.GetFncEnterPeerByDay)
			r.Get("/last_feast_came", h.GetFncLastFeastCame)
			r.Get("/more_then_time_peer", h.GetFncMoreThenTimePeer)
			r.Get("/early_entries", h.GetFncEarlyEntries)
		})
	})
}
