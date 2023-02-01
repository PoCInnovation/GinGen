package handlerparser

import (
	"regexp"
	"strconv"
)

// This regexp is used to match the handler start
var StartRegexp = regexp.MustCompile(`^[ |\t]*//@[ |\t]*(?i)(HandlerDeclaration_Start)`)

// This regexp is used to match the handler end
var EndRegexp = regexp.MustCompile(`^[ |\t]*//@[ |\t]*(?i)(Handlerdeclaration_end)`)

// This regexp is used to match the handler handlerid
var handlerIdRegexp = regexp.MustCompile(`^[ |\t]*//@[ |\t]*(?i)(handlerid)[ |\t]*:[ |\t]*(.*)`)

// This regexp is used to match the handler requestbody
var requestBodyRegexp = regexp.MustCompile(`^[ |\t]*//@[ |\t]*(?i)(requestbody)`)

// This regexp is used to match the handler requestbody required
var requiredRegexp = regexp.MustCompile(`^[ |\t]*//@-[ |\t]*(?i)(required)[ |\t]*:[ |\t]*(?i)(true|false)`)

// This regexp is used to match the handler response
var responseRegexp = regexp.MustCompile(`^[ |\t]*//@[ |\t]*(?i)(response)`)

// This regexp is used to match the handler response status
var statusRegexp = regexp.MustCompile(`^[ |\t]*//@-[ |\t]*(?i)(status)[ |\t]*:[ |\t]*(?i)([0-9]*)`)

// This regexp is used to match the handler requestbody or response description
var descriptionRegexp = regexp.MustCompile(`^[ |\t]*//@-[ |\t]*(?i)(description)[ |\t]*:[ |\t]*(.*)`)

// This regexp is used to match the handler requestbody or response schema
var schemaRegexp = regexp.MustCompile(`^[ |\t]*//@-[ |\t]*(?i)(schema)[ |\t]*:[ |\t]*(.*)`)

/** @brief This function iterates trough the comments passed as parametter and creates a new RequestBody based on the comments
 * @param comments The comments to parse
 * @return RequestBody a RequestBody created by the informations
 */
func getReqBody(comments []string) RequestBody {
	var redBody RequestBody

	for _, line := range comments {
		foundComment := false
		if descriptionRegexp.MatchString(line) {
			submatch := descriptionRegexp.FindStringSubmatch(line)
			redBody.Description = submatch[2]
			foundComment = true
		}
		if schemaRegexp.MatchString(line) {
			submatch := schemaRegexp.FindStringSubmatch(line)
			redBody.SchemaPath = submatch[2]
			foundComment = true
		}
		if requiredRegexp.MatchString(line) {
			submatch := requiredRegexp.FindStringSubmatch(line)
			redBody.IsRequired = submatch[2] == "true"
			foundComment = true
		}

		if !foundComment {
			break
		}
	}
	return redBody
}

/** @brief This function iterates trough the comments passed as parametter and creates a new ResponseBody based on the comments
 * @param comments The comments to parse
 * @return ResponseBody a ResponseBody created by the informations
 */
func getResBody(comments []string) ResponseBody {
	var resBody ResponseBody

	for _, line := range comments {
		foundComment := false
		if descriptionRegexp.MatchString(line) {
			submatch := descriptionRegexp.FindStringSubmatch(line)
			resBody.Description = submatch[2]
			foundComment = true
		}
		if schemaRegexp.MatchString(line) {
			submatch := schemaRegexp.FindStringSubmatch(line)
			resBody.SchemaPath = submatch[2]
			foundComment = true
		}
		if statusRegexp.MatchString(line) {
			submatch := statusRegexp.FindStringSubmatch(line)
			resBody.Status, _ = strconv.Atoi(submatch[2])
			foundComment = true
		}

		if !foundComment {
			break
		}
	}
	return resBody
}

/** @brief This function iterates trough the comments passed as parametter and creates a new handler based on the comments
 * @param comments The comments to parse
 * @return HandlerData a handlers created by the informations
 */
func HandlerParser(comments []string) HandlerData {
	handler := HandlerData{}
	for index, line := range comments {
		if handlerIdRegexp.MatchString(line) {
			submatch := handlerIdRegexp.FindStringSubmatch(line)
			handler.HandlerId = submatch[2]
		}
		if requestBodyRegexp.MatchString(line) {
			handler.RequestBodys = append(handler.RequestBodys, getReqBody(comments[index+1:]))
		}
		if responseRegexp.MatchString(line) {
			handler.ResponseBodys = append(handler.ResponseBodys, getResBody(comments[index+1:]))
		}
		if EndRegexp.MatchString(line) {
			break
		}
	}
	return handler
}
