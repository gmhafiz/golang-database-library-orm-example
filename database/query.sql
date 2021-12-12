-- name: CreateUser :one
INSERT INTO users (first_name, middle_name, last_name, email, password)
VALUES ($1, $2, $3, $4, $5)
RETURNING *;

-- name: GetUser :one
SELECT *
FROM users
WHERE id = $1
LIMIT 1;

-- name: ListUsers :many
SELECT *
FROM users
ORDER BY last_name;

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