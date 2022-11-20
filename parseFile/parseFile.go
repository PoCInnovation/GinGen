package parseFile

import (
	"fmt"
	"strings"
)

type Comment struct {
    Brief string
    FuncName string
    Params []string
    Return string
}

var clearComment = &Comment{}

/** @brief Clears a Comment type variable
 * @param comment *Comment - The Comment variable to be cleared
 */
func (comment *Comment) Reset() {
    *comment = *clearComment
}

/** @brief Gets the functions documentation and returns it as a slice of the struct Comment
 * @param content []string - The content of a file to be parsed
 * @return []Comment - The Comment struct slice containing for each element the documentation of a fonction
 */
func GetComments(content []string) []Comment {
	readComment := false
	var comments []Comment
	var newComment Comment

	for _, line := range content {
		fmt.Println(line)
		if (strings.HasPrefix(line, " */")) {
			comments = append(comments, newComment)
			readComment = false
		}
		if (strings.HasPrefix(line, "/*")) {
			readComment = true
			newComment.Params = nil
		}
		if (strings.HasPrefix(line, "func ")) {
			newComment.FuncName = line
		}
		if (readComment && strings.Contains(line, "@param")) {
			newComment.Params = append(newComment.Params, line)
		}
		if (readComment && strings.Contains(line, "@brief")) {
			newComment.Brief = line
		}
		if (readComment && strings.Contains(line, "@return")) {
			newComment.Return = line
		}
	}
	return comments
}