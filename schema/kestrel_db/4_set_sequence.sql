-- reset your sequence here because we've
-- inserted data into the tables manually,
-- so the sequence is not incremented

SELECT SETVAL(
    ( SELECT PG_GET_SERIAL_SEQUENCE('gear_set_option', 'id') ),
    ( SELECT MAX(id) FROM gear_set_option )
);
