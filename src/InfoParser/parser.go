package infoparser

import (
	"errors"
	"regexp"
)

// This regexp is used to match the info
var infoRegexp = regexp.MustCompile(`^[ |\t]*//@[ ]*(?i)(info)`)

// This regexp is used to match sub-info Title
var titleRegexp = regexp.MustCompile(`^[ |\t]*//@-[ ]*(?i)(title)[ ]*:[ ]*([A-z|0-9]*)`)

// This regexp is used to match sub-info Description
var infoDescriptionRegexp = regexp.MustCompile(`^[ |\t]*//@-[ ]*(?i)(description)[ ]*:[ ]*(.*)`)

// This regexp is used to match sub-info Version
var versionRegexp = regexp.MustCompile(`^[ |\t]*//@-[ ]*(?i)(version)[ ]*:[ ]*([0-9|\.]*)`)

func parseInfoField(comments []string) (Info, error) {
	var info Info
	for _, line := range comments {
		if titleRegexp.MatchString(line) {
			submatch := titleRegexp.FindStringSubmatch(line)
			info.Title = submatch[2]
		}
		if infoDescriptionRegexp.MatchString(line) {
			submatch := infoDescriptionRegexp.FindStringSubmatch(line)
			info.Description = submatch[2]
		}
		if versionRegexp.MatchString(line) {
			submatch := versionRegexp.FindStringSubmatch(line)
			info.Version = submatch[2]
		}
		if info.Title != "" && info.Description != "" && info.Version != "" {
			return info, nil
		}
	}
	if info.Title == "" || info.Description == "" || info.Version == "" {
		err := "Missing field: "
		if info.Title == "" {
			err += "title "
		}
		if info.Description == "" {
			err += "description "
		}
		if info.Version == "" {
			err += "version "
		}
		return Info{}, errors.New(err + "in info field")
	}
	return info, nil
}

/** @brief This function is used to parse the global information about the api from the comment
 * @param comments The comments to parse.
 * @return Info the global information about the api or an error if the info is not fully defined
 */
func ParseInfo(comments []string) (Info, error) {
	for index, line := range comments {
		if infoRegexp.MatchString(line) {
			info, err := parseInfoField(comments[index+1:])
			if err != nil {
				return Info{}, err
			} else {
				return info, nil
			}
		}
	}
	return Info{}, errors.New("no info declaration found")
}
