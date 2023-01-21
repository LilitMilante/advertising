CREATE TABLE announcements (
    id BIGSERIAL PRIMARY KEY,
    title TEXT NOT NULL,
    description TEXT NOT NULL,
    price BIGINT NOT NULL
);

CREATE TABLE photos (
    announcement_id BIGINT REFERENCES announcements(id) NOT NULL,
    link TEXT NOT NULL
);
