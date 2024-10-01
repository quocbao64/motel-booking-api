create table province
(
    id            bigserial
        primary key,
    created_at    timestamp with time zone default CURRENT_TIMESTAMP,
    updated_at    timestamp with time zone default CURRENT_TIMESTAMP,
    deleted_at    timestamp with time zone,
    province_name text,
    province_type text
);

alter table province
    owner to postgres;

create index idx_province_deleted_at
    on province (deleted_at);

