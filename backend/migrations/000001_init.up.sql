CREATE SCHEMA watchtogether;

CREATE TABLE watchtogether.videos (
    id         UUID PRIMARY KEY DEFAULT uuidv7(),
    title      TEXT NOT NULL,
    source     TEXT NOT NULL,
    url        TEXT NOT NULL,
    duration   BIGINT NOT NULL DEFAULT 0
);

CREATE TABLE watchtogether.users (
    id              UUID PRIMARY KEY DEFAULT uuidv7(),
    nickname        TEXT NOT NULL,
    password_hash   TEXT NOT NULL,
    email           TEXT UNIQUE NOT NULL,
    avatar          TEXT,
    created_at      TIMESTAMPTZ NOT NULL DEFAULT now()
);

CREATE TABLE watchtogether.rooms (
    id              TEXT PRIMARY KEY,
    title           TEXT NOT NULL,
    host_id         UUID NOT NULL REFERENCES watchtogether.users(id),
    visibility      TEXT NOT NULL DEFAULT 'public',
    password_hash   TEXT NOT NULL DEFAULT '',
    video_id        UUID REFERENCES watchtogether.videos(id),
    created_at      TIMESTAMPTZ NOT NULL DEFAULT now()
);
