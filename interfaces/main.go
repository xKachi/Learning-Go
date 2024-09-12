/*
- An interface can be said to be a contract that guarantees that a
certain value(typically a struct) has a certain method.

- An interface does not define the logic of a method, they simply define
the existence of a certain method by stating, it's NAME and RETURN value.

- It can also define the kind input values a method will accept
as parameters.

- An interface is also a type which means it can be used wherever a type is needed.


*/

package main

import (
	//"errors"
	"bufio"
	"fmt"
	"os"
	"strings"

	"example.com/note/note"
	"example.com/note/todo"
)

type saver interface {
	Save() error 
	/* A valid value must have Save() method, which takes
	no arguments and returns an error*/
	
}

// type displayer interface {
// 	Display()
// }

// embedded interface
type outputtable interface {
	saver
	Display()

	/*
	Every value of the outputtable interface must have all the 
	methods of the saver interface(we are embedding saver) and
	it must also have the display method
	*/
}

// type outputtable interface {
// 	Save() error
// 	Display()
// }

func main() {
	title, content := getNoteData()
	todoText := getUserInput("Todo text: ")

	todo, err := todo.New(todoText)

	if err != nil {
		fmt.Println(err)
		return
	}

	err = outputData(todo)

	if err != nil {
		return
	}

	userNote, err := note.New(title, content)

	if err != nil {
		fmt.Println(err)
		return
	}

	/*
	No need to handle error since function already ends here
	*/
  outputData(userNote)
}


func outputData(data outputtable) error  {
	data.Display()
	return saveData(data)
}

func saveData(data saver) error {
	err := data.Save()

	if err != nil {
		fmt.Println("Saving the note failed.")
		return err
	}

	fmt.Println("Saving the note succeeded!")
	return nil
}

func getTodoData() string {
	return getUserInput("Todo text: ")
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



