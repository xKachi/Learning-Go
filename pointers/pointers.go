//// WITHOUT A POINTER

// package main

// import "fmt"

// func main() {
// 	age := 32 // regular variable

// 	fmt.Println("Age:", age)

// 	adultYears := getAdultYears(age)
// 	fmt.Println(adultYears)
// }

// func getAdultYears(age int) int {
// 	return age - 18
// }

///// USING A POINTER as a value, and passing it to a function 
// package main

// import "fmt"

// func main() {
// 	age := 32 // regular variable

// 	var agePointer *int

// 	agePointer = &age

// 	fmt.Println("Age:", *agePointer)

// 	adultYears := getAdultYears(agePointer)
// 	fmt.Println(adultYears)
// }

// func getAdultYears(age *int) int {
// 	return *age - 18
// }



//// MUTATING VALUES WITH POINTERS
package main

import "fmt"

func main() {
	age := 32 // regular variable

	var agePointer *int

	agePointer = &age

	fmt.Println("Age:", *agePointer)

	editAgeToAdultYears(agePointer)
	fmt.Println(age)
}

func editAgeToAdultYears(age *int) {
	//return *age - 18
	*age = *age - 18
}