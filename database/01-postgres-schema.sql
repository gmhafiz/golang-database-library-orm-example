-- CREATE SCHEMA ent;

BEGIN;

CREATE TABLE IF NOT EXISTS countries
(
    id   BIGINT GENERATED ALWAYS AS IDENTITY PRIMARY KEY,
    code text not null,
    name text not null
);

CREATE TABLE IF NOT EXISTS addresses
(
    id         BIGINT GENERATED ALWAYS AS IDENTITY PRIMARY KEY,
    line_1     text not null,
    line_2     text,
    postcode   int,
    city       text,
    state      text,
    country_id bigint
        constraint addresses_countries_id_fk
            references countries ON DELETE CASCADE
);

CREATE TYPE valid_colours AS ENUM ('red', 'green', 'blue');
CREATE TABLE IF NOT EXISTS users
(
    id               BIGINT GENERATED ALWAYS AS IDENTITY PRIMARY KEY,
    first_name       text        not null,
    middle_name      text,
    last_name        text        not null,
    email            text        not null unique,
    password         text        not null,
    favourite_colour valid_colours        default 'green'::valid_colours not null,
    updated_at       timestamptz not null default NOW()
);

CREATE TABLE IF NOT EXISTS user_addresses
(
    user_id    bigint
        constraint user_addresses_users_id_fk
            references users ON DELETE CASCADE,
    address_id bigint
        constraint user_addresses_addresses_id_fk
            references addresses ON DELETE CASCADE,
    constraint user_addresses_pk
        primary key (user_id, address_id)
);

CREATE OR REPLACE FUNCTION trigger_set_updated_at()
    RETURNS TRIGGER AS
$$
BEGIN
    NEW.updated_at = NOW();
    RETURN NEW;
END;

$$ LANGUAGE plpgsql;

CREATE TRIGGER users_updated_at
    AFTER UPDATE
    on users
    FOR EACH ROW
EXECUTE PROCEDURE trigger_set_updated_at();

INSERT INTO countries (code, name)
VALUES ('AU', 'Australia');
INSERT INTO countries (code, name)
VALUES ('MY', 'Malaysia');
INSERT INTO countries (code, name)
VALUES ('ID', 'Indonesia');

INSERT INTO addresses (line_1, line_2, postcode, city, state, country_id)
VALUES ('Sydney Opera House', 'Bennelong Point', 2000, 'Sydney', 'NSW', 1);
INSERT INTO addresses (line_1, line_2, postcode, city, state, country_id)
VALUES ('Petronas Twin Towers', '', 50088, 'Kuala Lumpur',
        'Wilayah Persekutuan', 2);
INSERT INTO users (first_name, last_name, email, password, favourite_colour)
VALUES ('John', 'Doe', 'john@example.com',
        '$argon2id$v=19$m=16,t=2,p=1$SHVrWmRXc2tqOW5TWmVrRw$QCPRZ0MmOB/AEEMVB1LudA',
        'blue');
INSERT INTO users (first_name, last_name, email, password, favourite_colour)
VALUES ('Jane', 'Doe', 'jane@example.com',
        '$argon2id$v=19$m=16,t=2,p=1$UDB3RXNPd3ZEWHQ4ZTRNVg$LhHurQuz9Q9dDEG1VNzbFg',
        'blue');
INSERT INTO users (first_name, last_name, email, password, favourite_colour)
VALUES ('Jake', 'Doe', 'jake@example.com',
        '$argon2id$v=19$m=16,t=2,p=1$UDB3RXNPd3ZEWHQ4ZTRNVg$LhHurQuz9Q9dDEG1VNzbFg',
        'blue');
INSERT INTO users (first_name, last_name, email, password, favourite_colour)
VALUES ('Alice', 'Doe', 'alice@example.com',
        '$argon2id$v=19$m=16,t=2,p=1$UDB3RXNPd3ZEWHQ4ZTRNVg$LhHurQuz9Q9dDEG1VNzbFg',
        'blue');
INSERT INTO users (first_name, last_name, email, password, favourite_colour)
VALUES ('Bob', 'Doe', 'bob@example.com',
        '$argon2id$v=19$m=16,t=2,p=1$UDB3RXNPd3ZEWHQ4ZTRNVg$LhHurQuz9Q9dDEG1VNzbFg',
        'blue');
