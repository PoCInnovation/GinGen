package warning

import (
	"fmt"
	"gingen/src"
	"os"
)

var Reset = "\033[0m"
var Yellow = "\033[33m"
var Green = "\033[32m"
var Blue = "\033[34m"
var Bold = "\033[1m"
var arrow = " ➜ "

func printWarning(warning string) {
	fmt.Fprint(os.Stderr, Yellow+Bold+"[WARN]"+Green+arrow+Blue+warning+Reset+"\n")
}

/*  @brief This function is used to generate and print warnings.
 *  Warning case 1: GET method should not have request body
 *  Warning case 2: No response for endpoint
 *  @param info src.APIinfo struct containing the information about the API and endpoints
 */
func CheckWarning(info src.APIinfo) {
	shouldPrint := false
	for _, detail := range info.Details {
		if detail.EndPoint.Method == "GET" {
			if len(detail.Requests) != 0 {
				printWarning("GET method should not have request body")
				shouldPrint = true
			}
		}
		if len(detail.Responses) == 0 {
			printWarning("No response for endpoint : " + detail.EndPoint.Method + " " + detail.EndPoint.Path)
			shouldPrint = true
		}
	}
	if shouldPrint {
		printWarning("Your generated specifcation may not be valid because of the above reason.")
	}
}
