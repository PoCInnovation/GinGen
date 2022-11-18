package src

import (
	"bufio"
	"os"
)

/** @brief This function is used to get the file content and return it as an array of string with one line for each case.
 * @param path The path to the file. The path MUST BE VALID.
 * @return []string the content of the file
 */
func ReadFile(path string) []string {
	// The error is not handled here because it is handled in the main function by the ArgumentErrorHandler() function
	file, _ := os.Open(path)
	// Close the file when the function is done
	defer file.Close()
	scanner := bufio.NewScanner(file)
	scanner.Split(bufio.ScanLines)
	var text []string
	for scanner.Scan() {
		text = append(text, scanner.Text())
	}
	return text
}
