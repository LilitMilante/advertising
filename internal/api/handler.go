package api

import (
	"context"
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/LilitMilante/advertising/internal/services/entity"
	"github.com/gorilla/mux"
)

type Service interface {
	Create(ctx context.Context, ca entity.CreateAnnouncement) (int64, error)
	ByID(ctx context.Context, id int64) (entity.Announcement, error)
}

type AnnouncementHandler struct {
	s Service
}

func NewHandler(s Service) *AnnouncementHandler {
	return &AnnouncementHandler{s: s}
}

func (h *AnnouncementHandler) CreateAnnouncement(w http.ResponseWriter, r *http.Request) {
	var req entity.CreateAnnouncement

	err := json.NewDecoder(r.Body).Decode(&req)
	if err != nil {
		SendErrJSON(w, err)
		return
	}

	id, err := h.s.Create(r.Context(), req)
	if err != nil {
		SendErrJSON(w, err)
		return
	}

	resp := entity.Announcement{ID: id}

	SendJSON(w, resp)
}

func (h *AnnouncementHandler) AnnouncementByID(w http.ResponseWriter, r *http.Request) {
	idStr := mux.Vars(r)["id"]

	id, err := strconv.ParseInt(idStr, 10, 64)
	if err != nil {
		SendErrJSON(w, err)
		return
	}

	a, err := h.s.ByID(r.Context(), id)
	if err != nil {
		SendErrJSON(w, err)
		return
	}

	SendJSON(w, a)
}
