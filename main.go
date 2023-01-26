package main

import (
	"fmt"
	"gingen/src"
	endpointparser "gingen/src/EndpointParser"
	handlerparser "gingen/src/HandlerParser"
	info "gingen/src/InfoParser"
)

func main() {
	arguments := src.ArgumentGetter()
	src.ArgumentErrorHandler(arguments)
	content, err := src.ReadFile(arguments.InputFile, true)
	if err != nil {
		fmt.Println(err)
		return
	}
	endpoints := endpointparser.ParseEndpoint(content)
	handlers := handlerparser.GetHandlers(content)
	info, _ := info.ParseInfo(content)
	endpointDetails := src.MergeStructs(endpoints, handlers)
	apiInfo := src.APIinfo{OpenApiVersion: "3.0.3", Info: info, Details: endpointDetails}
	jsonDetails := src.ConvertJson(apiInfo)
	src.WriteFile(arguments.OutputFile, []string{string(jsonDetails)})
	// src.ConvertDetails(src.EndpointDetails{})
}
