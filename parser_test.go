package main

import (
	"gingen/src"
	"testing"
)

func TestParseEndPoint(t *testing.T) {
	content, err := src.ReadFile("test_data/endpoint.txt", true)
	if err != nil {
		t.Errorf("The file should be read without error but got %s", err)
	}
	endpoints, handlers := buildHandlersAndEndpoints(content)

	// ENDPOINTS
	if len(endpoints) != 3 {
		t.Errorf("The file should have 3 endpoints but got %d", len(endpoints))
	}
	if endpoints[0].Method != "GET" {
		t.Errorf("The first endpoint should have method GET but got %s", endpoints[0].Method)
	}
	if endpoints[0].Path != "/user" {
		t.Errorf("The first endpoint should have path /user but got %s", endpoints[0].Path)
	}
	if endpoints[0].HandlerID != "controllers.GetUser" {
		t.Errorf("The first endpoint should have handlerId controllers.GetUser but got %s", endpoints[0].HandlerID)
	}
	if endpoints[0].Summary != "GetUser" {
		t.Errorf("The first endpoint should have summary GetUser but got %s", endpoints[0].Summary)
	}
	if endpoints[0].Description != "Get a user based on user name" {
		t.Errorf("The first endpoint should have description Get a user based on user name but got %s", endpoints[0].Description)
	}
	if len(endpoints[0].Headers) != 1 {
		t.Errorf("The first endpoint should have 1 header but got %d", len(endpoints[0].Headers))
	}
	if endpoints[0].Headers[0].Key != "Manger" {
		t.Errorf("The first endpoint should have header Manger but got %s", endpoints[0].Headers[0].Key)
	}
	if endpoints[0].Headers[0].Description != "Description" {
		t.Errorf("The first endpoint should have header Manger with description Description but got %s", endpoints[0].Headers[0].Description)
	}
	if endpoints[0].Headers[0].IsRequired != true {
		t.Errorf("The first endpoint should have header Manger as required but got %t", endpoints[0].Headers[0].IsRequired)
	}
	if endpoints[1].Method != "POST" {
		t.Errorf("The second endpoint should have method POST but got %s", endpoints[1].Method)
	}
	if endpoints[1].Path != "/user" {
		t.Errorf("The second endpoint should have path /user but got %s", endpoints[1].Path)
	}
	if endpoints[1].HandlerID != "controllers.PostUser" {
		t.Errorf("The second endpoint should have handlerId controllers.PostUser but got %s", endpoints[1].HandlerID)
	}
	if endpoints[1].Summary != "PostUser" {
		t.Errorf("The second endpoint should have summary PostUser but got %s", endpoints[1].Summary)
	}
	if endpoints[1].Description != "Post user data" {
		t.Errorf("The second endpoint should have description Post user data but got %s", endpoints[1].Description)
	}
	if len(endpoints[1].Headers) != 1 {
		t.Errorf("The second endpoint should have 1 header but got %d", len(endpoints[1].Headers))
	}
	if endpoints[1].Headers[0].Key != "Hello" {
		t.Errorf("The second endpoint should have header Hello but got %s", endpoints[1].Headers[0].Key)
	}
	if endpoints[1].Headers[0].Description != "optional username" {
		t.Errorf("The second endpoint should have header Hello with description optional username but got %s", endpoints[1].Headers[0].Description)
	}
	if endpoints[1].Headers[0].IsRequired != false {
		t.Errorf("The second endpoint should have header Hello as optional but got %t", endpoints[1].Headers[0].IsRequired)
	}
	if endpoints[2].Method != "DELETE" {
		t.Errorf("The third endpoint should have method DELETE but got %s", endpoints[2].Method)
	}
	if endpoints[2].Path != "/user" {
		t.Errorf("The third endpoint should have path /user but got %s", endpoints[2].Path)
	}
	if endpoints[2].HandlerID != "controllers.DeleteUser" {
		t.Errorf("The third endpoint should have handlerId controllers.DeleteUser but got %s", endpoints[2].HandlerID)
	}
	if endpoints[2].Summary != "DeleteUser" {
		t.Errorf("The third endpoint should have summary DeleteUser but got %s", endpoints[2].Summary)
	}
	if endpoints[2].Description != "Delete user data" {
		t.Errorf("The third endpoint should have description Delete user data but got %s", endpoints[2].Description)
	}
	if len(endpoints[2].Headers) != 2 {
		t.Errorf("The third endpoint should have 2 header but got %d", len(endpoints[2].Headers))
	}
	if endpoints[2].Headers[0].Key != "Hello" {
		t.Errorf("The third endpoint should have header Hello but got %s", endpoints[2].Headers[0].Key)
	}
	if endpoints[2].Headers[0].Description != "optional username" {
		t.Errorf("The third endpoint should have header Hello with description optional username but got %s", endpoints[2].Headers[0].Description)
	}
	if endpoints[2].Headers[0].IsRequired != false {
		t.Errorf("The third endpoint should have header Hello as optional but got %t", endpoints[2].Headers[0].IsRequired)
	}
	if endpoints[2].Headers[1].Key != "no" {
		t.Errorf("The third endpoint should have header no but got %s", endpoints[2].Headers[1].Key)
	}
	if endpoints[2].Headers[1].Description != "test me please" {
		t.Errorf("The third endpoint should have header no with description test me please but got %s", endpoints[2].Headers[1].Description)
	}


	// HANDLERS
	if len(handlers) != 2 {
		t.Errorf("The file should have 2 handlers but got %d", len(handlers))
	}
	if handlers[0].HandlerId != "controllers.CreateUser" {
		t.Errorf("The first handler should have the ID 'controllers.CreateUser' but got '%s'", handlers[0].HandlerId)
	}


	if len(handlers[0].RequestBodys) != 1 {
		t.Errorf("The first handler should have 1 RequestBodys but got '%d'", len(handlers[0].RequestBodys))
	}
	if (handlers[0].RequestBodys[0].Description != "creates a new user based on the content of the body") {
		t.Errorf("The first RequestBodys of the first handler should have the Description 'creates a new user based on the content of the body' but got '%s'", handlers[0].RequestBodys[0].Description)
	}
	if (handlers[0].RequestBodys[0].SchemaPath != "/path/to/l afrique") {
		t.Errorf("The first RequestBodys of the first handler should have the SchemaPath '/path/to/l afrique' but got '%s'", handlers[0].RequestBodys[0].SchemaPath)
	}
	if (handlers[0].RequestBodys[0].IsRequired != true) {
		t.Errorf("The first RequestBodys of the first handler should have the IsRequired 'true' but got 'false'")
	}

	if len(handlers[0].ResponseBodys) != 2 {
		t.Errorf("The first handler should have 2 ResponseBodys but got '%d'", len(handlers[0].ResponseBodys))
	}
	if (handlers[0].ResponseBodys[0].Status != 200) {
		t.Errorf("The first ResponseBodys of the first handler should have the Status '200' but got '%d'", handlers[0].ResponseBodys[0].Status)
	}
	if (handlers[0].ResponseBodys[0].Description != "Was able to create a user") {
		t.Errorf("The first ResponseBodys of the first handler should have the Description 'Was able to create a user' but got '%s'", handlers[0].ResponseBodys[0].Description)
	}
	if (handlers[0].ResponseBodys[0].SchemaPath != "/path/to/mes couilles") {
		t.Errorf("The first ResponseBodys of the first handler should have the SchemaPath '/path/to/mes couilles' but got '%s'", handlers[0].ResponseBodys[0].SchemaPath)
	}
	if (handlers[0].ResponseBodys[1].Status != 400) {
		t.Errorf("The second ResponseBodys of the first handler should have the Status '200' but got '%d'", handlers[0].ResponseBodys[1].Status)
	}
	if (handlers[0].ResponseBodys[1].Description != "Couldn't create new user because the given email already exists in the data base") {
		t.Errorf("The second ResponseBodys of the first handler should have the Description 'Couldn't create new user because the given email already exists in the data base' but got '%s'", handlers[0].ResponseBodys[1].Description)
	}
	if (handlers[0].ResponseBodys[1].SchemaPath != "/path/to/mes couilles") {
		t.Errorf("The second ResponseBodys of the first handler should have the SchemaPath '/path/to/mes couilles' but got '%s'", handlers[0].ResponseBodys[1].SchemaPath)
	}
}
