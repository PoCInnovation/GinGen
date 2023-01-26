package handlerparser

type RequestBody struct {
	Description		string
	SchemaPath		string
	IsRequired		bool
}

type StatusDetails struct {
	Description	string
	SchemaPath	string
}

type ResponseBody struct {
	Status		map[int]StatusDetails
}

// This struct is used to store hander informations
type HandlerData struct {
	HandlerId		string
	RequestBodys 	[]RequestBody
	ResponseBodys 	[]ResponseBody
}