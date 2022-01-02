CREATE SCHEMA ent;

BEGIN;

CREATE TABLE IF NOT EXISTS countries
(
    id   bigint generated always as identity
        primary key,
    code text not null,
    name text not null
);

CREATE TABLE IF NOT EXISTS addresses
(
    id         bigint generated always as identity
        primary key,
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
    id              bigint generated always as identity
        primary key,
    first_name      text not null,
    middle_name     text,
    last_name       text not null,
    email           text not null unique,
    password        text not null,
    favourite_colour valid_colours default 'green'::valid_colours null
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

CREATE VIEW country_address as
select c.id,
       c.code,
       c.name,
       (
           select array_to_json(array_agg(row_to_json(addresslist.*))) as array_to_json
           from (
                    select a.*
                    from addresses a
                    where c.id = a.country_id
                ) addresslist) as address
from countries AS c;

INSERT INTO countries (code, name)
VALUES ('AU', 'Australia');
INSERT INTO countries (code, name)
VALUES ('MY', 'Malaysia');
INSERT INTO countries (code, name)
VALUES ('ID', 'Indonesia');

INSERT INTO addresses (line_1, line_2, postcode, city, state, country_id)
VALUES ('Sydney Opera House', 'Bennelong Point', 2000, 'Sydney', 'NSW', 1);
INSERT INTO addresses (line_1, line_2, postcode, city, state, country_id)
VALUES ('Petronas Twin Towers', '', 50088, 'Kuala Lumpur','Wilayah Persekutuan', 2);
INSERT INTO users (first_name, last_name, email, password)
VALUES ('John', 'Doe', 'john@example.com','$argon2id$v=19$m=16,t=2,p=1$SHVrWmRXc2tqOW5TWmVrRw$QCPRZ0MmOB/AEEMVB1LudA');
INSERT INTO users (first_name, last_name, email, password)
VALUES ('Jane', 'Doe', 'jane@example.com',  '$argon2id$v=19$m=16,t=2,p=1$UDB3RXNPd3ZEWHQ4ZTRNVg$LhHurQuz9Q9dDEG1VNzbFg');
INSERT INTO users (first_name, last_name, email, password)
VALUES ('Jake', 'Doe', 'jake@example.com',   '$argon2id$v=19$m=16,t=2,p=1$UDB3RXNPd3ZEWHQ4ZTRNVg$LhHurQuz9Q9dDEG1VNzbFg');
INSERT INTO users (first_name, last_name, email, password)
VALUES ('Alice', 'Doe', 'alice@example.com', '$argon2id$v=19$m=16,t=2,p=1$UDB3RXNPd3ZEWHQ4ZTRNVg$LhHurQuz9Q9dDEG1VNzbFg');
INSERT INTO users (first_name, last_name, email, password)
VALUES ('Bob', 'Doe', 'bob@example.com','$argon2id$v=19$m=16,t=2,p=1$UDB3RXNPd3ZEWHQ4ZTRNVg$LhHurQuz9Q9dDEG1VNzbFg');
INSERT INTO users (first_name, last_name, email, password)
VALUES ('Charlie', 'Doe', 'charlie@example.com','$argon2id$v=19$m=16,t=2,p=1$UDB3RXNPd3ZEWHQ4ZTRNVg$LhHurQuz9Q9dDEG1VNzbFg');
INSERT INTO users (first_name, last_name, email, password)
VALUES ('Duncan', 'Doe', 'duncan@example.com','$argon2id$v=19$m=16,t=2,p=1$UDB3RXNPd3ZEWHQ4ZTRNVg$LhHurQuz9Q9dDEG1VNzbFg');
INSERT INTO users (first_name, last_name, email, password)
VALUES ('Eric', 'Doe', 'eric@example.com','$argon2id$v=19$m=16,t=2,p=1$UDB3RXNPd3ZEWHQ4ZTRNVg$LhHurQuz9Q9dDEG1VNzbFg');
INSERT INTO users (first_name, last_name, email, password)
VALUES ('Finn', 'Doe', 'Finn@example.com','$argon2id$v=19$m=16,t=2,p=1$UDB3RXNPd3ZEWHQ4ZTRNVg$LhHurQuz9Q9dDEG1VNzbFg');
INSERT INTO users (first_name, last_name, email, password)
VALUES ('Garry', 'Doe', 'garry@example.com','$argon2id$v=19$m=16,t=2,p=1$UDB3RXNPd3ZEWHQ4ZTRNVg$LhHurQuz9Q9dDEG1VNzbFg');
INSERT INTO users (first_name, last_name, email, password)
VALUES ('Holden', 'Doe', 'holden@example.com','$argon2id$v=19$m=16,t=2,p=1$UDB3RXNPd3ZEWHQ4ZTRNVg$LhHurQuz9Q9dDEG1VNzbFg');
INSERT INTO users (first_name, last_name, email, password)
VALUES ('Ivy', 'Doe', 'ivy@example.com','$argon2id$v=19$m=16,t=2,p=1$UDB3RXNPd3ZEWHQ4ZTRNVg$LhHurQuz9Q9dDEG1VNzbFg');
INSERT INTO users (first_name, last_name, email, password, favourite_colour)
VALUES ('Jeff', 'Donovan', 'jeff@example.com', '$argon2id$v=19$m=16,t=2,p=1$UDB3RXNPd3ZEWHQ4ZTRNVg$LhHurQuz9Q9dDEG1VNzbFg','blue');
INSERT INTO users (first_name, last_name, email, password, favourite_colour)
VALUES ('Bruce', 'Campbell', 'bruce@example.com', '$argon2id$v=19$m=16,t=2,p=1$UDB3RXNPd3ZEWHQ4ZTRNVg$LhHurQuz9Q9dDEG1VNzbFg','blue');
INSERT INTO users (first_name, last_name, email, password, favourite_colour)
VALUES ('Gabrielle', 'Anwar', 'gabrielle@example.com', '$argon2id$v=19$m=16,t=2,p=1$UDB3RXNPd3ZEWHQ4ZTRNVg$LhHurQuz9Q9dDEG1VNzbFg','red');


INSERT INTO user_addresses (user_id, address_id) VALUES (1, 1);
INSERT INTO user_addresses (user_id, address_id) VALUES (2, 2);
INSERT INTO user_addresses (user_id, address_id) VALUES (2, 1);

COMMIT;
