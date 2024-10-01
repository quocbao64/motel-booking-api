create table ward
(
    id          bigserial
        primary key,
    created_at  timestamp with time zone default CURRENT_TIMESTAMP,
    updated_at  timestamp with time zone default CURRENT_TIMESTAMP,
    deleted_at  timestamp with time zone,
    ward_name   text,
    ward_type   text,
    district_id bigint
        constraint fk_ward_district
            references district
);

alter table ward
    owner to postgres;

create index idx_ward_deleted_at
    on ward (deleted_at);

