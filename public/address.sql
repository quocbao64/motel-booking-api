create table address
(
    id         bigserial
        primary key,
    created_at timestamp with time zone default CURRENT_TIMESTAMP,
    updated_at timestamp with time zone default CURRENT_TIMESTAMP,
    deleted_at timestamp with time zone,
    ward_id    bigint
        constraint fk_ward_addresses
            references ward,
    detail     text,
    user_id    bigint
        constraint fk_users_address
            references users
);

alter table address
    owner to postgres;

create index idx_address_deleted_at
    on address (deleted_at);

