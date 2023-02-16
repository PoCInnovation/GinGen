package src

import (
	"encoding/json"
	"fmt"
	endpointparser "gingen/src/EndpointParser"
	handlerparser "gingen/src/HandlerParser"
	info "gingen/src/InfoParser"
	"strconv"
	"strings"
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

// mapD := map[string]interface{}{"apple": 5, "test": map[string]int{"lettuce": 7}}
// mapB, _ := json.Marshal(mapD)
// fmt.Println(string(mapB))

func convertRequest(requestBodys []handlerparser.RequestBody) map[string]interface{} {
	result := make(map[string]interface{})
	for _, requestBody := range requestBodys {
		result["description"] = requestBody.Description
		result["required"] = requestBody.IsRequired
		result["content"] = requestBody.SchemaPath
	}
	return result
}

func convertResponse(responses []handlerparser.ResponseBody) map[string]interface{} {
	result := make(map[string]interface{})
	for _, response := range responses {
		for key, value := range response.Status {
			result[strconv.Itoa(key)] = map[string]interface{}{"description": value.Description, "content": value.SchemaPath}
		}
	}
	return result
}

func convertDetails(details []EndpointDetails) map[string]interface{} {
	result := make(map[string]interface{})
	for _, detail := range details {
		if result[detail.EndPoint.Path] != nil {
			result[detail.EndPoint.Path].(map[string]interface{})[strings.ToLower(detail.EndPoint.Method)] = map[string]interface{}{"summary": detail.EndPoint.Summary, "description": detail.EndPoint.Description}
		} else {
			result[detail.EndPoint.Path] = map[string]interface{}{strings.ToLower(detail.EndPoint.Method): map[string]interface{}{"summary": detail.EndPoint.Summary, "description": detail.EndPoint.Description}}
		}
		requestResult := convertRequest(detail.Requests)
		if len(requestResult) != 0 {
			result[detail.EndPoint.Path].(map[string]interface{})[strings.ToLower(detail.EndPoint.Method)].(map[string]interface{})["requestBody"] = requestResult
		}
		responseResult := convertResponse(detail.Responses)
		if len(responseResult) != 0 {
			result[detail.EndPoint.Path].(map[string]interface{})[strings.ToLower(detail.EndPoint.Method)].(map[string]interface{})["responses"] = responseResult
		}
	}
	return result
}

/** @brief This function is used to convert the api INFO in the right json forma
 *  @param data: The data to convert in json
 *  @return []byte the result of the marshal conversion
 */
func ConvertJson(data APIinfo) []byte {
	newData := map[string]interface{}{"openapi": "3.0.3", "info": data.Info, "paths": convertDetails(data.Details)}
	content, err := json.MarshalIndent(newData, "", "    ")
	if err != nil {
		fmt.Println(err)
	}
	return (content)
}

/** @brief This function is used to merge the endpoint and handler data
 *  @param endpoints: The endpoint data
 *  @param handlers: The handler data
 *  @return []EndpointDetails the result of the merge
 */
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
