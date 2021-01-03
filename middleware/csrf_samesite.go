// +build !go1.12

package middleware

import (
	"net/http"
)

const (
	// SameSiteNoneMode required to be redefined for Go 1.12 support (see #1524)
	SameSiteNoneMode http.SameSite = http.SameSiteNoneMode
)
