package src

import (
	"bufio"
	"fmt"
	"os"
)

type Argument struct {
	InputFile  string
	OutputFile string
}

/** @brief This function is used to get one argument from the command line. If the argument is not passed, the user will be prompted to enter it.
 * @param index The index of the argument
 * @param prompt The prompt to display to the user
 * @return string containing the argument
 */
func getOneArgument(index int, prompt string) string {
	// This slice is to skip the first arg which is the program full path
	if len(os.Args[1:]) <= index {
		reader := bufio.NewReader(os.Stdin)
		fmt.Print(prompt)
		input, err := reader.ReadString('\n')
		if err != nil {
			fmt.Println(err)
		}
		input = input[:len(input)-1]
		return input
	} else {
		return os.Args[1:][index]
	}
}

/** @brief This function handles the arguments passed to the program. The first argument is the input file and the second argument is the output file.
 * @return Argument struct containing the arguments
 */
func ArgumentGetter() Argument {
	argument := Argument{
		InputFile:  getOneArgument(0, "Please enter the input file path: "),
		// OutputFile: getOneArgument(1, "Please enter the output file path: "),
	}
	return argument
}

/** @brief This function handles the error of the arguments passed to the program.
 * @param arguments Argument struct containing the arguments passed to the program
 * @return Argument struct containing the arguments
 */
func ArgumentErrorHandler(arguments Argument) {
	stat, err := os.Stat(arguments.InputFile)
	if err != nil {
		fmt.Println("The input file does not exist")
		os.Exit(1)
	}
	if stat.IsDir() {
		fmt.Println("The input file is a directory")
		os.Exit(1)
	}
	// stat, err = os.Stat(arguments.OutputFile)
	// if err != nil {
	// 	fmt.Println("The output file will be created")
	// 	fd, err := os.Create(arguments.OutputFile)
	// 	if err != nil {
	// 		fmt.Printf("Error creating file: %v\n", err)
	// 		os.Exit(1)
	// 	}
	// 	fd.Close()
	// } else if stat.IsDir() {
	// 	fmt.Println("The output file is a directory")
	// 	os.Exit(1)
	// }
	// if arguments.InputFile == arguments.OutputFile {
	// 	fmt.Println("The input file and the output file are the same")
	// 	os.Exit(1)
	// }
}
