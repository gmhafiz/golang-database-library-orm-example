package sqlx

import (
	"fmt"
	"net/url"
)

type Filter struct {
	Address bool //  true / false
}

func Filters(queries url.Values) (*Filter, error) {
	var enabled bool

	if queries.Has("address") {
		truth := queries.Get("address")
		switch truth {
		case "true":
			enabled = true
		case "false":
			enabled = false
		default:
			return nil, fmt.Errorf("invalid value for address")
		}
	}

	return &Filter{
		Address: enabled,
	}, nil
}
