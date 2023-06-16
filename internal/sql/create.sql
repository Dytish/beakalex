-- DROP TABLE IF EXISTS user CASCADE;

CREATE TABLE users (
    id SERIAL PRIMARY KEY,
    email varchar(255) NOT NULL,
    login varchar(255) NOT NULL,
    password varchar(255) NOT NULL,
    created_at TIMESTAMP WITH TIME ZONE,
    updated_at TIMESTAMP WITH TIME ZONE,
    deleted_at TIMESTAMP WITH TIME ZONE
);

CREATE TABLE bots(
    id SERIAL PRIMARY KEY,
    id_user INTEGER NOT NULL,
    token varchar(255) NOT NULL,
    body json,
    created_at TIMESTAMP WITH TIME ZONE,
    updated_at TIMESTAMP WITH TIME ZONE,
    deleted_at TIMESTAMP WITH TIME ZONE,
    FOREIGN KEY (id_user) REFERENCES users (id)
);
--  drop table bots;