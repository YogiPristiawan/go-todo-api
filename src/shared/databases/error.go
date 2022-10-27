package databases

import (
	"errors"

	"github.com/jackc/pgx/v5"
)

// WrapError convert database error type
// into http response code
func WrapError(err error) (code int) {
	if err != nil {
		if errors.Is(err, pgx.ErrNoRows) {
			return 404
		}
		return 500
	}
	return
}
