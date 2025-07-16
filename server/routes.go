package server

import (
	"net/http"
	"rtc/db/postgres"
	"rtc/pkg/participant"
	"rtc/pkg/room"
	"rtc/pkg/signal"

	"github.com/gorilla/mux"
)

type ResponseMsg struct {
	Message string      `json:"message"`
	Data    interface{} `json:"data,omitempty"`
}

type Server struct {
	router          *mux.Router
	postgres        *postgres.Postgres
	roomRepo        room.RoomRepository
	signalRepo      signal.SignalRepository
	participantRepo participant.ParticipantRepository
}

func (s *Server) RegisterRoutes() {
	s.router.HandleFunc("/ping", s.HandlePong())
	// Signal routes
	s.router.HandleFunc("/signal", s.handleCreateSignal()).Methods(http.MethodPost, http.MethodOptions)
	s.router.HandleFunc("/signal", s.handleGetSignals()).Methods(http.MethodGet, http.MethodOptions)
	// Room routes
	s.router.HandleFunc("/room", s.handleCreateRoom()).Methods(http.MethodPost, http.MethodOptions)
	s.router.HandleFunc("/room", s.handleListRooms()).Methods(http.MethodGet, http.MethodOptions)
	s.router.HandleFunc("/room/{id}", s.handleGetRoomByID()).Methods(http.MethodGet, http.MethodOptions)
	s.router.HandleFunc("/room/{id}", s.handleDeleteRoom()).Methods(http.MethodDelete, http.MethodOptions)

	// Participant routes
	s.router.HandleFunc("/participant", s.handleCreateParticipant()).Methods(http.MethodPost, http.MethodOptions)
	s.router.HandleFunc("/participant", s.handleGetParticipantsByRoom()).Methods(http.MethodGet, http.MethodOptions)
	s.router.HandleFunc("/participant/{id}", s.handleDeleteParticipant()).Methods(http.MethodDelete, http.MethodOptions)

}

func (s *Server) HandlePong() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		s.respond(
			w,
			"pong",
			http.StatusOK,
			nil,
		)
	}
}
