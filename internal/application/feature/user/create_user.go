package userfeature

import (
	"stock-controll/internal/application/dto"
	"stock-controll/internal/domain/entity"
	"stock-controll/internal/domain/factory"
	"stock-controll/internal/domain/repository"
	uow "stock-controll/internal/infrastructure/unit_work"
)
type createUserFeatureBuilder struct {
	uow uow.IUnitOfWork
	userFactory       factory.IUserFactory
	credentialFactory factory.ICredentialFactory
	contactFactory    factory.IContactFactory
	userRepository    repository.IRepository[entity.User]
	credentialRepository repository.IRepository[entity.Credential]
	contactRepository repository.IRepository[entity.Contact]
}

func CreateUserFeatureBuilder() *createUserFeatureBuilder {
	return &createUserFeatureBuilder{}
} 

func (cub *createUserFeatureBuilder) SetUnitOfWork(uow uow.IUnitOfWork) *createUserFeatureBuilder {
	cub.uow = uow
	return cub
}

func (cub *createUserFeatureBuilder) SetFactories(
	userFactory factory.IUserFactory, 
	credentialFactory factory.ICredentialFactory,
	contactFactory factory.IContactFactory,
	) *createUserFeatureBuilder {

	cub.userFactory = userFactory
	cub.credentialFactory = credentialFactory
	cub.contactFactory = contactFactory

	return cub
}

func (cub *createUserFeatureBuilder) SetRepositories(
	userRepository repository.IRepository[entity.User],
	credentialRepository repository.IRepository[entity.Credential],
	contactRepository repository.IRepository[entity.Contact],
	) *createUserFeatureBuilder {

	cub.userRepository = userRepository
	cub.credentialRepository = credentialRepository
	cub.contactRepository = contactRepository

	return cub
}

func (cub *createUserFeatureBuilder) Build() *CreateUserFeature {
	return &CreateUserFeature{
		uow: cub.uow,
		userFactory: cub.userFactory,
		credentialFactory: cub.credentialFactory,
		contactFactory: cub.contactFactory,
		userRepository: cub.userRepository,
		credentialRepository: cub.credentialRepository,
		contactRepository: cub.contactRepository,
	}
}

type CreateUserFeature struct {
	uow               uow.IUnitOfWork
	userFactory       factory.IUserFactory
	credentialFactory factory.ICredentialFactory
	contactFactory    factory.IContactFactory
	userRepository    repository.IRepository[entity.User]
	credentialRepository repository.IRepository[entity.Credential]
	contactRepository repository.IRepository[entity.Contact]
}


func (uf *CreateUserFeature) CreateUser(userData dto.CreateUserRequestDTO) (*dto.CreateUserResponseDTO, error) {
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
