package param

import (
	"errors"
	"net/http"
	"net/url"
	"strconv"
	"strings"

	"github.com/go-chi/chi/v5"
)

var ErrParam = errors.New("error parsing param")

func UInt64(r *http.Request, param string) (uint64, error) {
	val, err := strconv.ParseInt(chi.URLParam(r, param), 10, 64)
	if err != nil {
		return 0, ErrParam
	}

	return uint64(val), nil
}

func Int64(r *http.Request, param string) (int64, error) {
	val, err := strconv.ParseInt(chi.URLParam(r, param), 10, 64)
	if err != nil {
		return 0, ErrParam
	}

	return val, nil
}

func String(r *http.Request, param string) string {
	return chi.URLParam(r, param)
}

// ToStrSlice turn comma separated query param to str slice
func ToStrSlice(v url.Values, s string) []string {
	values := v[s]
	if len(values) == 0 {
		return []string{}
	}

	str := strings.Split(values[0], ",")

	ints := make([]string, 0)

	for _, val := range str {
		ints = append(ints, val)
	}

	return ints
}
