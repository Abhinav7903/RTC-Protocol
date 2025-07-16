package server

import (
	"encoding/json"
	"net/http"
)

func (s *Server) handleCreateSignal() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req struct {
			RoomID     string          `json:"room_id"`
			SenderID   string          `json:"sender_id"`
			SignalType string          `json:"signal_type"`
			Payload    json.RawMessage `json:"payload"`
		}

		if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
			s.respond(w, &ResponseMsg{Message: "Invalid request body"}, http.StatusBadRequest, err)
			return
		}

		signal, err := s.signalRepo.CreateSignal(req.RoomID, req.SenderID, req.SignalType, req.Payload)
		if err != nil {
			s.respond(w, &ResponseMsg{Message: "Failed to store signal"}, http.StatusInternalServerError, err)
			return
		}

		s.respond(w, &ResponseMsg{
			Message: "Signal created successfully",
			Data:    signal,
		}, http.StatusCreated, nil)
	}
}

func (s *Server) handleGetSignals() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		roomID := r.URL.Query().Get("room_id")
		senderID := r.URL.Query().Get("sender_id")

		if roomID == "" || senderID == "" {
			s.respond(w, &ResponseMsg{Message: "Missing room_id or sender_id"}, http.StatusBadRequest, nil)
			return
		}

		signals, err := s.signalRepo.GetSignalsByRoomExcludingSender(roomID, senderID)
		if err != nil {
			s.respond(w, &ResponseMsg{Message: "Failed to retrieve signals"}, http.StatusInternalServerError, err)
			return
		}

		s.respond(w, &ResponseMsg{
			Message: "Signals fetched successfully",
			Data:    signals,
		}, http.StatusOK, nil)
	}
}
 