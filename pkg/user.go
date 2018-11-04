package root

type User struct {
}

type UserService interface {
	CreateUser(u *User) error
	GetByUsername(username string) (*User, error)
}
