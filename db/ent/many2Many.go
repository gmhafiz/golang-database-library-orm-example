package ent

import (
	"context"

	"godb/db/ent/ent/gen"
)

// ListM2M eager loads addresses for each user.
// Generated sql:
/*
SELECT
	DISTINCT "entUsers"."id",
	"entUsers"."first_name",
	"entUsers"."middle_name",
	"entUsers"."last_name",
	"entUsers"."email",
	"entUsers"."password"
FROM "entUsers" args = []

SELECT
	"user_id",
	"address_id"
FROM "address_entUsers"
WHERE "user_id" IN ($1, $2) args = [1 2]

SELECT
	DISTINCT "entAddresses"."id",
	"entAddresses"."line_1",
	"entAddresses"."line_2",
	"entAddresses"."postcode",
	"entAddresses"."state"
	FROM "entAddresses"
WHERE "entAddresses"."id" IN ($1, $2) args = [1 2]
*/
// Note that since `Password` field is made sensitive, it will not be serialized
// when return to the user.
func (r *database) ListM2M(ctx context.Context) ([]*gen.User, error) {
	return r.db.User.Query().
		Limit(30).
		WithAddresses().
		All(ctx)
}
