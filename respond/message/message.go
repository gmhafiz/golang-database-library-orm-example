package message

import "errors"

var (
	ErrDBScan        = errors.New("db scanning error")
	ErrBadRequest    = errors.New("bad request")
	ErrInternalError = errors.New("internal error")

	ErrUpdating   = errors.New("error updating")
	ErrDeleting   = errors.New("error deleting")
	ErrRetrieving = errors.New("error retrieving")
)
