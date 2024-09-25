///// Making Sense of Recursion

package main

import "fmt"

func main() {
	fact := factorial(5)
	fmt.Println(fact)
}

// factorial of 5 => 5 * 4 * 3 * 2 * 1 => 120

func factorial(number int) int {
	//-- Recursive Approach --//

	if number == 0 {
		return 1
	}
	return number * factorial(number - 1)
	


	//-- Loop based approach --//
	// result := 1

	// for i :=1; i <= number; i++ {
	// 	result = result * i
	// }

	// return result
}

