package api

import (
	"net/http"

	"github.com/gorilla/mux"
)

type Server struct {
	*http.Server
}

func NewServer(port string, ah *AnnouncementHandler) *Server {
	r := mux.NewRouter()

	setAnnouncementRoutes(r, ah)

	return &Server{
		&http.Server{
			Addr:    ":" + port,
			Handler: r,
		},
	}
}

func setAnnouncementRoutes(r *mux.Router, ah *AnnouncementHandler) {
	r.HandleFunc("/announcements", ah.CreateAnnouncement).Methods(http.MethodPost, http.MethodOptions)
	r.HandleFunc("/announcements/{id}", ah.AnnouncementByID).Methods(http.MethodGet, http.MethodOptions)
}
