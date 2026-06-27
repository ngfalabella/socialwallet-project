package domain

type User struct {
	ID int
	Name string
	Email string
	Active bool
}

func CreateNewUser ( id int , name string, email string ) (User, error) {
	if email == "" {
		return User{} , EmailEmptyInvalid
	}
	return User{
		ID: id ,
		Name : name,
		Email : email ,
		Active: true,
	} , nil 
}

