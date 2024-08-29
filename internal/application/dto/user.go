package dto

type CreateUserDTO struct {
	Name      string
	Gender    string
	BirthDate string
}

type CreateUserRequestDTO struct {
	Name      string
	Gender    string
	BirthDate string
	Email     string
	Phone     string
	Password  string
}

type CreateUserResponseDTO struct {
	ID int
	Name string
	Gender string
}