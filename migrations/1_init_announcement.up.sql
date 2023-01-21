CREATE TABLE announcements (
    id BIGSERIAL PRIMARY KEY,
    title TEXT NOT NULL,
    description TEXT NOT NULL,
    price BIGINT NOT NULL
);

CREATE TABLE photos (
    id_announcement BIGINT REFERENCES announcements(id) NOT NULL,
    link TEXT NOT NULL
);
