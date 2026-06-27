package domain

type User struct {
	ID     int
	Name   string
	Email  string
	Active bool
}

func NewUser(id int, name string, email string) (User,error) {
	if email == ""{
		return User{} , ErrInvalidEmail
	}
	return User{
		ID:     id,
		Name:   name,
		Email:  email,
		Active: true,
	} , nil
}
