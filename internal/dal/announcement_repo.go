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
	q := `INSERT INTO announcements (title, description, price) VALUES ($1, $2, $3) RETURNING id`

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

func (r *Repository) ByID(ctx context.Context, id int64) (entity.Announcement, error) {
	q := `
SELECT
    a.id,
    a.title,
    a.description,
    a.price,
    p.link
FROM announcements AS a 
JOIN photos AS p ON a.id = p.announcement_id
WHERE a.id = $1`

	var a entity.Announcement

	rows, err := r.db.QueryContext(ctx, q, id)
	if err != nil {
		return a, err
	}
	defer rows.Close()

	for rows.Next() {
		var link string

		err = rows.Scan(&a.ID, &a.Title, &a.Description, &a.Price, &link)
		if err != nil {
			return a, err
		}

		a.Photos = append(a.Photos, link)
	}

	err = rows.Err()
	if err != nil {
		return a, err
	}

	return a, nil
}
