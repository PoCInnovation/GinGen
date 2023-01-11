package handlerparser

type RequestBody struct {
	Description		string
	SchemaPath		string
	IsRequired		bool
}

type ResponseBody struct {
	Status		int
	Description	string
	SchemaPath	string

}

// This struct is used to store hander informations
type HandlerData struct {
	HandlerId		string
	RequestBodys 	[]RequestBody
	ResponseBodys 	[]ResponseBody
}