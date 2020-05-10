package users

import (
	domain "grpc-go/domain/users"
)

type UserService struct {
	UserRepository UserRepository
}

func (service *UserService) Save(data *domain.User) (*domain.User, error) {
	user, err := service.UserRepository.Create(data)
	if err != nil {
		return nil, err
	}

	return user, nil
}
