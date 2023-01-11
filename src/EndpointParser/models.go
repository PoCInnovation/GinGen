package endpointparser

// This enum is used to store the different types of HTTP methods
const (
	Get     = "GET"
	Post    = "POST"
	Head    = "HEAD"
	Put     = "PUT"
	Delete  = "DELETE"
	Connect = "CONNECT"
	Options = "OPTIONS"
	Trace   = "TRACE"
	Patch   = "PATCH"
)

// This struct is used to store the header part of the endpoint
type Header struct {
	Key         string
	IsRequired  bool
	Description string
}

// This struct is used to store endpoint information
type EndpointData struct {
	Method      string
	Path        string
	HandlerID   string
	Summary     string
	Description string
	Headers     []Header
}
