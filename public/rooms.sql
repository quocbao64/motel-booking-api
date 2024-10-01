create table rooms
(
    id             bigserial
        primary key,
    created_at     timestamp with time zone default CURRENT_TIMESTAMP,
    updated_at     timestamp with time zone default CURRENT_TIMESTAMP,
    deleted_at     timestamp with time zone,
    title          text,
    address_id     bigint,
    acreage        bigint,
    price          numeric,
    description    text,
    date_submitted text,
    owner_id       bigint,
    max_people     bigint,
    room_type      bigint,
    deposit        numeric,
    utilities      text,
    images         jsonb
);

alter table rooms
    owner to postgres;

create index idx_rooms_deleted_at
    on rooms (deleted_at);

