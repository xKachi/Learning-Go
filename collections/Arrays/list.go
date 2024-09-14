package list

import "fmt"

func main() {
	prices := []float64{10.99, 8.99}
	fmt.Println(prices[0:1])

	prices = append(prices, 5.99, 12.99, 22.99)
	fmt.Print(prices)

	discountPrices := []float64{101.99, 80.99, 20.59}
	prices = append(prices, discountPrices...)
	fmt.Println(prices)
}


// func main() {
// 	var productNames [4]string = [4]string{"A Book"}
// 	// we are storing a list of floating point values
// 	prices := [4]float64{10.99, 9.99, 45.99, 20.0}	
// 	// fmt.Println(productNames)
// 	// fmt.Println(prices)

// 	productNames[2] = "A Carpet"

// 	featuredPrices := prices[1:]
// 	featuredPrices[0] = 199.9
// 	highlightedPrices := featuredPrices[:1]
// 	fmt.Println(highlightedPrices)
// 	fmt.Println(prices)
// 	fmt.Println(featuredPrices)

// 	fmt.Println(len(featuredPrices), cap(featuredPrices))
// }
