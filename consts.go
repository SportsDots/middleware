package middleware

import "net/http"

const XRequestIDHeader = "X-Request-ID"

func GetXRequestIDHeaderFromRequest(request *http.Request) string {
	return request.Header.Get(XRequestIDHeader)
}

var (
	Headers = []string{
		"x-request-id",
		"authorization",
	}
)
