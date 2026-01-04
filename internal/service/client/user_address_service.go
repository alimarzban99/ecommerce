package client

import (
	dtoClient "github.com/alimarzban99/ecommerce/internal/dto/client"
	"github.com/alimarzban99/ecommerce/internal/repository"
	"github.com/alimarzban99/ecommerce/internal/resources/client"
)

type UserAddressService struct {
	repo repository.UserAddressRepositoryInterface
}

func NewUserAddressService(repo repository.UserAddressRepositoryInterface) *UserAddressService {
	return &UserAddressService{repo: repo}
}

func (s *UserAddressService) Index(dto dtoClient.ListUserAddressDTO, userId int) (*repository.PaginatedResponse[client.UserAddressListResource], error) {
	return s.repo.List(dto, userId)
}

func (s *UserAddressService) Show(id, userId int) (*client.UserAddressResource, error) {
	return s.repo.FindOne(id, userId)
}

func (s *UserAddressService) Store(dto *dtoClient.StoreUserAddressDTO, userId int) (*client.UserAddressResource, error) {
	return s.repo.Create(dto, userId)
}

func (s *UserAddressService) Update(id, userId int, dto *dtoClient.UpdateUserAddressDTO) error {
	return s.repo.Update(id, userId, dto)
}

func (s *UserAddressService) Destroy(id, userId int) error {
	return s.repo.Destroy(id, userId)
}
