package main

import "fmt"

func main() {

	//1

	//var hobbies [3]string = [3]string{"Reading", "Exercising", "Listening to Music"}
	//fmt.Println(hobbies)

	//2

	//fmt.Println(hobbies[0])
	//fmt.Println([2]string{hobbies[1],hobbies[2]})

	//3

	//newSlice := hobbies[:2]
	//newSlice2 := hobbies[0:2]

	//fmt.Println(newSlice, newSlice2)

	// 4

	//newSlice3 := hobbies[1:]
	

	//fmt.Println(newSlice3)

	//5

	//courseGoals := []string{"Build Project", "Write MongoDB Article"}

	//6

	// courseGoals[1] = "MongoDB Champion"

	// courseGoals = append(courseGoals, "Portfolio Project")

	// fmt.Println(courseGoals)

	//

	type Products struct {
		title string
		id int
		price float64
	}

	newProducts := []Products{
		{title:"Phone", id: 1, price: 10.99},
		{title:"Laptop", id: 2, price: 20.99},
		{title:"Charger", id: 3, price: 3.99},
}

fmt.Println(newProducts)

}

// Time to practice what you learned!

// 1) Create a new array (!) that contains three hobbies you have
// 		Output (print) that array in the command line.
// 2) Also output more data about that array:
//		- The first element (standalone)
//		- The second and third element combined as a new list
// 3) Create a slice based on the first element that contains
//		the first and second elements.
//		Create that slice in two different ways (i.e. create two slices in the end)
// 4) Re-slice the slice from (3) and change it to contain the second
//		and last element of the original array.
// 5) Create a "dynamic array" that contains your course goals (at least 2 goals)
// 6) Set the second goal to a different one AND then add a third goal to that existing dynamic array
// 7) Bonus: Create a "Product" struct with title, id, price and create a
//		dynamic list of products (at least 2 products).
//		Then add a third product to the existing list of products.