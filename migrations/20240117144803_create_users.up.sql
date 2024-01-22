CREATE TABLE IF NOT EXISTS users (
    id BIGSERIAL PRIMARY KEY,
    name VARCHAR(255) NOT NULL,
    surname VARCHAR(255) NOT NULL,
    patronymic VARCHAR(255),
    age INTEGER,
    gender VARCHAR(10),
    nationalities VARCHAR(255)
);
--     age bigint not null,
--     gender varchar not null,
--     nationalities varchar not null
