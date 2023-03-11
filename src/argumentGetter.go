package src

import (
	"encoding/json"
	"fmt"
	"os"
)

// This struct is used to store the arguments passed to the program
type Argument struct {
	InputFile     string
	OutputFile    string
	ComponentFile string
	Silent        bool
}

func checkFileExisting(FilePath string, ShouldCreated bool, FileName string) {
	stat, err := os.Stat(FilePath)
	if err != nil {
		if ShouldCreated {
			fmt.Printf("The %s file will be created\n", FileName)
			fd, err := os.Create(FilePath)
			if err != nil {
				fmt.Printf("Error while creating file %s: %v\n", FileName, err)
				os.Exit(1)
			}
			fd.Close()
		} else {
			fmt.Printf("The %s file does not exist\n", FileName)
			os.Exit(1)
		}
	} else if stat.IsDir() {
		fmt.Printf("The %s file is a directory\n", FileName)
		os.Exit(1)
	}
}

/** @brief This function handles the error of the arguments passed to the program.
 * @param arguments Argument struct containing the arguments passed to the program
 * @return Argument struct containing the arguments
 */
func ArgumentErrorHandler(arguments Argument) {
	checkFileExisting(arguments.ComponentFile, false, "component")
	checkFileExisting(arguments.InputFile, false, "input")
	if arguments.OutputFile == "" {
		fmt.Println("The output file is not specified")
		os.Exit(1)
	}
	checkFileExisting(arguments.OutputFile, true, "output")
	if arguments.InputFile == arguments.OutputFile {
		fmt.Println("The input file and the output file are the same")
		os.Exit(1)
	}
	component, err := ReadFile(arguments.ComponentFile, false)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	var value interface{}
	var rawData []byte
	for _, line := range component {
		rawData = append(rawData, []byte(line)...)
	}
	if json.Unmarshal(rawData, &value) != nil {
		fmt.Println("The component file is not valid json")
		os.Exit(1)
	}
}

/** @brief This function is used to get the components from the component file
 * @param componentFile The path to the component file
 * @return A map containing the components
 */
func GetComponents(componentFile string) (map[string]interface{}, error) {
	component, err := ReadFile(componentFile, false)
	if err != nil {
		return nil, err
	}
	var value map[string]interface{}
	var rawData []byte
	for _, line := range component {
		rawData = append(rawData, []byte(line)...)
	}
	if json.Unmarshal(rawData, &value) != nil {
		return nil, fmt.Errorf("The component file is not valid json")
	}
	return value, nil
}
