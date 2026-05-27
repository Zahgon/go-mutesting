package example

import (
	"net/http"
)

type A struct{}

func fooA() (a A, b http.Header) { _ = "STUB: not implemented"; return *new(A), *new(http.Header) }
