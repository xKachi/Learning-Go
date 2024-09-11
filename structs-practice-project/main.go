package main

import (
	//"errors"
	"bufio"
	"fmt"
	"os"
	"practice/note"
	"strings"
)

func main() {
	title, content := getNoteData()

	userNote, err := note.New(title, content)

	if err != nil {
		fmt.Println(err)
		return
	}

	// Do something with note
	userNote.Display()
	err = userNote.Save()

	if err !=nil {
		fmt.Println("Saving the note failed.")
		return
	}

	fmt.Println("Saving the note succeeded!")
}

func getNoteData() (string, string) {
	title := getUserInput("Note title:")

	content := getUserInput("Note content:")

	return title, content
}

func getUserInput(prompt string) string {
	fmt.Printf("%v ", prompt)
	// var value string
	//fmt.Scanln(&value);
	// Using a more complex solution for user input
	reader := bufio.NewReader(os.Stdin)

	/*
	Function signature The function is named getUserInput and takes a single argument prompt of type string. It returns a string value.

Printing the prompt The function starts by printing the prompt string to the console using fmt.Printf. The %v format verb is used to print the value of prompt.

Reading user input Instead of using the simple fmt.Scanln approach, you've opted for a more complex solution using bufio.NewReader(os.Stdin). This creates a new reader that reads from the standard input (usually the keyboard).

The ReadString method is used to read a string from the input until a newline character (\n) is encountered. The error returned by ReadString is checked, and if it's not nil, an empty string is returned.

Trimming the input The input string is then trimmed using strings.TrimSuffix to remove any trailing newline (\n) or carriage return (\r) characters.

Returning the input Finally, the trimmed input string is returned by the function.
	
	*/

	text, err := reader.ReadString('\n')

	if err != nil {
		return ""
	}

	text = strings.TrimSuffix(text, "\n")
	text = strings.TrimSuffix(text, "\r")

	return text
}