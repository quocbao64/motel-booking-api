create table district
(
    id            bigserial
        primary key,
    created_at    timestamp with time zone default CURRENT_TIMESTAMP,
    updated_at    timestamp with time zone default CURRENT_TIMESTAMP,
    deleted_at    timestamp with time zone,
    district_name text,
    district_type text,
    province_id   bigint
        constraint fk_district_province
            references province
);

alter table district
    owner to postgres;

create index idx_district_deleted_at
    on district (deleted_at);

