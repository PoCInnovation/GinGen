package src

import (
	"encoding/json"
	"fmt"
	endpointparser "gingen/src/EndpointParser"
	handlerparser "gingen/src/HandlerParser"
	info "gingen/src/InfoParser"
)

type APIinfo struct {
	OpenApiVersion string       `json:"openapi"`
	Info           info.Info    `json:"info"`
	Details        []APIDetails `json:"paths"`
}

type APIDetails struct {
	EndPoint endpointparser.EndpointData `json:"endpoint"`
	Handlers []handlerparser.HandlerData `json:"handlers"`
}

// "paths": {
// 	"/pet": {
// 		"put": {
// 			"summary": "Update an existing pet",
// 			"description": "Update an existing pet by Id",
// 			"operationId"

func ConvertDetails(details []APIDetails) []interface{} {
	// mapD := map[string]interface{}{"apple": 5, "test": map[string]int{"lettuce": 7}}
	// mapB, _ := json.Marshal(mapD)
	// fmt.Println(string(mapB))
	var result []interface{}
	for _, detail := range details {
		result = append(result, map[string]interface{}{detail.EndPoint.Path: map[string]interface{}{detail.EndPoint.Method: map[string]interface{}{"summary": detail.EndPoint.Summary, "description": detail.EndPoint.Description, "requestBody": detail.Handlers, "responses": detail.Handlers}}})
	}
	return result
}

func ConvertJson(data APIinfo) []byte {
	newData := map[string]interface{}{"openapi": "3.0.3", "info": data.Info, "paths": ConvertDetails(data.Details)}
	content, err := json.MarshalIndent(newData, "", "    ")
	if err != nil {
		fmt.Println(err)
	}
	return (content)
}

func MergeStructs(endpoints []endpointparser.EndpointData, handlers []handlerparser.HandlerData) []APIDetails {
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
