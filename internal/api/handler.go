package api

import (
	"context"
	"encoding/json"
	"net/http"

	"github.com/LilitMilante/advertising/internal/services/entity"
)

type Service interface {
	Create(ctx context.Context, ca entity.CreateAnnouncement) (int64, error)
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
