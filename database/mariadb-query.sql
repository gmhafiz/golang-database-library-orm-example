-- name: ListDynamicUsers :many
SELECT id, first_name, middle_name, last_name, email, password, favourite_colour
FROM users
WHERE (@first_name = '' OR first_name LIKE '%' || @first_name || '%')
  AND (@email = '' OR email = LOWER(@email) )
ORDER BY (CASE
              WHEN @first_name_desc = 'first_name' THEN first_name
              WHEN @email_desc = 'email' THEN email
    END) DESC,
         (CASE
              WHEN @first_name_asc = 'first_name' THEN first_name
              WHEN @email_asc = 'email' THEN email
             END);
-- OFFSET ? LIMIT ? ;


-- name: CreateUser :exec
INSERT INTO users (first_name, middle_name, last_name, email, password, favourite_colour)
VALUES (?, ?, ?, ?, ?, ?);

-- name: GetUser :one
SELECT id, first_name, middle_name, last_name, email, favourite_colour
FROM users
WHERE id = ?;

-- name: ListUsers :many
SELECT id, first_name, middle_name, last_name, email, favourite_colour
FROM users
ORDER BY id
LIMIT 30
OFFSET 0;

-- name: DeleteUser :exec
DELETE
FROM users
WHERE id = ?;

-- name: UpdateUser :exec
UPDATE users
SET first_name = ?,
    middle_name = ?,
    last_name = ?,
    email = ?,
    favourite_colour = ?
WHERE id = ?;

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
select * from country_address;

-- name: SelectUsers :many
SELECT u.id, u.first_name, u.middle_name, u.last_name, u.email, u.favourite_colour
FROM users u
LIMIT 30;

-- name: SelectUserAddresses :many
SELECT DISTINCT ua.user_id, ua.address_id
FROM addresses a
         LEFT JOIN user_addresses ua ON a.id = ua.address_id
WHERE ua.user_id = ?;

-- name: SelectAddress :many
SELECT a.*
FROM addresses a
WHERE a.id = ?;

-- name: SelectWhereInLastNames :many
SELECT * FROM users WHERE last_name in (?);
