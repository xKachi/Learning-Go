//////////////////// DEFINING STRUCTS ////////////////////////////////////////////////////////////

// package main
// import (
// 	"fmt"
// 	"time"
// )

// type user struct {
// 	firstName string
// 	lastName string
// 	birthdate string
// 	createdAt time.Time
// }

// func main() {
// 	userFirstName := getUserData("Please enter your first name: ")
// 	userLastName := getUserData("Please enter your last name: ")
// 	userBirthdate := getUserData("Please enter your birthdate (MM/DD/YYYY): ")

// 	var appUser user

// 	appUser = user{
// 		firstName: userFirstName,
// 		lastName: userLastName,
// 		birthdate: userBirthdate,
// 		createdAt: time.Now(),
// 	}

// 	// ... do something awesome with that gathered data!

// 	outPutUserDetails(appUser)
// }

// func outPutUserDetails(u user) {
// 	//...
// 	fmt.Println(u.firstName, u.lastName, u.birthdate)
// }

// func getUserData(promptText string) string {
// 	fmt.Print(promptText)
// 	var value string
// 	fmt.Scan(&value)
// 	return value
// }

//////////////////////////////////////////////////////////////////////////////////////////
//// STRUCTS AND POINTERS

// package main

// import (
// 	"fmt"
// 	"time"
// )

// type user struct {
// 	firstName string
// 	lastName string
// 	birthdate string
// 	createdAt time.Time
// }

// func main() {
// 	userFirstName := getUserData("Please enter your first name: ")
// 	userLastName := getUserData("Please enter your last name: ")
// 	userBirthdate := getUserData("Please enter your birthdate (MM/DD/YYYY): ")

// 	var appUser user

// 	appUser = user{
// 		firstName: userFirstName,
// 		lastName: userLastName,
// 		birthdate: userBirthdate,
// 		createdAt: time.Now(),
// 	}

// 	// ... do something awesome with that gathered data!

// 	outPutUserDetails(appUser)
// }

// func outPutUserDetails(u user) {
// 	//...
// 	fmt.Println(u.firstName, u.lastName, u.birthdate)
// }

// func getUserData(promptText string) string {
// 	fmt.Print(promptText)
// 	var value string
// 	fmt.Scan(&value)
// 	return value
// }

///////////////////////////////////////////////////////////////////////////

//// ADDING METHODS TO STRUCTS

// package main

// import (
// 	"fmt"
// 	"time"
// )

// type user struct {
// 	firstName string
// 	lastName string
// 	birthdate string
// 	createdAt time.Time
// }

// func (u user) outPutUserDetails() {
// 	//...
// 	fmt.Println(u.firstName, u.lastName, u.birthdate)
// }

// func (u *user) clearUserName() {
// 	u.firstName = ""
// 	u.lastName = ""
// }

// func main() {
// 	userFirstName := getUserData("Please enter your first name: ")
// 	userLastName := getUserData("Please enter your last name: ")
// 	userBirthdate := getUserData("Please enter your birthdate (MM/DD/YYYY): ")

// 	var appUser user

// 	appUser = user{
// 		firstName: userFirstName,
// 		lastName: userLastName,
// 		birthdate: userBirthdate,
// 		createdAt: time.Now(),
// 	}

// 	// ... do something awesome with that gathered data!

// 	appUser.outPutUserDetails()
// 	appUser.clearUserName()
// 	appUser.outPutUserDetails()
// }

// func getUserData(promptText string) string {
// 	fmt.Print(promptText)
// 	var value string
// 	fmt.Scan(&value)
// 	return value
// }

//////////////////////////////////////////////////////////////////////////////
///// CONSTRUCTOR FUNCTION IN STRUCTS
// package main

// import (
// 	"fmt"
// 	"time"
// )

// type user struct {
// 	firstName string
// 	lastName string
// 	birthdate string
// 	createdAt time.Time
// }

// func (u user) outPutUserDetails() {
// 	//...
// 	fmt.Println(u.firstName, u.lastName, u.birthdate)
// }

// func (u *user) clearUserName() {
// 	u.firstName = ""
// 	u.lastName = ""
// }

// // constructor

// func newUser(firstName , lastName , birthdate string) *user {
// 	return &user{
// 		firstName: firstName,
// 		lastName: lastName,
// 		birthdate: birthdate,
// 		createdAt: time.Now(),
// 	}
// }

