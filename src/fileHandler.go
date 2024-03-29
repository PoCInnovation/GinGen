package src

import (
	"bufio"
	"os"
	"regexp"
)

var comment_regexp = regexp.MustCompile(`^[ |\t]*//.*`)

/** @brief This function is used to get the file content and return it as an array of string with one line for each case.
 * @param path The path to the file. The path MUST BE VALID.
 * @param only_comment If true, the function will return only the comments. If false, the function will return the whole file.
 * @return []string the content of the file or an error if the file cannot be read.
 */
func ReadFile(path string, only_comment bool) ([]string, error) {
	file, err := os.Open(path)
	if err != nil {
		return nil, err
	}
	// Close the file when the function is done
	defer file.Close()
	scanner := bufio.NewScanner(file)
	scanner.Split(bufio.ScanLines)
	var text []string
	// Read the file line by line and add it to the text slice
	for scanner.Scan() {
		line := scanner.Text()
		if only_comment && comment_regexp.MatchString(line) {
			text = append(text, line)
		} else if !only_comment {
			text = append(text, line)
		}
	}
	return text, nil
}

/** @brief This function is used to write content in the file path given in parameter.
 * @param path The path to the file.
 * @param content The content to write in the file.
 * @param ShouldCreated If true, the file will be created if it does not exist. If false, the file will be appended if it exists.
 * @return []string the content of the file
 */
func WriteFile(path string, content []string, ShouldCreated bool) error {
	var file *os.File
	var err error
	if ShouldCreated {
		file, err = os.Create(path)
	} else {
		file, err = os.OpenFile(path, os.O_APPEND|os.O_WRONLY, 0600)
	}
	if err != nil {
		return err
	}
	defer func() error {
		err := file.Close()
		if err != nil {
			return err
		}
		return nil
	}()
	for _, line := range content {
		_, err := file.WriteString(line + "\n")
		if err != nil {
			return err
		}
	}
	return nil
}
