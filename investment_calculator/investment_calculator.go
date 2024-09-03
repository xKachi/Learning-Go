package main

import (
	"fmt"
	"math"
)

const inflationRate float64 = 2

func main() {
	var investmentAmount float64
	expectedReturnRate := 5.5
	var years float64

	//fmt.Print("Investment Amount: ")
	outputText("Investment Amount")
	fmt.Scan(&investmentAmount)

	//fmt.Print("Expected Return Rate: ")
	outputText("Expected Return Rate")
	fmt.Scan(&expectedReturnRate)

	//fmt.Print("Years: ")
	outputText("Years")
	fmt.Scan(&years)

	futureValue, futureRealValue := calculateFutureValues(investmentAmount, expectedReturnRate, years)

	//futureValue := investmentAmount * math.Pow(1 + expectedReturnRate / 100, years)
	//futureRealValue := futureValue / math.Pow(1 + inflationRate/100, years)

	fmt.Println("Future Value:",futureValue)
	fmt.Println("Future Value (adjusted for Inflation):", futureRealValue)
	
}

func outputText(text string) {
	fmt.Print(text)
}

func calculateFutureValues(investmentAmount float64, expectedReturnRate float64, years float64) (float64, float64) {
	fv := investmentAmount * math.Pow(1 + expectedReturnRate / 100, years)
	rfv := fv / math.Pow(1 + inflationRate/100, years)
	return fv, rfv
}