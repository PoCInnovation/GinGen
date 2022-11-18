package main

import (
	"bufio"
	"encoding/json"
	"fmt"
	"os"
)

func convert_json(data interface{}) []byte {
	content, err := json.Marshal(data)
	if err != nil {
		fmt.Println(err)
	}
	return (content)
}

func createOutputFile(path string) {
	_, err := os.Stat(path)
	if err != nil {
		fmt.Println("The output file does not exist")
		fd, err := os.Create(path)
		if err != nil {
			fmt.Printf("Error creating file: %v\n", err)
			os.Exit(1)
		}
		fmt.Println("The output file has been created")
		fd.Close()
	}
}

func argumentHandler() string {
	var input string
	if len(os.Args) == 1 {
		reader := bufio.NewReader(os.Stdin)
		fmt.Print("Enter the input file: ")
		input, _ = reader.ReadString('\n')
		input = input[:len(input)-1]
		createOutputFile(input)
		return input
	} else {
		input = os.Args[1]
		createOutputFile(input)
		return os.Args[1]
	}
}

type People []Person

type Job struct {
	Company string
	Role    string
}

type Person struct {
	Name string
	Age  int
	Job  Job
}

func main() {
	path := argumentHandler()
	person := Person{
		Name: "John",
		Age:  30,
		Job: Job{
			Company: "Google",
			Role:    "Software Engineer",
		},
	}
	people := People{person, person, person}
	content := convert_json(people)
	fd, err := os.OpenFile(path, os.O_APPEND|os.O_WRONLY, 0644)
	if err != nil {
		fmt.Println(err)
	}
	defer fd.Close()
	if _, err = fd.WriteString(string(content)); err != nil {
		fmt.Println(err)
	}
}
