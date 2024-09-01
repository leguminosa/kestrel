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

INSERT INTO gear_set_option (id, name, set_count, status, create_time) VALUES (1, 'ATK', 2, 1, '2023-02-25 21:09:35');
INSERT INTO gear_set_option (id, name, set_count, status, create_time) VALUES (2, 'DEF', 2, 1, '2023-02-25 21:09:35');
INSERT INTO gear_set_option (id, name, set_count, status, create_time) VALUES (3, 'HP', 2, 1, '2023-02-25 21:09:35');
INSERT INTO gear_set_option (id, name, set_count, status, create_time) VALUES (4, 'CDR', 4, 1, '2023-02-25 21:09:35');
INSERT INTO gear_set_option (id, name, set_count, status, create_time) VALUES (5, 'ASPD', 4, 1, '2023-02-25 21:09:35');
-- reset your sequence here because we've
-- inserted data into the tables manually,
-- so the sequence is not incremented

SELECT SETVAL(
    ( SELECT PG_GET_SERIAL_SEQUENCE('gear_set_option', 'id') ),
    ( SELECT MAX(id) FROM gear_set_option )
);
