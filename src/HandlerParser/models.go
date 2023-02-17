package handlerparser


type Reference struct {
	SchemaPath	string `json:"$ref"`
}

type Schema struct {
	Ref Reference `json:"schema"`
}

type Content struct {
	ContentInfo map[string]Schema
}

type StatusDetails struct {
	Description	string
	Content	Content
}

type RequestBody struct {
	Description		string
	Content			Content
	IsRequired		bool
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