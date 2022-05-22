package constants

import (
	"os"
	"strconv"
)

var (
	ACCCESS_TOKEN_EXPIRE_TIME, _ = strconv.ParseInt(os.Getenv("ACCESS_TOKEN_EXPIRE_TIME"), 10, 64)
	ACCESS_TOKEN_SECRET          = os.Getenv("ACCESS_TOKEN_SECRET")
)
