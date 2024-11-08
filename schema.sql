create extension if not exists pgcrypto;

-- таблица с пользователями
CREATE TABLE IF NOT EXISTS users (
    id BIGSERIAL PRIMARY KEY,
    login TEXT NOT NULL UNIQUE,
    password_hash TEXT NOT NULL,
    created_at TIMESTAMPTZ NOT NULL DEFAULT NOW()
);


-- таблица сессий
CREATE TABLE IF NOT EXISTS sessions (
    id TEXT PRIMARY KEY DEFAULT encode(gen_random_bytes(16), 'hex'),
    uid BIGINT NOT NULL,
    -- user id
    created_at TIMESTAMPTZ NOT NULL DEFAULT NOW()
);


-- таблица с файлами
CREATE TABLE IF NOT EXISTS assets (
    name TEXT NOT NULL,
    uid BIGINT NOT NULL,
    data BYTEA NOT NULL,
    created_at TIMESTAMPTZ NOT NULL DEFAULT NOW(),
    PRIMARY KEY (name, uid)
);


-- user id
CREATE TABLE IF NOT EXISTS assets (
    name TEXT NOT NULL,
    uid BIGINT NOT NULL,
    data BYTEA NOT NULL,
    created_at TIMESTAMPTZ NOT NULL DEFAULT NOW(),
    PRIMARY KEY (name, uid)
);


-- тестовый пользователь
insert into users
(login, password_hash)
values
('alice', encode(digest('secret', 'md5'),'hex'))
on conflict do nothing;