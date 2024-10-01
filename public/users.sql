create table users
(
    id              bigserial
        primary key,
    created_at      timestamp with time zone default CURRENT_TIMESTAMP,
    updated_at      timestamp with time zone default CURRENT_TIMESTAMP,
    deleted_at      timestamp with time zone,
    full_name       text,
    email           text,
    img_url         text,
    password        text,
    phone           text not null,
    role            jsonb,
    refresh_token   text,
    identity_number text
);

alter table users
    owner to postgres;

create index idx_users_deleted_at
    on users (deleted_at);

