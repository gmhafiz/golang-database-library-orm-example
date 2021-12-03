package param

import (
	"errors"
	"net/http"
	"strconv"

	"github.com/go-chi/chi/v5"
)

var ErrParam = errors.New("error parsing param")

func UInt64(w http.ResponseWriter, r *http.Request, param string) (uint64, error) {
	val, err := strconv.ParseInt(chi.URLParam(r, param), 10, 64)
	if err != nil {
		return 0, ErrParam
	}

	return uint64(val), nil
}

func Int64(w http.ResponseWriter, r *http.Request, param string) (int64, error) {
	val, err := strconv.ParseInt(chi.URLParam(r, param), 10, 64)
	if err != nil {
		return 0, ErrParam
	}

	return val, nil
}

func String(r *http.Request, param string) string {
	return chi.URLParam(r, param)
}
