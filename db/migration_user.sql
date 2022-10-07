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
--

--
insert into orders(user_id, item)
values 
(1, 'item1'),
(1, 'item2'),
(1, 'item3');
--

-- JOIN
select *
from orders o 
	join users u on o.user_id = u.id;
--