package solid

// The Single Responsibility Principle states that a class or module should have only one reason to change.
//  In other words, a type should have a single responsibility, making the code easier to understand and maintain.
// Consider an example where we have a struct User and two methods, GetFullName() and Save():

type User struct {
	FirstName string
	LastName  string
}

func (u *User) GetFullName() string {
	return u.FirstName + " " + u.LastName
}

func (u *User) Save() error {
	return nil
}
