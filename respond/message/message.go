package message

import (
	"errors"
)

var (
	ErrDBScan        = errors.New("db scanning error")
	ErrBadRequest    = errors.New("bad request")
	ErrInternalError = errors.New("internal error")

	ErrUniqueKeyViolation = errors.New("error unique key violation")
	ErrRecordNotFound     = errors.New("error record not found")
	ErrDeleting           = errors.New("error deleting")
	ErrRetrieving         = errors.New("error retrieving")

	ErrDefault = errors.New("whoops, something wrong happened")
)
