package entity

type Announcement struct {
	ID          int64    `json:"id"`
	Title       string   `json:"title"`
	Description string   `json:"description"`
	Photos      []string `json:"photos"`
	Price       int64    `json:"price"`
}

type CreateAnnouncement struct {
	Title       string   `json:"title"`
	Description string   `json:"description"`
	Photos      []string `json:"photos"`
	Price       int64    `json:"price"`
}
