package example

import (
	"net/http"
)

func fooB() (a A, b http.Header) { _ = "STUB: not implemented"; return *new(A), *new(http.Header) }
