CREATE DATABASE IF NOT EXISTS echo;
USE echo;

CREATE TABLE IF NOT EXISTS users
(
    `id` INT NOT NULL AUTO_INCREMENT,
    `name` VARCHAR(255) NOT NULL,
    `email` VARCHAR(255) NOT NULL,
    PRIMARY KEY (`id`)
);

INSERT INTO users
    (name, email)
VALUES
    ('John', 'aaa@mail.com'),
    ('Paul', 'bbb@mail.com'),
    ('George', 'ccc@mail.com'),
    ('Ringo', 'ddd@mail.com');