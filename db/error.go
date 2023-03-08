package db

import (
	"errors"
	"strings"
)

func IsUniqueIndexConflictErr(err error) bool {
	return strings.Contains(err.Error(), "Error 1062: Duplicate entry")
}

var ErrBatchesStop = errors.New("batches stop")
