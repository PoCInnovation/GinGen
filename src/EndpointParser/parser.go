package endpointparser

import (
	"errors"
	"regexp"
	"strings"
)

var methodsRegexp = regexp.MustCompile(`^[ ]*//@[ ]*(?i)(method)[ ]*:[ ]*(?i)(GET|POST|PUT|HEAD|DELETE|CONNECT|OPTIONS|TRACE|PATCH)`)
var pathRegexp = regexp.MustCompile(`^[ ]*//@[ ]*(?i)(path)[ ]*:[ ]*(/.*)*`)
var handlerIdRegexp = regexp.MustCompile(`^[ ]*//@[ ]*(?i)(handlerid)[ ]*:[ ]*(.*)`)
var summaryRegexp = regexp.MustCompile(`^[ ]*//@[ ]*(?i)(summary)[ ]*:[ ]*(.*)`)
var descriptionRegexp = regexp.MustCompile(`^[ ]*//@[ ]*(?i)(description)[ ]*:[ ]*(.*)`)
var headersRegexp = regexp.MustCompile(`^[ ]*//@[ ]*(?i)(headers)`)
var subHeaderRegexp = regexp.MustCompile(`^[ ]*//@-[ ]*(.*)[ ]*:[ ]*(?i)(true|false)[ ]*,(.*)`)

func parseHeader(comments []string) ([]Header, error) {
	var headers []Header
	for _, line := range comments {
		if !subHeaderRegexp.MatchString(line) {
			if len(headers) > 0 {
				return headers, nil
			} else {
				return []Header{}, errors.New("no header found")
			}
		}
		submatch := subHeaderRegexp.FindStringSubmatch(line)
		headers = append(headers, Header{
			Key:         submatch[1],
			IsRequired:  strings.ToLower(submatch[2]) == "true",
			Description: submatch[3],
		})
	}
	if len(headers) > 0 {
		return headers, nil
	} else {
		return []Header{}, errors.New("no header found")
	}
}

// This function is used to parse the endpoint information from the comment
func ParseEndpoint(comments []string) ([]EndpointData, error) {
	var endpoints []EndpointData
	currentEndpoint := EndpointData{}
	for index, line := range comments {
		if methodsRegexp.MatchString(line) {
			submatch := methodsRegexp.FindStringSubmatch(line)
			currentEndpoint.Method = strings.ToUpper(submatch[2])
		}
		if headersRegexp.MatchString(line) {
			headers, err := parseHeader(comments[index+1:])
			if err != nil {
				return []EndpointData{}, err
			}
			currentEndpoint.Headers = append(currentEndpoint.Headers, headers...)
		}
		if pathRegexp.MatchString(line) {
			submatch := pathRegexp.FindStringSubmatch(line)
			currentEndpoint.Path = submatch[2]
		}
		if handlerIdRegexp.MatchString(line) {
			submatch := handlerIdRegexp.FindStringSubmatch(line)
			currentEndpoint.HandlerID = submatch[2]
		}
		if summaryRegexp.MatchString(line) {
			submatch := summaryRegexp.FindStringSubmatch(line)
			currentEndpoint.Summary = submatch[2]
		}
		if descriptionRegexp.MatchString(line) {
			submatch := descriptionRegexp.FindStringSubmatch(line)
			currentEndpoint.Description = submatch[2]
		}
		is_not_empty := currentEndpoint.Method != "" && currentEndpoint.Path != "" && currentEndpoint.HandlerID != "" && currentEndpoint.Summary != "" && len(currentEndpoint.Headers) != 0
		if is_not_empty {
			endpoints = append(endpoints, currentEndpoint)
			currentEndpoint = EndpointData{}
		}
	}
	is_not_empty := currentEndpoint.Method != "" && currentEndpoint.Path != "" && currentEndpoint.HandlerID != "" && currentEndpoint.Summary != "" && len(currentEndpoint.Headers) != 0
	if len(endpoints) == 0 && !is_not_empty {
		return []EndpointData{}, errors.New("no endpoint found")
	}
	return append(endpoints, currentEndpoint), nil
}
