package main

import (
	"encoding/json"
	"fmt"
	"gingen/src"
	endpointparser "gingen/src/EndpointParser"
	info "gingen/src/InfoParser"
	handlerparser "gingen/src/HandlerParser"
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
	apiDetails := mergeStructs(endpoints, handlers)
	apiInfo := APIinfo{Info: info, Details: apiDetails}
	jsonDetails := convert_json(apiInfo)
	src.WriteFile(arguments.OutputFile, []string{string(jsonDetails)})
}
