package merrors

var errorType = struct {
	BadRequest         string
	server             string
	Unauthorized       string
	conflict           string
	ServiceUnavailable string
	Forbidden          string
	Downstream         string
	NotFound           string
	TooManyRequests    string
}{
	BadRequest:         "bad request",
	server:             "server",
	Unauthorized:       "unauthorized",
	conflict:           "conflict",
	ServiceUnavailable: "service unavailable",
	Forbidden:          "forbidden",
	Downstream:         "downstream",
	NotFound:           "not found",
	TooManyRequests: "too many requests",
}
