CREATE TABLE users
(
    id            BIGSERIAL PRIMARY KEY,
    role_id       integer references roles (id),
    name          varchar(255) not null,
    username      varchar(255) not null unique,
    password_hash varchar(255) not null
);



CREATE TABLE carts
(
    id         BIGSERIAL PRIMARY KEY,
    user_id    integer references users (id),
    created_at timestamp not null DEFAULT current_timestamp,
    removed    BOOLEAN            DEFAULT FALSE,
    delivered  BOOLEAN            DEFAULT FALSE

);

CREATE TABLE roles
(
    id   bigserial primary key,
    name text not null
);




CREATE TABLE items
(
    id            BIGSERIAL PRIMARY KEY, -- думаю тоже тут сделать ключ уникальным
    title         TEXT    NOT NULL,
    categories_id INTEGER REFERENCES categories (id),
    price         INTEGER NOT NULL,
    count         INTEGER

);


CREATE TABLE carts_users
(
    cart_id integer references carts (id),
    user_id integer references users (id)

);


CREATE TABLE carts_items
(
    cart_id INTEGER REFERENCES carts (id),
    item_id INTEGER REFERENCES items (id),
    count   INTEGER NOT NULL
);


CREATE TABLE categories
(
    id   BIGSERIAL PRIMARY KEY,
    name TEXT NOT NULL UNIQUE
);


CREATE TABLE orders
(
    id        bigserial PRIMARY KEY,
    phone     TEXT NOT NULL,
    address   TEXT NOT NULL,
    cart_id   INTEGER references carts (id),
    delivered BOOLEAN DEFAULT FALSE

);

CREATE TABLE orders_items
(
    items_id INTEGER references items (id),
    carts_id INTEGER references carts (id),
    order_id INTEGER references orders (id)
);





