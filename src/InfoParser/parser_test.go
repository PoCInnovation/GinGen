package infoparser

import (
	"testing"
)

func TestInfoParserBasique(t *testing.T) {
	info, err := ParseInfo([]string{
		"//@info",
		"//@-title: test",
		"//@-description: description test",
		"//@-version: 1.3",
	})
	if err != nil {
		t.Errorf("The info should be parse without error but got %s", err)
	}
	if info.Title != "test" {
		t.Errorf("The title should be test but got %s", info.Title)
	}
	if info.Description != "description test" {
		t.Errorf("The description should be \"description test\" test but got %s", info.Description)
	}
	if info.Version != "1.3" {
		t.Errorf("The version should be 1.3 but got %s", info.Version)
	}
}

func TestInfoParserAdvance(t *testing.T) {
	info, err := ParseInfo([]string{
		"//@    info   	 		",
		"//@-    description: description test",
		"			//@-   version   : 1.3 87983",
		"		   	//@-    title   :     test oh auehd",
	})
	if err != nil {
		t.Errorf("The info should be parse without error but got %s", err)
	}
	if info.Title != "test" {
		t.Errorf("The title should be test but got %s", info.Title)
	}
	if info.Description != "description test" {
		t.Errorf("The description should be \"description test\" test but got %s", info.Description)
	}
	if info.Version != "1.3" {
		t.Errorf("The version should be 1.3 but got %s", info.Version)
	}
}

func TestInfoParserFail(t *testing.T) {
	info, err := ParseInfo([]string{
		"//@info",
		"//@-title: test",
		"//@-description: description test",
		"//@-version: jojoq",
	})
	if err == nil {
		t.Errorf("The info should not be parsed and produce an error but got %s", info)
	}
	info, err = ParseInfo([]string{
		"//@-title: test",
		"//@-description: description test",
		"//@-version: 1.3",
	})
	if err == nil {
		t.Errorf("The info should not be parsed and produce an error but got %s", info)
	}
}