// /*
// N/B: In the above code you can also return a pointer, this will enable
// the returned value to no be copied. As seen in the preceding code.

// -- The code  below is the non-pointer version --

// func newUser(firstName , lastName , birthdate string) user {
// 	return user{
// 		firstName: firstName,
// 		lastName: lastName,
// 		birthdate: birthdate,
// 		createdAt: time.Now(),
// 	}
// }

// */

// func main() {
// 	userFirstName := getUserData("Please enter your first name: ")
// 	userLastName := getUserData("Please enter your last name: ")
// 	userBirthdate := getUserData("Please enter your birthdate (MM/DD/YYYY): ")

// 	var appUser *user

// 	appUser = newUser(userFirstName, userLastName, userBirthdate)

// 	// ... do something awesome with that gathered data!

// 	appUser.outPutUserDetails()
// 	appUser.clearUserName()
// 	appUser.outPutUserDetails()
// }

// func getUserData(promptText string) string {
// 	fmt.Print(promptText)
// 	var value string
// 	fmt.Scan(&value)
// 	return value
// }

//////////////////////////////////////////////////////////////////////////
// USING CONSTRUCTOR FUNCTIONS FOR VALIDATION

// package main

// import (
// 	"errors"
// 	"fmt"
// 	"time"
// )

// type user struct {
// 	firstName string
// 	lastName string
// 	birthdate string
// 	createdAt time.Time
// }

// func (u user) outPutUserDetails() {
// 	//...
// 	fmt.Println(u.firstName, u.lastName, u.birthdate)
// }

// func (u *user) clearUserName() {
// 	u.firstName = ""
// 	u.lastName = ""
// }

// // constructor

// func newUser(firstName , lastName , birthdate string) (*user, error) {
// 	if firstName == "" || lastName == "" || birthdate == "" {
// 		return nil, errors.New("first name, last name and birthdate are required.")
// 	}

// 	return &user{
// 		firstName: firstName,
// 		lastName: lastName,
// 		birthdate: birthdate,
// 		createdAt: time.Now(),
// 	}, nil
// }

// /*
// N/B: In the above code you can also return a pointer, this will enable
// the returned value to no be copied. As seen in the preceding code.

// -- The code  below is the non-pointer version --

// func newUser(firstName , lastName , birthdate string) user {
// 	return user{
// 		firstName: firstName,
// 		lastName: lastName,
// 		birthdate: birthdate,
// 		createdAt: time.Now(),
// 	}
// }

// */

// func main() {
// 	userFirstName := getUserData("Please enter your first name: ")
// 	userLastName := getUserData("Please enter your last name: ")
// 	userBirthdate := getUserData("Please enter your birthdate (MM/DD/YYYY): ")

// 	var appUser *user

// 	appUser, err := newUser(userFirstName, userLastName, userBirthdate)

// 	if err != nil {
// 		fmt.Println(err)
// 		return
// 	}

// 	// ... do something awesome with that gathered data!

// 	appUser.outPutUserDetails()
// 	appUser.clearUserName()
// 	appUser.outPutUserDetails()
// }

// func getUserData(promptText string) string {
// 	fmt.Print(promptText)
// 	var value string
// 	fmt.Scanln(&value)
// 	return value
// }

///////////////////////////////////////////////////////////////////

/* STRUCTS, PACKAGES, AND EXPORTS CODE*/

package main

import (
	"fmt"
	"structs/user"
)




func main() {
	userFirstName := getUserData("Please enter your first name: ")
	userLastName := getUserData("Please enter your last name: ")
	userBirthdate := getUserData("Please enter your birthdate (MM/DD/YYYY): ")

	var appUser *user.User
	

	appUser, err := user.New(userFirstName, userLastName, userBirthdate)

	if err != nil {
		fmt.Println(err)
		return 
	}

	admin := user.NewAdmin("test@example", "test123")

	admin.OutPutUserDetails()
	admin.ClearUserName()
	admin.OutPutUserDetails()


	// ... do something awesome with that gathered data!

	appUser.OutPutUserDetails()
	appUser.ClearUserName()
	appUser.OutPutUserDetails()
}

func getUserData(promptText string) string {
	fmt.Print(promptText)
	var value string
	fmt.Scanln(&value)
	return value
}
