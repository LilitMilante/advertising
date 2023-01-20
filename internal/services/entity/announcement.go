package entity

type Announcement struct {
	ID          int64    `json:"id,omitempty"`
	Title       string   `json:"title,omitempty"`
	Description string   `json:"description,omitempty"`
	Photos      []string `json:"photos,omitempty"`
	Price       int64    `json:"price,omitempty"`
}

type CreateAnnouncement struct {
	Title       string   `json:"title"`
	Description string   `json:"description"`
	Photos      []string `json:"photos"`
	Price       int64    `json:"price"`
}
