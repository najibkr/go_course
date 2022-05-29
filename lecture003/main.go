// package main

// import (
// 	"fmt"

// 	"github.com/najibkr/go_course/lecture003/common"
// 	"github.com/najibkr/go_course/lecture003/store"
// )

// func main() {
// 	newUser := store.Raees{
// 		Name:     common.GetUserInput("Enter your name", false),
// 		Email:    common.GetUserInput("Enter your new email", false),
// 		Password: common.GetUserInput("Enter your new password", false),
// 	}

// 	for {
// 		fmt.Printf("Welcome to login page\n\n")
// 		correctEmail := newUser.ValidateEmail(common.GetUserInput("Enter your saved email", true))
// 		correctPass := newUser.ValidatePassword(common.GetUserInput("Enter your saved passowrd", true))
// 		if correctEmail && correctPass {
// 			fmt.Println("Welcome dear", newUser.Email)
// 			break
// 		} else {
// 			fmt.Println("Wrong email or password")
// 		}

// 	}
// }

package main

import (
	"fmt"

	"github.com/najibkr/go_course/lecture003/database"
)

func main() {
	var database database.Database = &database.MySQL{}

	// APP
	saved := database.SaveData("Najib")
	if saved {
		fmt.Println("Data saved Successfully")
	} else {
		fmt.Println("Could not save the data")
	}

	result := database.ReadData()
	fmt.Println("Saved data: ", result)

}
