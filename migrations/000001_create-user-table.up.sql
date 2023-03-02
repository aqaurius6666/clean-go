create table users
(
    id       uuid default gen_random_uuid() not null primary key,
    name     text                           not null,
    email    text                           not null,
    password text                           not null
);

create unique index idx_users_email on users (email);