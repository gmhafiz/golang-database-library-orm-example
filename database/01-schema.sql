BEGIN;

create table countries
(
    id   bigint generated always as identity
        primary key,
    code text not null,
    name text not null
);

create table addresses
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

create table users
(
    id          bigint generated always as identity
        primary key,
    first_name  text not null,
    middle_name text,
    last_name   text not null,
    email       text not null unique,
    password    text not null
);

create table user_addresses
(
    user_id    bigint
        constraint user_addresses_users_id_fk
            references users ON DELETE CASCADE ,
    address_id bigint
        constraint user_addresses_addresses_id_fk
            references addresses ON DELETE CASCADE ,
    constraint user_addresses_pk
        primary key (user_id, address_id)
);

CREATE VIEW country_address as
select c.id, c.code, c.name,
       (
           select array_to_json(array_agg(row_to_json(addresslist.*))) as array_to_json
           from (
                    select a.*
                    from addresses a
                    where c.id = a.country_id
                ) addresslist) as addresses
from countries AS c;

INSERT INTO countries (code, name)
VALUES ('AU', 'Australia');
INSERT INTO countries (code, name)
VALUES ('MY', 'Malaysia');

INSERT INTO addresses (line_1, line_2, postcode, city, state, country_id)
VALUES ('Sydney Opera House', 'Bennelong Point', 2000, 'Sydney', 'NSW', 1);
INSERT INTO addresses (line_1, line_2, postcode, city, state, country_id)
VALUES ('Petronas Twin Towers', '', 50088, 'Kuala Lumpur', 'Wilayah Persekutuan', 2);

INSERT INTO users (first_name, last_name, email, password)
VALUES ('John', 'Doe', 'john@example.com', '$argon2id$v=19$m=16,t=2,p=1$SHVrWmRXc2tqOW5TWmVrRw$QCPRZ0MmOB/AEEMVB1LudA'); -- password
INSERT INTO users (first_name, last_name, email, password)
VALUES ('Jane', 'Doe', 'jane@example.com', '$argon2id$v=19$m=16,t=2,p=1$UDB3RXNPd3ZEWHQ4ZTRNVg$LhHurQuz9Q9dDEG1VNzbFg'); -- password

COMMIT;