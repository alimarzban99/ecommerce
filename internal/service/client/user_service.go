package client

import (
	dtoAdmin "github.com/alimarzban99/ecommerce/internal/dto/client"
	"github.com/alimarzban99/ecommerce/internal/repository"
	"github.com/alimarzban99/ecommerce/internal/resources/client"
)

type UserService struct {
	repo *repository.UserRepository
}

func NewUserService() *UserService {
	return &UserService{repo: repository.NewUserRepository()}
}

func (s UserService) Profile(id int) (*client.UserResource, error) {
	return s.repo.FindOne(id)
}

func (s UserService) Update(id int, dto *dtoAdmin.UpdateUserDTO) error {
	return s.repo.Update(id, dto)
}

func (s UserService) UpdateProfile(id int, dto *dtoAdmin.UpdateProfileDTO) (*client.UserResource, error) {
	// Convert UpdateProfileDTO to map for partial update
	updateMap := make(map[string]interface{})
	if dto.FirstName != nil {
		updateMap["first_name"] = *dto.FirstName
	}
	if dto.LastName != nil {
		updateMap["last_name"] = *dto.LastName
	}
	if dto.Email != nil {
		updateMap["email"] = *dto.Email
	}

	// Perform the update
	err := s.repo.UpdateProfile(id, updateMap)
	if err != nil {
		return nil, err
	}

	// Return updated user
	return s.repo.FindOne(id)
}
