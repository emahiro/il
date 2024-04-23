package application

const (
	HttpVersion = "Http/1.0"

	HttpStatusOK HttpStatus = iota
	HttpStatusCreated
	HttpStatusNotFound
	HttpStatusInternalServerError
)

type HttpRequest struct {
	Method  string
	URI     string
	Version string
	Headers map[string]string
	Body    string
}

type HttpResponse struct {
	Version string
	Status  string
	Headers map[string]string
	Body    string
}

type HttpStatus int

func (s HttpStatus) String() string {
	switch s {
	case HttpStatusOK:
		return "200 OK"
	case HttpStatusCreated:
		return "201 Created"
	case HttpStatusNotFound:
		return "404 Not Found"
	case HttpStatusInternalServerError:
		return "500 Internal Server Error"
	default:
		return "500 Internal Server Error"
	}
}

func ParseHttpRequest(raw string) (*HttpRequest, error) {
	return nil, nil
}
