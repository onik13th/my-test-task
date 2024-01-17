CREATE TABLE users (
    id bigserial not null primary key,
    name varchar not null unique,
    surname varchar not null unique,
    patronymic varchar not null
);
