package error

import (
	"strings"
	"strconv"
)

type RequestError struct {
	StatusCode int
	Exist      bool
}

func (err RequestError) Error() string {
	return strings.Join([]string{"status code error:", strconv.Itoa(err.StatusCode)}, " ")
}