CREATE TABLE users (
    id int NOT NULL AUTO_INCREMENT PRIMARY KEY,
    name VARCHAR (255),
    address VARCHAR (255),
    email VARCHAR (255),
    born_date TIMESTAMP
);


-- CREATE DATABASE
-- migrate -database "mysql://root:1234@tcp(localhost:3306)/go_gin_gonic" -path database/migrations up

-- DROP DATABASE
-- migrate -database "mysql://root:1234@tcp(localhost:3306)/go_gin_gonic" -path database/migrations down

-- CREATE TABLE
-- migrate create -ext sql -dir db/migrations create_users_table