-- name: CreateUser :one
INSERT INTO users (first_name, middle_name, last_name, email, password)
VALUES ($1, $2, $3, $4, $5)
RETURNING *;

-- name: GetUser :one
SELECT id, first_name, middle_name, last_name, email
FROM users
WHERE id = $1;

-- name: ListUsers :many
SELECT id, first_name, middle_name, last_name, email
FROM users
LIMIT 30
OFFSET 0;

-- name: DeleteUser :exec
DELETE
FROM users
WHERE id = $1;

-- name: UpdateUser :exec
UPDATE users
SET first_name=$1,
    middle_name=$2,
    last_name=$3,
    email=$4
WHERE id = $5;

-- name: CountriesWithAddress :many
SELECT c.id AS country_id,
       c.name,
       c.code,
       a.id AS address_id,
       a.line_1,
       a.line_2,
       a.postcode,
       a.city,
       a.state
FROM countries c
         LEFT JOIN addresses a on c.id = a.country_id
ORDER BY c.id;


-- name: CountriesWithAddressAggregate :many
select row_to_json(row) from (select * from country_address) row;

-- name: SelectUsers :many
SELECT u.id, u.first_name, u.middle_name, u.last_name, u.email
FROM "users" u
LIMIT 30;

-- name: SelectUserAddresses :many
SELECT DISTINCT ua.user_id, ua.address_id
FROM "addresses" a
         LEFT JOIN "user_addresses" ua ON a.id = ua.address_id
WHERE ua.user_id = ANY($1::int[]);;

-- name: SelectAddress :many
SELECT a.*
FROM addresses a
WHERE a.id = ANY($1::int[]);