package src

import (
	"encoding/json"
	"fmt"
	endpointparser "gingen/src/EndpointParser"
	handlerparser "gingen/src/HandlerParser"
	info "gingen/src/InfoParser"
)

type APIinfo struct {
	OpenApiVersion string            `json:"openapi"`
	Info           info.Info         `json:"info"`
	Details        []EndpointDetails `json:"paths"`
}

type EndpointDetails struct {
	EndPoint  endpointparser.EndpointData  `json:"endpoint"`
	Requests  []handlerparser.RequestBody  `json:"requestBody"`
	Responses []handlerparser.ResponseBody `json:"responses"`
}

// "paths": {
// 	"/pet": {
// 		"put": {
// 			"summary": "Update an existing pet",
// 			"description": "Update an existing pet by Id",
// 			"operationId"

func ConvertDetails(details []EndpointDetails) []interface{} {
	// mapD := map[string]interface{}{"apple": 5, "test": map[string]int{"lettuce": 7}}
	// mapB, _ := json.Marshal(mapD)
	// fmt.Println(string(mapB))
	var result []interface{}
	for _, detail := range details {
		result = append(result, map[string]interface{}{detail.EndPoint.Path: map[string]interface{}{detail.EndPoint.Method: map[string]interface{}{"summary": detail.EndPoint.Summary, "description": detail.EndPoint.Description, "requestBody": detail.Requests, "responses": detail.Responses}}})
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

func MergeStructs(endpoints []endpointparser.EndpointData, handlers []handlerparser.HandlerData) []EndpointDetails {
	var endpointDetails []EndpointDetails
	for _, endpoint := range endpoints {
		endpointDetails = append(endpointDetails, EndpointDetails{EndPoint: endpoint})
		for _, handler := range handlers {
			if handler.HandlerId == endpoint.HandlerID {
				endpointDetails[len(endpointDetails)-1].Requests = append(endpointDetails[len(endpointDetails)-1].Requests, handler.RequestBodys...)
				endpointDetails[len(endpointDetails)-1].Responses = append(endpointDetails[len(endpointDetails)-1].Responses, handler.ResponseBodys...)
			}
		}
	}
	return endpointDetails
}
