package main

import "fmt"


func main () {
	/* We want Go to understand that when we pass integers to the add(), 
	we want the return values to also be integers, we will do this with generics*/
	result := add(1, 2)

	fmt.Println(result)
}

////////// Dynamic Types & Generics ////////


func printSomething(value any) {
	// switch value.(type) {
	// case int:
	// 	fmt.Println("Integer:", value)
	// case float64:
	// 	fmt.Println("Float:", value)
	// case string:
	// 	fmt.Println(value)
	// }
	// fmt.Println(value)

	// Alternative syntax ///

	intVal, ok := value.(int)

	if ok {
		fmt.Println("Integer:", intVal)
		return
	}

	floatVal, ok := value.(int)

	if ok {
		fmt.Println("Float:", floatVal)
		return
	}

	stringVal, ok := value.(int)

	if ok {
		fmt.Println("String:", stringVal)
		return
	}
}


// This function allows us to dynamically add values of different types
/* func add(a, b interface{}) interface{} {
	aInt, aIsInt := a.(int)
	bInt, bIsInt := b.(int)

	if aIsInt && bIsInt {
		return aInt + bInt
	}

	aFloat, aIsFloat := a.(int)
	bFloat, bIsFloat := b.(int)

	if aIsFloat && bIsFloat {
		return aFloat+ bFloat
	}


	aString, aIsString := a.(int)
	bString, bIsString := b.(int)

	if aIsString && bIsString {
		return aString + bString
	}

	return nil
}

*/


///// Rewriting the dynamic function with generics /////

/*
We are telling Go that the concrete type of values being received/returned in the add() function
will only be set and known when the add() function is called, which is done in the main() function

We are telling Go that T can be ANY of the following type int |  float64 | string
*/
func add[T int| float64 | string](a, b T) T  {
	return a + b
}