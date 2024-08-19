package entity

type User struct {
	ID int
	Name string
	Gender string
	BirthDate string
	Occupation string
	Password string
}

func NewUser() *User {
	return nil
}