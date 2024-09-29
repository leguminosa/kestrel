SET TIMEZONE = 'Asia/Jakarta';

create type character_rarity_type as enum('N', 'R', 'SR', 'SSR', 'ASSR');
create type character_faction_type as enum('COUNTER', 'SOLDIER', 'MECH');
create type character_status_type as enum('ACTIVE', 'INACTIVE', 'DELETED');

create table if not exists characters (
    id BIGSERIAL not null constraint characters_pkey primary key,
    rarity character_rarity_type not null,
    faction character_faction_type not null,
    name VARCHAR(255) not null default '',
    cost SMALLINT not null default 0,
    status character_status_type not null default 'INACTIVE',
    created_at TIMESTAMPTZ not null default now(),
    updated_at TIMESTAMPTZ
);
insert into characters (id, rarity, faction, name, cost, status, created_at) values
(1, 'SSR', 'COUNTER', 'Xiao Lin', 4, 'ACTIVE', '2023-03-23 21:09:35+07'),
(2, 'SSR', 'COUNTER', 'Kestrel Xiao Lin', 4, 'ACTIVE', '2023-03-23 21:09:35+07'),
(3, 'SSR', 'COUNTER', 'Nest Keeper Xiao Lin', 4, 'ACTIVE', '2023-03-23 21:09:35+07')
;
SELECT SETVAL(
    ( SELECT PG_GET_SERIAL_SEQUENCE('characters', 'id') ),
    ( SELECT MAX(id) FROM characters )
);
