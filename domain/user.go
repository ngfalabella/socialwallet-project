package domain

type User struct {
	UserID int 
	UserName string
	UserEmail string
	UserActive bool
}

func NewUser( id int , name string ,email string ) (User,error) {
	if name == "" {
		return User{} , InvalidNameUser
	}
	if email == "" {
		return User{} , InvalidEmailUser
	}
	return User{
		UserID: id,
		UserName: name,
		UserEmail: email,
		UserActive: true,
	 } , nil
}