package handlerparser

import (
	"bufio"
	"os"
	"regexp"
	"testing"
)

var comment_regexp = regexp.MustCompile(`^[ |\t]*//.*`)

/** @brief This function is used to get the file content and return it as an array of string with one line for each case.
 * @param path The path to the file. The path MUST BE VALID.
 * @param only_comment If true, the function will return only the comments. If false, the function will return the whole file.
 * @return []string the content of the file or an error if the file cannot be read.
 */
func readFile(path string, only_comment bool) ([]string, error) {
	file, err := os.Open(path)
	if err != nil {
		return nil, err
	}
	// Close the file when the function is done
	defer file.Close()
	scanner := bufio.NewScanner(file)
	scanner.Split(bufio.ScanLines)
	var text []string
	// Read the file line by line and add it to the text slice
	for scanner.Scan() {
		line := scanner.Text()
		if only_comment && comment_regexp.MatchString(line) {
			text = append(text, line)
		} else if !only_comment {
			text = append(text, line)
		}
	}
	return text, nil
}

func TestGetHandlers(t *testing.T) {
	content, err := readFile("test_data/file_test", true)
	if err != nil {
		t.Errorf("The file should be read without error but got %s", err)
	}
	handlers := GetHandlers(content)

	if len(handlers) != 2 {
		t.Errorf("The file should have 2 handlers but got %d", len(handlers))
	}
	if handlers[0].HandlerId != "controllers.CreateUser" {
		t.Errorf("The first handler should have the ID 'controllers.CreateUser' but got '%s'", handlers[0].HandlerId)
	}

	if len(handlers[0].RequestBodys) != 1 {
		t.Errorf("The first handler should have 1 RequestBodys but got '%d'", len(handlers[0].RequestBodys))
	}
	if handlers[0].RequestBodys[0].Description != "creates a new user based on the content of the body" {
		t.Errorf("The first RequestBodys of the first handler should have the Description 'creates a new user based on the content of the body' but got '%s'", handlers[0].RequestBodys[0].Description)
	}
	if handlers[0].RequestBodys[0].Content.ContentInfo["application/json"].Ref.SchemaPath != "/path/to/l afrique" {
		t.Errorf("The first RequestBodys of the first handler should have the SchemaPath '/path/to/l afrique' but got '%s'", handlers[0].RequestBodys[0].Content.ContentInfo["application/json"].Ref.SchemaPath)
	}
	if handlers[0].RequestBodys[0].IsRequired != true {
		t.Errorf("The first RequestBodys of the first handler should have the IsRequired 'true' but got 'false'")
	}

	if len(handlers[0].ResponseBodys) != 2 {
		t.Errorf("The first handler should have 2 ResponseBodys but got '%d'", len(handlers[0].ResponseBodys))
	}
	if len(handlers[0].ResponseBodys[0].Status) != 1 {
		t.Errorf("The first ResponseBodys should have 1 statuses but is '%d'", len(handlers[0].ResponseBodys[0].Status))
	}
	if len(handlers[0].ResponseBodys[1].Status) != 1 {
		t.Errorf("The first ResponseBodys should have 1 statuses but is '%d'", len(handlers[0].ResponseBodys[1].Status))
	}

	value200, ok200 := handlers[0].ResponseBodys[0].Status[200]
	if !ok200 {
		t.Errorf("The first ResponseBodys should have a status of 200 but got nothing")
	}
	if value200.Description != "Was able to create a user" {
		t.Errorf("The first ResponseBodys of the first handler should have the Description 'Was able to create a user' but got '%s'", value200.Description)
	}
	if value200.Content.ContentInfo["application/json"].Ref.SchemaPath != "/path/to/mes couilles" {
		t.Errorf("The first ResponseBodys of the first handler should have the SchemaPath '/path/to/mes couilles' but got '%s'", value200.Content.ContentInfo["application/json"].Ref.SchemaPath)
	}
	
	value400, ok400 := handlers[0].ResponseBodys[1].Status[400]

	if !ok400 {
		t.Errorf("The second ResponseBodys of the first handler should have the Status '400' but got nothing")
	}
	if value400.Description != "Couldn't create new user because the given email already exists in the data base" {
		t.Errorf("The second ResponseBodys of the first handler should have the Description 'Couldn't create new user because the given email already exists in the data base' but got '%s'", value400.Description)
	}
	if value400.Content.ContentInfo["application/json"].Ref.SchemaPath != "/path/to/mes couilles" {
		t.Errorf("The second ResponseBodys of the first handler should have the SchemaPath '/path/to/mes couilles' but got '%s'", value400.Content.ContentInfo["application/json"].Ref.SchemaPath)
	}
}
