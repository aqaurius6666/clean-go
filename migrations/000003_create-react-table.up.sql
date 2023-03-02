create table reacts
(
    user_id uuid,
    post_id uuid,
    type    varchar(12)
);


create unique index idx_reacts_user_id_post_id
    on reacts (user_id, post_id);

