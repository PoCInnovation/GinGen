package endpointparser

import (
	"errors"
	"regexp"
	"strings"
)

// This regexp is used to match the endpoint http method
var methodsRegexp = regexp.MustCompile(`^[ |\t]*//@[ ]*(?i)(method)[ ]*:[ ]*(?i)(GET|POST|PUT|HEAD|DELETE|CONNECT|OPTIONS|TRACE|PATCH)`)

// This regexp is used to match the endpoint path
var pathRegexp = regexp.MustCompile(`^[ |\t]*//@[ ]*(?i)(path)[ ]*:[ ]*(/.*)*`)

// This regexp is used to match the endpoint handlerid
var handlerIdRegexp = regexp.MustCompile(`^[ |\t]*//@[ ]*(?i)(handlerid)[ ]*:[ ]*(.*)`)

// This regexp is used to match the endpoint summary
var summaryRegexp = regexp.MustCompile(`^[ |\t]*//@[ ]*(?i)(summary)[ ]*:[ ]*(.*)`)

// This regexp is used to match the endpoint description
var descriptionRegexp = regexp.MustCompile(`^[ |\t]*//@[ ]*(?i)(description)[ ]*:[ ]*(.*)`)

// This regexp is used to match the endpoint headers
var headersRegexp = regexp.MustCompile(`^[ |\t]*//@[ ]*(?i)(headers)`)

// This function is used to parse the endpoint headers sub-information from the comment
var subHeaderRegexp = regexp.MustCompile(`^[ |\t]*//@-[ ]*([A-z|1-9]*)[ ]*:[ ]*(?i)(true|false)[ ]*,[ ]*(.*)`)

/** @brief This function is used to parse the endpoint information for an headers from the comment
 * @param comments The comments to parse. This parameter is a slice of string begin at the line after the regexp HeaderRegexp match.
 * @return []Header the list of headers found in the comments params
 */
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

/** @brief This function is used to parse the endpoint information from the comment
 * @param comments The comments to parse.
 * @return []EndpointData the list of endpoint found in the comments params or an error if the endpoint is not fully defined
 */
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
				return []EndpointData{}, errors.New(err.Error() + " in endpoint " + currentEndpoint.Path)
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
	if is_not_empty {
		return append(endpoints, currentEndpoint), nil
	}
	return endpoints, nil
}
