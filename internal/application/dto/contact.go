package dto

type CreateContactDTO struct {
	UID int
	Phone string
	Email string
}

type UpdateContactDTO struct {
	UID int
	Email string
	Phone string
}
