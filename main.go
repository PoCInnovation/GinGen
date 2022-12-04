package main

import (
	"fmt"
	endpointparser "gingen/src/EndpointParser"
)

func main() {
	// arguments := src.ArgumentGetter()
	// src.ArgumentErrorHandler(arguments)
	// content := src.ReadFile(arguments.InputFile, false)
	// comments := parseFile.GetComments(content)
	// fmt.Println(comments)
	var text = []string{
		"//@Method: GET",
		"//@Path: /user",
		"//@HandlerId: controllers.GetUser",
		"//@Summary: GetUser",
		"//@Description: Get a user based on my ass",
		"//@Headers",
		"//@-Manger: true, Description",
	}
	endpoints, err := endpointparser.ParseEndpoint(text)
	if err != nil {
		panic(err)
	}
	fmt.Println(endpoints)
}
