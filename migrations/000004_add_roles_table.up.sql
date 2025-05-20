CREATE TABLE IF NOT EXISTS roles (
    id UUID NOT NULL PRIMARY KEY DEFAULT uuid_generate_v4(),
    role VARCHAR(50) NOT NULL
);

INSERT INTO roles (role)
VALUES ('user'),
       ('admin');