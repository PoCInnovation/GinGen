package handlerparser

// import (
// 	"gingen/src"
// 	"testing"

// )

// func TestGetHandlers(t *testing.T) {
// 	content, err := src.ReadFile("test_data/file_test", true)
// 	if err != nil {
// 		t.Errorf("The file should be read without error but got %s", err)
// 	}
// 	handlers := GetHandlers(content)

// 	if len(handlers) != 2 {
// 		t.Errorf("The file should have 2 handlers but got %d", len(handlers))
// 	}
// 	if handlers[0].HandlerId != "controllers.CreateUser" {
// 		t.Errorf("The first handler should have the ID 'controllers.CreateUser' but got '%s'", handlers[0].HandlerId)
// 	}


// 	if len(handlers[0].RequestBodys) != 1 {
// 		t.Errorf("The first handler should have 1 RequestBodys but got '%d'", len(handlers[0].RequestBodys))
// 	}
// 	if (handlers[0].RequestBodys[0].Description != "creates a new user based on the content of the body") {
// 		t.Errorf("The first RequestBodys of the first handler should have the Description 'creates a new user based on the content of the body' but got '%s'", handlers[0].RequestBodys[0].Description)
// 	}
// 	if (handlers[0].RequestBodys[0].SchemaPath != "/path/to/l afrique") {
// 		t.Errorf("The first RequestBodys of the first handler should have the SchemaPath '/path/to/l afrique' but got '%s'", handlers[0].RequestBodys[0].SchemaPath)
// 	}
// 	if (handlers[0].RequestBodys[0].IsRequired != true) {
// 		t.Errorf("The first RequestBodys of the first handler should have the IsRequired 'true' but got 'false'")
// 	}

// 	if len(handlers[0].ResponseBodys) != 2 {
// 		t.Errorf("The first handler should have 2 ResponseBodys but got '%d'", len(handlers[0].ResponseBodys))
// 	}
// 	if (handlers[0].ResponseBodys[0].Status != 200) {
// 		t.Errorf("The first ResponseBodys of the first handler should have the Status '200' but got '%d'", handlers[0].ResponseBodys[0].Status)
// 	}
// 	if (handlers[0].ResponseBodys[0].Description != "Was able to create a user") {
// 		t.Errorf("The first ResponseBodys of the first handler should have the Description 'Was able to create a user' but got '%s'", handlers[0].ResponseBodys[0].Description)
// 	}
// 	if (handlers[0].ResponseBodys[0].SchemaPath != "/path/to/mes couilles") {
// 		t.Errorf("The first ResponseBodys of the first handler should have the SchemaPath '/path/to/mes couilles' but got '%s'", handlers[0].ResponseBodys[0].SchemaPath)
// 	}
// 	if (handlers[0].ResponseBodys[1].Status != 400) {
// 		t.Errorf("The second ResponseBodys of the first handler should have the Status '200' but got '%d'", handlers[0].ResponseBodys[1].Status)
// 	}
// 	if (handlers[0].ResponseBodys[1].Description != "Couldn't create new user because the given email already exists in the data base") {
// 		t.Errorf("The second ResponseBodys of the first handler should have the Description 'Couldn't create new user because the given email already exists in the data base' but got '%s'", handlers[0].ResponseBodys[1].Description)
// 	}
// 	if (handlers[0].ResponseBodys[1].SchemaPath != "/path/to/mes couilles") {
// 		t.Errorf("The second ResponseBodys of the first handler should have the SchemaPath '/path/to/mes couilles' but got '%s'", handlers[0].ResponseBodys[1].SchemaPath)
// 	}
// }
