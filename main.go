package main

import (
	"encoding/json"
	"fmt"
	"gingen/src"
	endpointparser "gingen/src/EndpointParser"
	handlerparser "gingen/src/HandlerParser"
	info "gingen/src/InfoParser"
	warning "gingen/src/Warning"
)

type APIinfo struct {
	Info    info.Info
	Details []APIDetails
}

type APIDetails struct {
	EndPoint endpointparser.EndpointData
	Handlers []handlerparser.HandlerData
}

func convert_json(data interface{}) []byte {
	content, err := json.MarshalIndent(data, "", "    ")
	if err != nil {
		fmt.Println(err)
	}
	return (content)
}

func mergeStructs(endpoints []endpointparser.EndpointData, handlers []handlerparser.HandlerData) []APIDetails {
	var apiDetails []APIDetails
	for _, endpoint := range endpoints {
		apiDetails = append(apiDetails, APIDetails{EndPoint: endpoint})
		for _, handler := range handlers {
			if handler.HandlerId == endpoint.HandlerID {
				apiDetails[len(apiDetails)-1].Handlers = append(apiDetails[len(apiDetails)-1].Handlers, handler)
			}
		}
	}
	return apiDetails
}

func buildHandlersAndEndpoints(comments []string) ([]endpointparser.EndpointData, []handlerparser.HandlerData) {
	var endpoints []endpointparser.EndpointData
	var handlers []handlerparser.HandlerData
	for index, line := range comments {
		if endpointparser.StartRegexp.MatchString(line) {
			endpoints = append(endpoints, endpointparser.ParseOneEndpoint(comments[index+1:]))
		}
		if handlerparser.StartRegexp.MatchString(line) {
			handlers = append(handlers, handlerparser.HandlerParser(comments[index+1:]))
		}
	}
	return endpoints, handlers
}

func main() {
	arguments := src.ArgumentGetter()
	src.ArgumentErrorHandler(arguments)
	content, err := src.ReadFile(arguments.InputFile, true)
	if err != nil {
		fmt.Println(err)
		return
	}
	endpoints, handlers := buildHandlersAndEndpoints(content)
	info, _ := info.ParseInfo(content)
	endpointDetails := src.MergeStructs(endpoints, handlers)
	apiInfo := src.APIinfo{OpenApiVersion: "3.0.3", Info: info, Details: endpointDetails}
	warning.CheckWarning(apiInfo)
	jsonDetails := src.ConvertJson(apiInfo)
	src.WriteFile(arguments.OutputFile, []string{string(jsonDetails)})
}
