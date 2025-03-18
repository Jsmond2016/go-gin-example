package service

import "github.com/EDDYCJY/go-gin-example/internal/repository"

type Services struct {
	User UserService
}

func NewServices(repos *repository.Repositories) *Services {
	return &Services{
		User: NewUserService(repos.User),
	}
}