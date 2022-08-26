package service

import (
	"fmt"
	"os"

	"lightsaid.com/weblogs/internal/models"
	"lightsaid.com/weblogs/pkg/utils"
)

func CreateDefaultUsername() string {
	return fmt.Sprintf("%s_%s", os.Getenv("USERNAME_PREFIX"), utils.RandomString(6))
}

func CreateDefaultAvatar() string {
	return fmt.Sprintf("%s/static/avatar/%d.png", os.Getenv("ASSETS_PREFIX"), utils.RandomInt(1, 5))
}

func (s *Service) Register(arg CreateUserRequest) (models.User, error) {
	user, err := s.Repository.InsertUser(arg.Email, arg.Username, arg.Password, arg.Avatar)
	return user, err
}

func (s *Service) Login() {

}
