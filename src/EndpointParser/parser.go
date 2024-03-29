package endpointparser

import (
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
var endpointDescriptionRegexp = regexp.MustCompile(`^[ |\t]*//@[ ]*(?i)(description)[ ]*:[ ]*(.*)`)

// This regexp is used to match the endpoint headers
var headersRegexp = regexp.MustCompile(`^[ |\t]*//@[ ]*(?i)(headers)`)

// This function is used to parse the endpoint headers sub-information from the comment
var subHeaderRegexp = regexp.MustCompile(`^[ |\t]*//@-[ ]*([A-z|0-9]*)[ ]*:[ ]*(?i)(true|false)[ ]*,[ ]*(.*)`)

// This regexp is used to match the handler start
var StartRegexp = regexp.MustCompile(`^[ |\t]*//@[ |\t]*(?i)(EndPointDeclaration_Start)`)

// This regexp is used to match the handler end
var EndRegexp = regexp.MustCompile(`^[ |\t]*//@[ |\t]*(?i)(EndPointDeclaration_end)`)

/** @brief This function is used to parse the endpoint information for an headers from the comment
 * @param comments The comments to parse. This parameter is a slice of string begin at the line after the regexp HeaderRegexp match.
 * @return []Header the list of headers found in the comments params
 */
func parseHeader(comments []string) []Header {
	var headers []Header
	for _, line := range comments {
		if !subHeaderRegexp.MatchString(line) {
			return headers
		}
		submatch := subHeaderRegexp.FindStringSubmatch(line)
		headers = append(headers, Header{
			Key:         submatch[1],
			IsRequired:  strings.ToLower(submatch[2]) == "true",
			Description: submatch[3],
		})
	}
	return headers
}

func ParseOneEndpoint(comments []string) EndpointData {
	currentEndpoint := EndpointData{}
	for index, line := range comments {
		if methodsRegexp.MatchString(line) {
			submatch := methodsRegexp.FindStringSubmatch(line)
			currentEndpoint.Method = strings.ToUpper(submatch[2])
		}
		if headersRegexp.MatchString(line) {
			headers := parseHeader(comments[index+1:])
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
		if endpointDescriptionRegexp.MatchString(line) {
			submatch := endpointDescriptionRegexp.FindStringSubmatch(line)
			currentEndpoint.Description = submatch[2]
		}
		if EndRegexp.MatchString(line) {
			break
		}
	}
	return currentEndpoint
}
