package filter

import (
	"net/url"
	"strconv"
	"strings"
)

const (
	paginationDefaultPage = 1
	paginationDefaultSize = 10

	queryParamPage          = "page"
	queryParamLimit         = "limit"
	queryParamOffset        = "offset"
	queryParamDisablePaging = "disable_paging"
	queryParamSort          = "sort"

	SqlAsc  = "ASC"
	SqlDesc = "DESC"
)

type Filter struct {
	Page          int  `json:"page"`
	Offset        int  `json:"offset"`
	Limit         int  `json:"size"`
	DisablePaging bool `json:"disable_paging"`

	Sort   map[string]string `json:"sort"`
	Search bool
}

func New(queries url.Values) *Filter {
	var page, limit, offset int
	page, err := strconv.Atoi(queries.Get(queryParamPage))
	if err != nil {
		page = paginationDefaultPage
	}
	limit, err = strconv.Atoi(queries.Get(queryParamLimit))
	if err != nil {
		limit = paginationDefaultSize
	}

	offset, err = strconv.Atoi(queries.Get(queryParamOffset))
	if err != nil {
		offset = limit * (page - 1) // calculates offset
	}

	disablePaging, _ := strconv.ParseBool(queries.Get(queryParamDisablePaging))

	sortKey := make(map[string]string)
	if queries.Has(queryParamSort) {
		s := queries[queryParamSort]
		for _, val := range s {
			key, _, found := strings.Cut(val, ",")
			if found {
				sortKey[key] = SqlDesc
			} else {
				sortKey[key] = SqlAsc
			}
		}
	}

	return &Filter{
		Page:          page,
		Offset:        offset,
		Limit:         limit,
		DisablePaging: disablePaging,
		Sort:          sortKey,
	}
}
