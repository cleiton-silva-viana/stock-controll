package dto

type CreateContactDTO struct {
	UID string
	Phone string
	Device string
	Email string
}

type UpdateContactDTO struct {
	UID string
	Email string
	Phone string
}
