-- name: ListDynamicUsers :many
SELECT id, first_name, middle_name, last_name, email, password, favourite_colour, updated_at
FROM users
WHERE (@first_name::text = '' OR first_name ILIKE '%' || @first_name || '%')
  AND (@email::text = '' OR email = LOWER(@email) )
--   AND (@favourite_colour IS NOT NULL OR favourite_colour = @favourite_colour )
--   AND (@favourite_colour_present::text = '' OR favourite_colour = @favourite_colour )
--   AND (@favourite_colour_present::valid_colours = '' OR favourite_colour = @favourite_colour )
ORDER BY (CASE
              WHEN @first_name_desc::text = 'first_name' THEN first_name
              WHEN @email_desc::text = 'email' THEN email
--               WHEN @favourite_colour_desc::text = 'favourite_colour' THEN favourite_colour
    END) DESC,
         (CASE
              WHEN @first_name_asc::text = 'first_name' THEN first_name
              WHEN @email_asc::text = 'email' THEN email
--               WHEN @favourite_colour_asc::text = 'favourite_colour' THEN favourite_colour
             END)

OFFSET @sql_offset LIMIT @sql_limit ;
-- SELECT *
-- FROM users
-- WHERE (@first_name::text = '' OR first_name = @first_name)
--   AND (@email::text = '' OR email ILIKE '%' || @email || '%')
-- --   AND (@favourite_colour::text = '' OR favourite_colour ILIKE '%' || @favourite_colour || '%')
-- LIMIT 30
-- OFFSET 0;

-- name: CreateUser :one
INSERT INTO users (first_name, middle_name, last_name, email, password, favourite_colour)
VALUES ($1, $2, $3, $4, $5, $6)
RETURNING *;

-- name: GetUser :one
SELECT id, first_name, middle_name, last_name, email, favourite_colour, updated_at
FROM users
WHERE id = $1;

-- name: ListUsers :many
SELECT id, first_name, middle_name, last_name, email, favourite_colour, updated_at
FROM users
ORDER BY id
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
    email=$4,
    favourite_colour=$5
WHERE id = $6;

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
SELECT u.id, u.first_name, u.middle_name, u.last_name, u.email, u.favourite_colour, updated_at
FROM "users" u
LIMIT 30;

-- name: SelectUserAddresses :many
SELECT DISTINCT ua.user_id, ua.address_id
FROM "addresses" a
         LEFT JOIN "user_addresses" ua ON a.id = ua.address_id
WHERE ua.user_id = ANY($1::int[]);

-- name: SelectAddress :many
SELECT a.*
FROM addresses a
WHERE a.id = ANY($1::int[]);

-- name: SelectWhereInLastNames :many
SELECT * FROM users WHERE last_name = ANY(@last_name::text[]);
