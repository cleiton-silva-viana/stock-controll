package userfeature

import (
	"net/http"
	"stock-controll/internal/application/dto"
	"stock-controll/internal/application/failure"
	"stock-controll/internal/domain/factory"
	uow "stock-controll/internal/infrastructure/unit_work"
)

type UserFeature struct {
	uow               uow.IUnitOfWork
	userFactory       factory.IUserFactory
	credentialFactory factory.ICredentialFactory
	contactFactory    factory.IContactFactory
}

func NewUserFeature(
	UOW uow.IUnitOfWork,
	userFactory factory.IUserFactory,
	credentialFactory factory.ICredentialFactory,
	contactFactory factory.IContactFactory) *UserFeature {
	return &UserFeature{
		uow:               UOW,
		userFactory:       userFactory,
		credentialFactory: credentialFactory,
		contactFactory:    contactFactory,
	}
}

func (u *UserFeature) CreateUser(userData dto.CreateUserRequestDTO) (*dto.CreateUserResponseDTO, error) {
	err := u.uow.Begin()
	if err != nil {
		return nil, err
	}

	userDTO, credentialDTO, contactDTO := parseCreateUserRequestDTO(userData)

	user, errorList := u.userFactory.Create(userDTO)
	if errorList != nil {
		return nil, &failure.Error{
			Status:    http.StatusBadRequest,
			Message:   "Error data",
			ErrorList: errorList,
		}
	}

	id, err := u.uow.UserRepository().Save(*user)
	if err != nil {
		// Criar um erro default de persistência
		// Refatorar !!! - Preciso implementar um método que, a partir da persistence, retorno o possível status code e descrição do erro
		u.uow.Rollback()
		return nil, &failure.Error{
			Status:  http.StatusInternalServerError,
			Message: "internal server error",
		}
	}

	credentialDTO.ID = id
	contactDTO.ID = id

	credential, errorList := u.credentialFactory.Create(credentialDTO)
	if errorList != nil {
		u.uow.Rollback()
		return nil, &failure.Error{
			Status:    http.StatusBadRequest,
			Message:   "Error data",
			ErrorList: errorList,
		}
	}

	err = u.uow.CredentialRepository().Save(*credential)
	if err != nil {
		u.uow.Rollback()
		return nil, &failure.Error{
			Status:  http.StatusInternalServerError,
			Message: "internal server error",
		}
	}

	contact, errorList := u.contactFactory.Create(contactDTO)
	if errorList != nil {
		u.uow.Rollback()
		return nil, &failure.Error{
			Status:    http.StatusBadRequest,
			Message:   "Error data",
			ErrorList: errorList,
		}
	}

	err = u.uow.ContactRepository().Save(*contact)
	if err != nil {
		u.uow.Rollback()
		return nil, &failure.Error{
			Status:  http.StatusInternalServerError,
			Message: "internal server error",
		}
	}

	return &dto.CreateUserResponseDTO{
		ID:     id,
		Name:   userDTO.Name,
		Gender: userDTO.Gender,
	}, nil
}

func parseCreateUserRequestDTO(DTO dto.CreateUserRequestDTO) (dto.CreateUserDTO, dto.CreateCredentialDTO, dto.CreateContactDTO) {
	user := dto.CreateUserDTO{
		Name:      DTO.Name,
		Gender:    DTO.Gender,
		BirthDate: DTO.BirthDate,
	}

	credential := dto.CreateCredentialDTO{
		Password: DTO.Password,
	}

	contact := dto.CreateContactDTO{
		Email: DTO.Email,
		Phone: DTO.Phone,
	}

	return user, credential, contact
}