INSERT INTO users (first_name, last_name, email, password, favourite_colour)
VALUES ('Charlie', 'Doe', 'charlie@example.com',
        '$argon2id$v=19$m=16,t=2,p=1$UDB3RXNPd3ZEWHQ4ZTRNVg$LhHurQuz9Q9dDEG1VNzbFg',
        'blue');
INSERT INTO users (first_name, last_name, email, password, favourite_colour)
VALUES ('Duncan', 'Doe', 'duncan@example.com',
        '$argon2id$v=19$m=16,t=2,p=1$UDB3RXNPd3ZEWHQ4ZTRNVg$LhHurQuz9Q9dDEG1VNzbFg',
        'blue');
INSERT INTO users (first_name, last_name, email, password, favourite_colour)
VALUES ('Eric', 'Doe', 'eric@example.com',
        '$argon2id$v=19$m=16,t=2,p=1$UDB3RXNPd3ZEWHQ4ZTRNVg$LhHurQuz9Q9dDEG1VNzbFg',
        'blue');
INSERT INTO users (first_name, last_name, email, password, favourite_colour)
VALUES ('Finn', 'Doe', 'Finn@example.com',
        '$argon2id$v=19$m=16,t=2,p=1$UDB3RXNPd3ZEWHQ4ZTRNVg$LhHurQuz9Q9dDEG1VNzbFg',
        'blue');
INSERT INTO users (first_name, last_name, email, password, favourite_colour)
VALUES ('Garry', 'Doe', 'garry@example.com',
        '$argon2id$v=19$m=16,t=2,p=1$UDB3RXNPd3ZEWHQ4ZTRNVg$LhHurQuz9Q9dDEG1VNzbFg',
        'blue');
INSERT INTO users (first_name, last_name, email, password, favourite_colour)
VALUES ('Holden', 'Doe', 'holden@example.com',
        '$argon2id$v=19$m=16,t=2,p=1$UDB3RXNPd3ZEWHQ4ZTRNVg$LhHurQuz9Q9dDEG1VNzbFg',
        'blue');
INSERT INTO users (first_name, last_name, email, password, favourite_colour)
VALUES ('Ivy', 'Doe', 'ivy@example.com',
        '$argon2id$v=19$m=16,t=2,p=1$UDB3RXNPd3ZEWHQ4ZTRNVg$LhHurQuz9Q9dDEG1VNzbFg',
        'blue');
INSERT INTO users (first_name, last_name, email, password, favourite_colour)
VALUES ('Jeff', 'Donovan', 'jeff@example.com',
        '$argon2id$v=19$m=16,t=2,p=1$UDB3RXNPd3ZEWHQ4ZTRNVg$LhHurQuz9Q9dDEG1VNzbFg',
        'blue');
INSERT INTO users (first_name, last_name, email, password, favourite_colour)
VALUES ('Bruce', 'Campbell', 'bruce@example.com',
        '$argon2id$v=19$m=16,t=2,p=1$UDB3RXNPd3ZEWHQ4ZTRNVg$LhHurQuz9Q9dDEG1VNzbFg',
        'blue');
INSERT INTO users (first_name, last_name, email, password, favourite_colour)
VALUES ('Gabrielle', 'Anwar', 'gabrielle@example.com',
        '$argon2id$v=19$m=16,t=2,p=1$UDB3RXNPd3ZEWHQ4ZTRNVg$LhHurQuz9Q9dDEG1VNzbFg',
        'red');


INSERT INTO user_addresses (user_id, address_id)
VALUES (1, 1);
INSERT INTO user_addresses (user_id, address_id)
VALUES (2, 2);
INSERT INTO user_addresses (user_id, address_id)
VALUES (2, 1);

CREATE VIEW country_address as
select c.id,
       c.code,
       c.name,
       (select array_to_json(array_agg(row_to_json(addresslist.*))) as array_to_json
        from (select a.*
              from addresses a
              where c.id = a.country_id) addresslist) as address
from countries AS c;

CREATE COLLATION case_insensitive (provider = icu, locale = 'und-u-ks-level2', deterministic = false);

COMMIT;
