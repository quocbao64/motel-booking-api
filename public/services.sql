create table services
(
    id         bigserial
        primary key,
    created_at timestamp with time zone default CURRENT_TIMESTAMP,
    updated_at timestamp with time zone default CURRENT_TIMESTAMP,
    deleted_at timestamp with time zone,
    name       text,
    icon_url   text,
    price      numeric
);

alter table services
    owner to postgres;

create index idx_services_deleted_at
    on services (deleted_at);

