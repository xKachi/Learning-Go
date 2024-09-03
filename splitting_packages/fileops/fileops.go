
package fileops

import (
	"errors"
	"fmt"
	"os"
	"strconv"
)



func GetFloatFromFile(fileName string) (float64, error) {
	data, err := os.ReadFile(fileName)


	if err != nil {
		return 1000, errors.New("failed to find balance file")
	}

	/*
		if the error value is not a useful value use 
		this instead[1000 as default value], 
		and [errors.New... as a custom value] 
	*/ 

	valueText := string(data)
	value, err := strconv.ParseFloat(valueText, 64)

	if err != nil {
		return 1000,  errors.New("failed to parse stored  value")
	}

	return value, nil

	/*
 	Conversion was made from byte to string, and from string to float64
 	[104 101 108 108 111] → string(data) → strconv.ParseFloat(balanceText, 64)
 	before the value was returned
	*/
 }

func WriteFloatToFile(value float64, filename string) {
	valueText := fmt.Sprint(value)
	os.WriteFile(filename, []byte(valueText), 0644)
}

 /*
 - The function takes a single argument balance of type float64.
 - You use fmt.Sprint to convert the balance value to a string, which is stored in the balanceText variable.
 - Then, you use the os.WriteFile function to write the contents of balanceText to a file
  named "balance.txt". The file permissions are set to 0644, which means the owner has
  read and write permissions, while the group and others have
  read-only permissions.

 s := "hello"
 b := []byte(s)
fmt.Println(b) // Output: [104 101 108 108 111]

*/