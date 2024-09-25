///// FUNCTION AS VALUES AND FUNCTION AS TYPES /////////

// package main

// import "fmt"

// type transformFn func(int) int 

// func main() {
// 	numbers := []int{1,2,3,4}
// 	doubled := transformNumbers(&numbers, double)
// 	tripled := transformNumbers(&numbers, triple)

// 	fmt.Println(tripled)
// 	fmt.Println(doubled)
// }

// func transformNumbers(numbers *[]int, transform transformFn) []int {
// 	dNumbers := []int{}
// 	for _, val := range *numbers {
// 		dNumbers = append(dNumbers, transform(val))
// 	}

// 	return dNumbers
// }

// func double(number int) int {
// 	return number * 2
// }

// func triple(number int) int {
// 	return number * 3
// }


/////// RETURNING FUNCTION AS VALUES /////////////

// package main

// import "fmt"

// type transformFn func(int) int 

// func main() {
// 	numbers := []int{1,2,3,4}
// 	moreNumbers := []int{5, 1, 2}
// 	doubled := transformNumbers(&numbers, double)
// 	tripled := transformNumbers(&numbers, triple)

// 	fmt.Println(tripled)
// 	fmt.Println(doubled)

// 	transformerFn1 := getTransformerFunction(&numbers)
// 	transformerFn2 := getTransformerFunction(&moreNumbers)

// 	transformedNumbers := transformNumbers(&numbers, transformerFn1)
// 	moreTransformedNumbers := transformNumbers(&numbers, transformerFn2)

// 	fmt.Println(transformedNumbers)
// 	fmt.Println(moreTransformedNumbers)
// }

// func transformNumbers(numbers *[]int, transform transformFn) []int {
// 	dNumbers := []int{}
// 	for _, val := range *numbers {
// 		dNumbers = append(dNumbers, transform(val))
// 	}

// 	return dNumbers
// }

// func getTransformerFunction(numbers *[]int) transformFn {
// 	if(*numbers)[0] == 1 {
// 		return double
// 	} else {
// 		return triple
// 	}
	
// }

// func double(number int) int {
// 	return number * 2
// }

// func triple(number int) int {
// 	return number * 3
// }


//////////// ANONYMOUS FUNCTIONS ////////////

package main

import "fmt"

type transformFn func(int) int 

func main() {
	numbers := []int{1,2,3,4}
	moreNumbers := []int{5, 1, 2}
	doubled := transformNumbers(&numbers, double)
	tripled := transformNumbers(&numbers, triple)

	fmt.Println(tripled)
	fmt.Println(doubled)

	transformerFn1 := getTransformerFunction(&numbers)
	transformerFn2 := getTransformerFunction(&moreNumbers)

	transformedNumbers := transformNumbers(&numbers, transformerFn1)
	moreTransformedNumbers := transformNumbers(&numbers, transformerFn2)

	fmt.Println(transformedNumbers)
	fmt.Println(moreTransformedNumbers)
}

func transformNumbers(numbers *[]int, transform transformFn) []int {
	dNumbers := []int{}
	for _, val := range *numbers {
		dNumbers = append(dNumbers, transform(val))
	}

	return dNumbers
}

func getTransformerFunction(numbers *[]int) transformFn {
	if(*numbers)[0] == 1 {
		return double
	} else {
		return triple
	}
	
}

func double(number int) int {
	return number * 2
}

func triple(number int) int {
	return number * 3
}



