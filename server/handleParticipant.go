package server

import (
	"encoding/json"
	"net/http"
	"strconv"
)

func (s *Server) handleCreateParticipant() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req struct {
			RoomID      int    `json:"room_id"`
			DisplayName string `json:"display_name"`
		}

		if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
			s.respond(w, &ResponseMsg{Message: "Invalid request"}, http.StatusBadRequest, err)
			return
		}

		participant, err := s.participantRepo.CreateParticipant(req.RoomID, req.DisplayName)
		if err != nil {
			s.respond(w, &ResponseMsg{Message: "Failed to create participant"}, http.StatusInternalServerError, err)
			return
		}

		s.respond(w, &ResponseMsg{Message: "Participant joined", Data: participant}, http.StatusCreated, nil)
	}
}

func (s *Server) handleGetParticipantsByRoom() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		roomID := r.URL.Query().Get("room_id")
		if roomID == "" {
			s.respond(w, &ResponseMsg{Message: "Missing room_id"}, http.StatusBadRequest, nil)
			return
		}
		// Assuming roomID is an integer, convert it
		roomIDInt, err := strconv.Atoi(roomID)
		if err != nil {
			s.respond(w, &ResponseMsg{Message: "Invalid room_id"}, http.StatusBadRequest, err)
			return
		}

		participants, err := s.participantRepo.GetParticipantsByRoom(roomIDInt)
		if err != nil {
			s.respond(w, &ResponseMsg{Message: "Failed to get participants"}, http.StatusInternalServerError, err)
			return
		}

		s.respond(w, &ResponseMsg{Message: "Participants retrieved", Data: participants}, http.StatusOK, nil)
	}
}

func (s *Server) handleDeleteParticipant() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		id := r.URL.Query().Get("id")
		if id == "" {
			s.respond(w, &ResponseMsg{Message: "Missing participant ID"}, http.StatusBadRequest, nil)
			return
		}
		// Assuming id is an integer, convert it
		participantID, err := strconv.Atoi(id)
		if err != nil {
			s.respond(w, &ResponseMsg{Message: "Invalid participant ID"}, http.StatusBadRequest, err)
			return
		}
		err = s.participantRepo.DeleteParticipant(participantID)
		if err != nil {
			s.respond(w, &ResponseMsg{Message: "Failed to delete participant"}, http.StatusInternalServerError, err)
			return
		}

		s.respond(w, &ResponseMsg{Message: "Participant removed"}, http.StatusOK, nil)
	}
}
