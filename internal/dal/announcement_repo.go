package dal

import (
	"context"
	"database/sql"

	"github.com/LilitMilante/advertising/internal/services/entity"
)

type Repository struct {
	db *sql.DB
}

func NewRepository(db *sql.DB) *Repository {
	return &Repository{
		db: db,
	}
}

// Add new announcement.
func (r *Repository) Add(ctx context.Context, ac entity.CreateAnnouncement) (int64, error) {
	tx, err := r.db.Begin()
	if err != nil {
		return 0, err
	}
	defer tx.Rollback()

	var id int64
	q := `INSERT INTO announcements (title, description, price) VALUES ($1, $2, $3)`

	err = tx.QueryRowContext(ctx, q, ac.Title, ac.Description, ac.Price).Scan(&id)
	if err != nil {
		return 0, err
	}

	q = `INSERT INTO photos (announcement_id, link) VALUES ($1, $2)`

	stmt, err := tx.PrepareContext(ctx, q)
	if err != nil {
		return 0, err
	}
	defer stmt.Close()

	for _, v := range ac.Photos {
		_, err = stmt.ExecContext(ctx, id, v)
		if err != nil {
			return 0, err
		}
	}

	err = tx.Commit()
	if err != nil {
		return 0, err
	}

	return id, nil
}
