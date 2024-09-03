////// No Loop

// package main

// import "fmt"

// func main () {
// 	var accountBalance = 1000.0

// 	fmt.Println("Welcome  to Go Bank!")
// 	fmt.Println("what do you want to do")
// 	fmt.Println("1. Check balance")
// 	fmt.Println("2. Deposit money")
// 	fmt.Println("3. Withdraw money")
// 	fmt.Println("4. Exit")

// 	var choice int
// 	fmt.Print("What do want to do: ")
// 	fmt.Scan(&choice)

// 	if choice == 1 {
// 		fmt.Println("Your balance is", accountBalance)
// 	} else if choice == 2 {
// 		fmt.Print("Your deposit: ")
// 		var depositAmount float64
// 		fmt.Scan(&depositAmount)

// 	if depositAmount <= 0 {
// 			fmt.Println("Invalid amount. Must be greater than 0.")
// 			return
// 		}

// 		accountBalance += depositAmount // accountBalance + depositAmount
// 		fmt.Println("Balance updated! New amount:", accountBalance)
// 	} else if choice == 3 {
// 		fmt.Print("How much are withdrawing: ")
// 		var withdrawalAmount float64

// 		fmt.Scan(&withdrawalAmount)

// 	if withdrawalAmount <= 0 {
// 			fmt.Println("Invalid amount. Must be greater than 0.")
// 			return
// 		}

// 		if withdrawalAmount > accountBalance {
// 			fmt.Println("Invalid amount. You can't withdraw more than you have")
// 			return
// 		}

// 		accountBalance -= withdrawalAmount // accountBalance - withDrawalAmount
// 		fmt.Println("New Balance:", accountBalance)
// 	} else {
// 		fmt.Println("Goodbye")
// 	}

// 	fmt.Println("Your choice:", choice)
// }

////////// FINITE LOOP

// package main

// import "fmt"

// func main () {
// 	var accountBalance = 1000.0

// 	fmt.Println("Welcome  to Go Bank!")

// 	for i := 0; i < 2; i++ {
// 		fmt.Println("what do you want to do")
// 		fmt.Println("1. Check balance")
// 		fmt.Println("2. Deposit money")
// 		fmt.Println("3. Withdraw money")
// 		fmt.Println("4. Exit")

// 		var choice int
// 		fmt.Print("What do want to do: ")
// 		fmt.Scan(&choice)

// 		if choice == 1 {
// 			fmt.Println("Your balance is", accountBalance)
// 		} else if choice == 2 {
// 			fmt.Print("Your deposit: ")
// 			var depositAmount float64
// 			fmt.Scan(&depositAmount)

// 		if depositAmount <= 0 {
// 				fmt.Println("Invalid amount. Must be greater than 0.")
// 				return
// 			}

// 			accountBalance += depositAmount // accountBalance + depositAmount
// 			fmt.Println("Balance updated! New amount:", accountBalance)
// 		} else if choice == 3 {
// 			fmt.Print("How much are withdrawing: ")
// 			var withdrawalAmount float64

// 			fmt.Scan(&withdrawalAmount)

// 		if withdrawalAmount <= 0 {
// 				fmt.Println("Invalid amount. Must be greater than 0.")
// 				return
// 			}

// 			if withdrawalAmount > accountBalance {
// 				fmt.Println("Invalid amount. You can't withdraw more than you have")
// 				return
// 			}

// 			accountBalance -= withdrawalAmount // accountBalance - withDrawalAmount
// 			fmt.Println("New Balance:", accountBalance)
// 		} else {
// 			fmt.Println("Goodbye")
// 		}

// 		fmt.Println("Your choice:", choice)
// 		fmt.Println("---------------------")

// 	}
// }

/////////// INFINITE LOOP - continue and break

// package main

// import "fmt"

// func main () {
// 	var accountBalance = 1000.0

// 	fmt.Println("Welcome  to Go Bank!")

// 	for {
// 		fmt.Println("what do you want to do")
// 		fmt.Println("1. Check balance")
// 		fmt.Println("2. Deposit money")
// 		fmt.Println("3. Withdraw money")
// 		fmt.Println("4. Exit")

// 		var choice int
// 		fmt.Print("What do want to do: ")
// 		fmt.Scan(&choice)

// 		if choice == 1 {
// 			fmt.Println("Your balance is", accountBalance)
// 		} else if choice == 2 {
// 			fmt.Print("Your deposit: ")
// 			var depositAmount float64
// 			fmt.Scan(&depositAmount)

