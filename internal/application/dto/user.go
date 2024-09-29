package dto

type UserDTO struct {
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
	UID int
	Name string
	Gender string
}

