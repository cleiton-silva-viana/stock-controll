package contact_DTO

type CreateContactDTO struct {
	Phone string
	Email string
}

type UpdateContactDTO struct {
	ID int
	Email *string
	Phone *string
}