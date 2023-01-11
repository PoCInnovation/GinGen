package src

import (
	"os"
	"testing"
)

func TestReadFile(t *testing.T) {
	content, err := ReadFile("test_data/text.txt", false)
	if err != nil {
		t.Errorf("The file should be read without error but got %s", err)
	}
	if len(content) != 3 {
		t.Errorf("The file should have 3 lines but got %d", len(content))
	}
	if content[0] != "Hello," {
		t.Errorf("The first line should be 'Hello,' but got '%s'", content[0])
	}
	if content[1] != "// I am a test !" {
		t.Errorf("The second line should be 'I am a test !' but got '%s'", content[1])
	}
	if content[2] != "How are you ?" {
		t.Errorf("The third line should be 'How are you ?' but got '%s'", content[2])
	}
}

func TestReadFileComment(t *testing.T) {
	content, err := ReadFile("test_data/text.txt", true)
	if err != nil {
		t.Errorf("The file should be read without error but got %s", err)
	}
	if len(content) != 1 {
		t.Errorf("The file should have 1 lines but got %d", len(content))
	}
	if content[0] != "// I am a test !" {
		t.Errorf("The first line should be '// I am a test !' but got '%s'", content[0])
	}
}

func TestWriteFile(t *testing.T) {
	content := []string{"Hello,", "I am a test !", "How are you ?"}
	err := WriteFile("test_data/write.txt", content)
	if err != nil {
		t.Errorf("The file should be written without error but got %s", err)
	}
	file, err := ReadFile("test_data/write.txt", false)
	if err != nil {
		t.Errorf("The file should be read without error but got %s", err)
	}
	if len(file) != 3 {
		t.Errorf("The file should have 3 lines but got %d", len(file))
	}
	if file[0] != "Hello," {
		t.Errorf("The first line should be 'Hello,' but got '%s'", file[0])
	}
	if file[1] != "I am a test !" {
		t.Errorf("The second line should be 'I am a test !' but got '%s'", file[1])
	}
	if file[2] != "How are you ?" {
		t.Errorf("The third line should be 'How are you ?' but got '%s'", file[2])
	}
	os.Remove("test_data/write.txt")
}

func TestArgumentGetter(t *testing.T) {
	os.Args = []string{"gingen", "test_data/argument.txt", "test_data/output.txt"}
	argument := ArgumentGetter()
	if argument.InputFile != "test_data/argument.txt" {
		t.Errorf("The input file should be 'test_data/argument.txt' but got '%s'", argument.InputFile)
	}
	if argument.OutputFile != "test_data/output.txt" {
		t.Errorf("The output file should be 'test_data/output.txt' but got '%s'", argument.OutputFile)
	}
}

func TestArgumentErrorHandler(t *testing.T) {
	os.Args = []string{"gingen", "test_data/argument.txt", "test_data/output.txt"}
	argument := ArgumentGetter()
	ArgumentErrorHandler(argument)
	file, err := os.Open("test_data/output.txt")
	file.Close()
	if err != nil {
		t.Errorf("The file should be open without error but got %s", err)
	}
	os.Remove("test_data/output.txt")
}
