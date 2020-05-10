package oauth

import (
	"grpc-go/application/users"
	"grpc-go/domain/oauth"
	"grpc-go/framework/utils/bcrypt"
)

type OAuthService struct {
	OAuthRepository users.UserRepository
}

func (service *OAuthService) DoLogin(email string, password string) (*oauth.OAuth, error) {
	var auth oauth.OAuth
	user, err := service.OAuthRepository.Fetch(email)
	if err != nil {
		return nil, err
	}

	isValid := bcrypt.ComparedPassword(user.Password, password)
	if isValid == false {
		return nil, err
	}

	auth.Token = "token"

	return &auth, nil
}
