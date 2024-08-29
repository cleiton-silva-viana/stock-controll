package dto

type CreateContactDTO struct {
	ID int
	Phone string
	Email string
}

type UpdateContactDTO struct {
	ID int
	Email string
	Phone string
}
