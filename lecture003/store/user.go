package store

type User interface {
	ValidateEmail(enteredEmail string) bool
	ValidatePassword(enteredPassword string) bool
}

type Raees struct {
	Name     string
	Email    string
	Password string
}

func (r Raees) ValidateEmail(enteredEmail string) bool {
	return true
}
func (r Raees) ValidatePassword(enteredPassword string) bool {
	return true
}

// type User struct {
// 	Name     string
// 	Email    string
// 	Password string
// }

// func (user User) ValidateEmail(enteredEmail string) bool {
// 	return user.Email == enteredEmail
// }

// func (user User) ValidatePassword(enteredPassword string) bool {
// 	return user.Password == enteredPassword
// }
