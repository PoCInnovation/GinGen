package main

import (
	"encoding/json"
	"fmt"
	"gingen/src"
	endpointparser "gingen/src/EndpointParser"
	handlerparser "gingen/src/handlerParser"
)

func convert_json(data interface{}) []byte {
	content, err := json.MarshalIndent(data, "", "    ")
	if err != nil {
		fmt.Println(err)
	}
	return (content)
}

func main() {
	arguments := src.ArgumentGetter()
	src.ArgumentErrorHandler(arguments)
	content, err := src.ReadFile(arguments.InputFile, true)
	if err != nil {
		fmt.Println(err)
		return
	}
	endpoints, _ := endpointparser.ParseEndpoint(content)
	handlers := handlerparser.GetHandlers(content)
	// if err != nil {
	// 	fmt.Println(err)
	// 	return
	// }
	jsonHandlers := convert_json(handlers)
	jsonEndpoints := convert_json(endpoints)
	src.WriteFile(arguments.OutputFile, []string{string(jsonEndpoints)})
	src.WriteFile(arguments.OutputFile, []string{string(jsonHandlers)})
}
