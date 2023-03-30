create table posts
(
    id         uuid default gen_random_uuid() not null primary key,
    title      text,
    content    text,
    creator_id uuid
);

create unique index idx_posts_creator_id on posts (creator_id);