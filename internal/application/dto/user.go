package dto

import "time"

type BaseUserDTO struct {
	Role      string
	Position  string
	FirstName string
	LastName  string
	Gender    string
}

type UserDTO struct {
	BaseUserDTO
	BirthDate time.Time
}

type CreateUserRequestDTO struct {
	BaseUserDTO
	BirthDate time.Time
	CPF       string
	Email     string
	Phone     string
	Password  string
}

type CreateUserResponseDTO struct {
	UUID string
	BaseUserDTO
}
