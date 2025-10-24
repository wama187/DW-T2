package user

type User struct {
	ID       uint
	Name     string
	Email    string
	Password string
}

type UserRepository interface {
	Create(user *User) error
	FindByEmail(email string) (*User, error)
}
