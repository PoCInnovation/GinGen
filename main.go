package main

import (
	"encoding/json"
	"fmt"
	"gingen/src"
	endpointparser "gingen/src/EndpointParser"
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
	content := src.ReadFile(arguments.InputFile, true)
	endpoints, err := endpointparser.ParseEndpoint(content)
	if err != nil {
		fmt.Println(err)
		return
	}
	json_content := convert_json(endpoints)
	src.WriteFile(arguments.OutputFile, []string{string(json_content)})
}
