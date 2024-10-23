package userfeature

import (
	"stock-controll/internal/application/dto"
	"stock-controll/internal/domain/entity"
	"stock-controll/internal/domain/factory"
	"stock-controll/internal/infrastructure/repository"
	uow "stock-controll/internal/infrastructure/unit_work"
)

type UserFeature struct {
	uow                  uow.IUnitOfWork
	userFactory          factory.IFactory[dto.UserDTO, entity.User]
	credentialFactory    factory.IFactory[dto.CreateCredentialDTO, entity.Credential]
	contactFactory       factory.IFactory[dto.CreateContactDTO, entity.Contact]
	userRepository       repository.IRepository[entity.User]
	credentialRepository repository.IRepository[entity.Credential]
	contactRepository    repository.IRepository[entity.Contact]
}

func (uf *UserFeature) CreateUser(userData dto.CreateUserRequestDTO) (*dto.CreateUserResponseDTO, error) {
	userDTO, credentialDTO, contactDTO := parseCreateUserRequestDTO(userData)

	user, err := uf.userFactory.Create(userDTO)
	if err != nil {
		return nil, err
	}

	userUID := user.GetUID()
	credentialDTO.UID = userUID
	contactDTO.UID = userUID

	credential, err := uf.credentialFactory.Create(credentialDTO)
	if err != nil {
		return nil, err
	}

	contact, err := uf.contactFactory.Create(contactDTO)
	if err != nil {
		return nil, err
	}

	err = uf.uow.Begin()
	if err != nil {
		return nil, err
	}

	err = uf.userRepository.Save(*user)
	if err != nil {
		uf.uow.Rollback()
		return nil, err
	}

	err = uf.userRepository.Save(*user)
	if err != nil {
		uf.uow.Rollback()
		return nil, err
	}

	err = uf.credentialRepository.Save(*credential)
	if err != nil {
		uf.uow.Rollback()
		return nil, err
	}
	err = uf.contactRepository.Save(*contact)
	if err != nil {
		uf.uow.Rollback()
		return nil, err
	}

	err = uf.uow.Commit()
	if err != nil {
		return nil, err
	}

	return &dto.CreateUserResponseDTO{
		UID:    userUID,
		Name:   userDTO.Name,
		Gender: userDTO.Gender,
	}, nil
}

type userFeatureBuilder struct {
	uow                  uow.IUnitOfWork
	userFactory          factory.IFactory[dto.UserDTO, entity.User]
	credentialFactory    factory.IFactory[dto.CreateCredentialDTO, entity.Credential]
	contactFactory       factory.IFactory[dto.CreateContactDTO, entity.Contact]
	userRepository       repository.IRepository[entity.User]
	credentialRepository repository.IRepository[entity.Credential]
	contactRepository    repository.IRepository[entity.Contact]
}

func NewUserFeatureBuilder() *userFeatureBuilder {
	return &userFeatureBuilder{}
}

func (cuf *userFeatureBuilder) SetUnitOfWork(uow uow.IUnitOfWork) *userFeatureBuilder {
	cuf.uow = uow
	return cuf
}

func (cuf *userFeatureBuilder) SetFactories(
	userFactory factory.IFactory[dto.UserDTO, entity.User],
	credentialFactory factory.IFactory[dto.CreateCredentialDTO, entity.Credential],
	contactFactory factory.IFactory[dto.CreateContactDTO, entity.Contact],
) *userFeatureBuilder {

	cuf.userFactory = userFactory
	cuf.credentialFactory = credentialFactory
	cuf.contactFactory = contactFactory

	return cuf
}

func (cuf *userFeatureBuilder) SetRepositories(
	userRepository repository.IRepository[entity.User],
	credentialRepository repository.IRepository[entity.Credential],
	contactRepository repository.IRepository[entity.Contact],
) *userFeatureBuilder {

	cuf.userRepository = userRepository
	cuf.credentialRepository = credentialRepository
	cuf.contactRepository = contactRepository

	return cuf
}

func (cuf *userFeatureBuilder) Build() *UserFeature {
	return &UserFeature{
		uow:                  cuf.uow,
		userFactory:          cuf.userFactory,
		credentialFactory:    cuf.credentialFactory,
		contactFactory:       cuf.contactFactory,
		userRepository:       cuf.userRepository,
		credentialRepository: cuf.credentialRepository,
		contactRepository:    cuf.contactRepository,
	}
}

func parseCreateUserRequestDTO(DTO dto.CreateUserRequestDTO) (dto.UserDTO, dto.CreateCredentialDTO, dto.CreateContactDTO) {
	user := dto.UserDTO{
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
