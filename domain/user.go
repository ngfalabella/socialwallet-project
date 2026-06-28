package domain

type User struct {
	UserID int
	UserName string
	UserEmail string
	UserActive bool
}


func NewUserConstructor ( id int , name string, email string ) (User , error ) {
	if email == "" {
		return User{} , ErrInvalidEmail
	}
	return User{
		UserID : id,
	  UserName : name,
	  UserEmail : email,
	  UserActive : true,
	} , nil 
}