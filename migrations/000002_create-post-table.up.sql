create table posts
(
    id         uuid default gen_random_uuid() not null primary key,
    title      text,
    content    text,
    creator_id uuid
);

alter table posts
    add constraint fk_posts_creator_id_users_id
        foreign key (creator_id) references users (id);