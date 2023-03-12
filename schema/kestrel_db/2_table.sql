create table if not exists gear_set_option (
    id              SMALLSERIAL     not null
        constraint gear_set_option_pkey
            primary key,
    name            VARCHAR         not null    default '',
    set_count       SMALLINT        not null,
    status          SMALLINT        not null    default 2,
    create_time     TIMESTAMPTZ     not null    default now(),
    update_time     TIMESTAMPTZ
);
