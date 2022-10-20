CREATE DATABASE test
    WITH 
    OWNER = postgres
    ENCODING = 'UTF8';

CREATE TABLE IF NOT EXISTS employees(
    id SERIAL NOT NULL PRIMARY KEY,
    first_name CHARACTER VARYING(45) NOT NULL,
    last_name CHARACTER VARYING(45) NOT NULL,
    email CHARACTER VARYING(75) NOT NULL,
    status BOOLEAN NOT NULL DEFAULT true,
    created_at TIMESTAMP NOT NULL DEFAULT NOW(),
    updated_at TIMESTAMP
);