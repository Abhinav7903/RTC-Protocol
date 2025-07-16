package server

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/sirupsen/logrus"
)

func (s *Server) handleCreateSignal() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		logrus.Debug("Running in handleCreateSignal")

		var req struct {
			RoomID     int             `json:"room_id"`   // changed to int
			SenderID   int             `json:"sender_id"` // changed to int
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
		logrus.Debug("Running in handleGetSignals")

		roomIDStr := r.URL.Query().Get("room_id")
		senderIDStr := r.URL.Query().Get("sender_id")

		if roomIDStr == "" || senderIDStr == "" {
			s.respond(w, &ResponseMsg{Message: "Missing room_id or sender_id"}, http.StatusBadRequest, nil)
			return
		}
		// Assuming roomID and senderID are integers, convert them
		roomID, err := strconv.Atoi(roomIDStr)
		if err != nil {
			s.respond(w, &ResponseMsg{Message: "Invalid room_id"}, http.StatusBadRequest, err)
			return
		}
		senderID, err := strconv.Atoi(senderIDStr)
		if err != nil {
			s.respond(w, &ResponseMsg{Message: "Invalid sender_id"}, http.StatusBadRequest, err)
			return
		}

		signals, err := s.signalRepo.GetSignalsByRoomExcludingSender(roomID, senderID)
		if err != nil {
			logrus.Error("Failed to retrieve signals:", err)
			s.respond(w, &ResponseMsg{Message: "Failed to retrieve signals"}, http.StatusInternalServerError, err)
			return
		}

		s.respond(w, &ResponseMsg{
			Message: "Signals fetched successfully",
			Data:    signals,
		}, http.StatusOK, nil)
	}
}