// 			if depositAmount <= 0 {
// 				fmt.Println("Invalid amount. Must be greater than 0.")
// 				//return
// 				continue
// 			}

// 			accountBalance += depositAmount // accountBalance + depositAmount
// 			fmt.Println("Balance updated! New amount:", accountBalance)
// 		} else if choice == 3 {
// 			fmt.Print("How much are withdrawing: ")
// 			var withdrawalAmount float64

// 			fmt.Scan(&withdrawalAmount)

// 			if withdrawalAmount <= 0 {
// 				fmt.Println("Invalid amount. Must be greater than 0.")
// 				//return
// 				continue
// 			}

// 			if withdrawalAmount > accountBalance {
// 				fmt.Println("Invalid amount. You can't withdraw more than you have")
// 				//return
// 				continue

// 			}

// 			accountBalance -= withdrawalAmount // accountBalance - withDrawalAmount
// 			fmt.Println("New Balance:", accountBalance)
// 		} else {
// 			fmt.Println("Goodbye")
// 			//return
// 			break
// 		}

// 		fmt.Println("---------------------")
// 	}

// 	fmt.Println("Thanks for using our bank")
// }

// package main

// import "fmt"

// func main () {
// 	var accountBalance = 1000.0

// 	fmt.Println("Welcome  to Go Bank!")

// 	for {
// 		fmt.Println("what do you want to do")
// 		fmt.Println("1. Check balance")
// 		fmt.Println("2. Deposit money")
// 		fmt.Println("3. Withdraw money")
// 		fmt.Println("4. Exit")

// 		var choice int
// 		fmt.Print("What do want to do: ")
// 		fmt.Scan(&choice)

// 		switch choice {
// 			case 1:
// 				fmt.Println("Your balance is", accountBalance)
// 			case 2:
// 				fmt.Print("Your deposit: ")
// 				var depositAmount float64
// 				fmt.Scan(&depositAmount)

// 				if depositAmount <= 0 {
// 					fmt.Println("Invalid amount. Must be greater than 0.")
// 					//return
// 					continue
// 				}

// 				accountBalance += depositAmount // accountBalance + depositAmount
// 				fmt.Println("Balance updated! New amount:", accountBalance)
// 			case 3:
// 				fmt.Print("How much are withdrawing: ")
// 				var withdrawalAmount float64

// 				fmt.Scan(&withdrawalAmount)

// 				if withdrawalAmount <= 0 {
// 					fmt.Println("Invalid amount. Must be greater than 0.")
// 					//return
// 					continue
// 				}

// 				if withdrawalAmount > accountBalance {
// 					fmt.Println("Invalid amount. You can't withdraw more than you have")
// 					//return
// 					continue
// 				}
// 			default:
// 				fmt.Println("Goodbye")
// 				fmt.Println("---------------------")
// 				fmt.Println("Thanks for using our bank")
// 				return
// 				//break

// 		}
// 	}

// }

/// WRITING TO & READING FROM A FILE

// package main

// import (
// 	"fmt"
// 	"os"
// 	"strconv"
// )

// const accountBalanceFile = "balance.txt"

// func getBalanceFromFile() float64 {
// 	data, _ := os.ReadFile(accountBalanceFile)
// 	balanceText := string(data)
// 	balance, _ := strconv.ParseFloat(balanceText, 64)
// 	return balance

// 	/*
// 	Conversion was made from byte to string, and from string to float64
// 	[104 101 108 108 111] → string(data) → strconv.ParseFloat(balanceText, 64)
// 	before the value was returned
// 	*/
// }

// func writeBalanceToFile(balance float64) {
// 	balanceText := fmt.Sprint(balance)
// 	os.WriteFile(accountBalanceFile, []byte(balanceText), 0644)
// }

// /*
// - The function takes a single argument balance of type float64.
// - You use fmt.Sprint to convert the balance value to a string, which is stored in the balanceText variable.
// - Then, you use the os.WriteFile function to write the contents of balanceText to a file
//  named "balance.txt". The file permissions are set to 0644, which means the owner has
//  read and write permissions, while the group and others have
//  read-only permissions.

// s := "hello"
// b := []byte(s)
// fmt.Println(b) // Output: [104 101 108 108 111]

// */

// func main () {
// 	var accountBalance = getBalanceFromFile()

// 	fmt.Println("Welcome  to Go Bank!")

// 	for {
// 		fmt.Println("what do you want to do")
// 		fmt.Println("1. Check balance")
// 		fmt.Println("2. Deposit money")
// 		fmt.Println("3. Withdraw money")
// 		fmt.Println("4. Exit")

