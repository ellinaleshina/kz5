CREATE TABLE if not exists users
(
    id       serial PRIMARY KEY,
    username varchar
);
CREATE TABLE if not exists posts
(
    id     serial PRIMARY KEY,
    author int,
    constraint author_fk foreign key (author) references users (id),
    posted timestamp,
    post_text text
)