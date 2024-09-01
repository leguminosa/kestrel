create table if not exists characters (
    id BIGSERIAL not null constraint characters_pkey primary key,
    character_rarity_id SMALLINT not null default 0,
    character_faction_id SMALLINT not null default 0,
    name VARCHAR(255) not null default '',
    cost SMALLINT not null default 0
);
insert into characters (id, character_rarity_id, character_faction_id, name, cost) values
(1, 4, 1, 'Xiao Lin', 4),
(2, 4, 1, 'Kestrel Xiao Lin', 4),
(3, 4, 1, 'Nest Keeper Xiao Lin', 4)
;
SELECT SETVAL(
    ( SELECT PG_GET_SERIAL_SEQUENCE('characters', 'id') ),
    ( SELECT MAX(id) FROM characters )
);

create table if not exists character_rarities (
    id SMALLSERIAL not null constraint character_rarities_pkey primary key,
    code VARCHAR(255) not null default '',
    name VARCHAR(255) not null default '',
    weight SMALLINT not null default 0
);
insert into character_rarities (id, code, name, weight) values
(1, 'N', 'Normal', 1),
(2, 'R', 'Rare', 2),
(3, 'SR', 'Super Rare', 3),
(4, 'SSR', 'Super Super Rare', 4),
(5, 'ASSR', 'Awakened SSR', 5)
;
SELECT SETVAL(
    ( SELECT PG_GET_SERIAL_SEQUENCE('character_rarities', 'id') ),
    ( SELECT MAX(id) FROM character_rarities )
);

create table if not exists character_factions (
    id SMALLSERIAL not null constraint character_factions_pkey primary key,
    name VARCHAR(255) not null default ''
);
insert into character_factions (id, name) values
(1, 'Counter'),
(2, 'Soldier'),
(3, 'Mech')
;
SELECT SETVAL(
    ( SELECT PG_GET_SERIAL_SEQUENCE('character_factions', 'id') ),
    ( SELECT MAX(id) FROM character_factions )
);