// 		var choice int
// 		fmt.Print("What do want to do: ")
// 		fmt.Scan(&choice)

// 		switch choice {
// 			case 1:
// 				fmt.Println("Your balance is", accountBalance)
// 			case 2:
// 				fmt.Print("Your deposit: ")
// 				var depositAmount float64
// 				fmt.Scan(&depositAmount)

// 				if depositAmount <= 0 {
// 					fmt.Println("Invalid amount. Must be greater than 0.")
// 					//return
// 					continue
// 				}

// 				accountBalance += depositAmount // accountBalance + depositAmount
// 				fmt.Println("Balance updated! New amount:", accountBalance)
// 				writeBalanceToFile(accountBalance)
// 			case 3:
// 				fmt.Print("How much are withdrawing: ")
// 				var withdrawalAmount float64

// 				fmt.Scan(&withdrawalAmount)

// 				if withdrawalAmount <= 0 {
// 					fmt.Println("Invalid amount. Must be greater than 0.")
// 					//return
// 					continue
// 				}

// 				if withdrawalAmount > accountBalance {
// 					fmt.Println("Invalid amount. You can't withdraw more than you have")
// 					//return
// 					continue
// 				}
// 				accountBalance -= withdrawalAmount
// 				fmt.Println("Balance updated! New amount:", accountBalance)
// 				writeBalanceToFile(accountBalance)
// 			default:
// 				fmt.Println("Goodbye")
// 				fmt.Println("---------------------")
// 				fmt.Println("Thanks for using our bank")
// 				return
// 				//break

// 		}
// 	}

// }

//////// ERROR HANDLING ///////////

package main

import (
	"errors"
	"fmt"
	"os"
	"strconv"
)

const accountBalanceFile = "balance.txt"

func getBalanceFromFile() (float64, error) {
	data, err := os.ReadFile(accountBalanceFile)


	if err != nil {
		return 1000, errors.New("failed to find balance file")
	}

	/*
		if the error value is not a useful value use 
		this instead[1000 as default value], 
		and [errors.New... as a custom value] 
	*/ 

	balanceText := string(data)
	balance, err := strconv.ParseFloat(balanceText, 64)

	if err != nil {
		return 1000,  errors.New("failed to parse stored balance value")
	}

	return balance, nil

	/*
 	Conversion was made from byte to string, and from string to float64
 	[104 101 108 108 111] → string(data) → strconv.ParseFloat(balanceText, 64)
 	before the value was returned
	*/
 }

func writeBalanceToFile(balance float64) {
	balanceText := fmt.Sprint(balance)
	os.WriteFile(accountBalanceFile, []byte(balanceText), 0644)
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

func main () {
	var accountBalance, err = getBalanceFromFile()

	if err != nil {
		fmt.Println("ERROR")
		fmt.Println(err)
		fmt.Println("----------------")
		// panic("Can't continue, sorry")
	}

	fmt.Println("Welcome  to Go Bank!")

	for {
		fmt.Println("what do you want to do")
		fmt.Println("1. Check balance")
		fmt.Println("2. Deposit money")
		fmt.Println("3. Withdraw money")
		fmt.Println("4. Exit")

		var choice int
		fmt.Print("What do want to do: ")
		fmt.Scan(&choice)

		switch choice {
			case 1:
				fmt.Println("Your balance is", accountBalance)
			case 2:
				fmt.Print("Your deposit: ")
				var depositAmount float64
				fmt.Scan(&depositAmount)

				if depositAmount <= 0 {
					fmt.Println("Invalid amount. Must be greater than 0.")
					//return
					continue
				}

				accountBalance += depositAmount // accountBalance + depositAmount
				fmt.Println("Balance updated! New amount:", accountBalance)
				writeBalanceToFile(accountBalance)
			case 3:
				fmt.Print("How much are withdrawing: ")
				var withdrawalAmount float64

				fmt.Scan(&withdrawalAmount)

				if withdrawalAmount <= 0 {
					fmt.Println("Invalid amount. Must be greater than 0.")
					//return
					continue
				}

				if withdrawalAmount > accountBalance {
					fmt.Println("Invalid amount. You can't withdraw more than you have")
					//return
					continue
				}
				accountBalance -= withdrawalAmount
				fmt.Println("Balance updated! New amount:", accountBalance)
				writeBalanceToFile(accountBalance)
			default:
				fmt.Println("Goodbye")
				fmt.Println("---------------------")
				fmt.Println("Thanks for using our bank")
				return
				//break

		}
	}

}

