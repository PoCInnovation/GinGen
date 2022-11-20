package main

import (
	"fmt"
	"gingen/parseFile"
	"gingen/src"
)

// func convert_json(data interface{}) []byte {
// 	content, err := json.Marshal(data)
// 	if err != nil {
// 		fmt.Println(err)
// 	}
// 	return (content)
// }

func main() {
	arguments := src.ArgumentGetter()
	src.ArgumentErrorHandler(arguments)
	content := src.ReadFile(arguments.InputFile)
	comments := parseFile.GetComments(content)
	fmt.Println(comments)

}
