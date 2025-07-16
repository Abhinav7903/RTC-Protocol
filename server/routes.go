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
