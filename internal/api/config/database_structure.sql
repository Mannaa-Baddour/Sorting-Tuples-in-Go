CREATE DATABASE users_files_db;

\c users_files_db

CREATE TABLE users(
      id BIGSERIAL PRIMARY KEY NOT NULL,
      username VARCHAR(30) NOT NULL UNIQUE,
      email VARCHAR(60) NOT NULL CONSTRAINT email_constraint
          UNIQUE,
      password VARCHAR(30) NOT NULL
);

CREATE TABLE files(
      id BIGSERIAL PRIMARY KEY NOT NULL,
      name VARCHAR(120) NOT NULL,
      type VARCHAR(60) NOT NULL,
      user_id BIGINT REFERENCES users(id) NOT NULL
);