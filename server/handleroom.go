package server

import (
	"encoding/json"
	"net/http"

	"github.com/sirupsen/logrus"
)

func (s *Server) handleCreateRoom() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req struct {
			Name     string `json:"name"`
			RoomType string `json:"room_type"` // "single" or "group"
		}

		if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
			s.respond(w, &ResponseMsg{Message: "Invalid request"}, http.StatusBadRequest, err)
			return
		}

		room, err := s.roomRepo.CreateRoom(req.Name, req.RoomType)
		if err != nil {
			logrus.Error("Failed to create room:", err)
			s.respond(w, &ResponseMsg{Message: "Failed to create room"}, http.StatusInternalServerError, err)
			return
		}

		s.respond(w, &ResponseMsg{Message: "Room created", Data: room}, http.StatusCreated, nil)
	}
}

func (s *Server) handleListRooms() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		rooms, err := s.roomRepo.ListRooms()
		if err != nil {
			s.respond(w, &ResponseMsg{Message: "Failed to fetch rooms"}, http.StatusInternalServerError, err)
			return
		}

		s.respond(w, &ResponseMsg{Message: "Rooms retrieved", Data: rooms}, http.StatusOK, nil)
	}
}

func (s *Server) handleGetRoomByID() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		id := r.URL.Query().Get("id")
		if id == "" {
			s.respond(w, &ResponseMsg{Message: "Missing room ID"}, http.StatusBadRequest, nil)
			return
		}
		room, err := s.roomRepo.GetRoomByID(id)
		if err != nil {
			s.respond(w, &ResponseMsg{Message: "Room not found"}, http.StatusNotFound, err)
			return
		}

		s.respond(w, &ResponseMsg{Message: "Room found", Data: room}, http.StatusOK, nil)
	}
}

func (s *Server) handleDeleteRoom() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		id := r.URL.Query().Get("id")
		if id == "" {
			s.respond(w, &ResponseMsg{Message: "Missing room ID"}, http.StatusBadRequest, nil)
			return
		}
		err := s.roomRepo.DeleteRoom(id)
		if err != nil {
			s.respond(w, &ResponseMsg{Message: "Failed to delete room"}, http.StatusInternalServerError, err)
			return
		}

		s.respond(w, &ResponseMsg{Message: "Room deleted"}, http.StatusOK, nil)
	}
}
