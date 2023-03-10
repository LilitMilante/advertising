package services

import (
	"context"

	"github.com/LilitMilante/advertising/internal/services/entity"
)

type Repository interface {
	Add(ctx context.Context, ca entity.CreateAnnouncement) (int64, error)
	ByID(ctx context.Context, id int64) (entity.Announcement, error)
}

type Announcement struct {
	repo Repository
}

func NewAnnouncement(r Repository) *Announcement {
	return &Announcement{repo: r}
}

func (a *Announcement) Create(ctx context.Context, ca entity.CreateAnnouncement) (int64, error) {
	return a.repo.Add(ctx, ca)
}

func (a *Announcement) ByID(ctx context.Context, id int64) (entity.Announcement, error) {
	return a.repo.ByID(ctx, id)
}
