-- creating user table in postgres DATABASE
CREATE TABLE users(
    id SERIAL NOT NULL PRIMARY KEY,
    last_name TEXT,
    first_name TEXT,
    email TEXT
);

-- creating order table
create table orders(
    id SERIAL NOT NULL PRIMARY KEY,
    user_id int,
    item text,
    FOREIGN KEY (user_id) REFERENCES users(id)
);