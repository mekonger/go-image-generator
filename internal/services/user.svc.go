package services

import "github.com/mekonger/go-image-generator/internal/repo"

type UserService interface {
	Hello(name string) string
	PostHello(name string, email string) string
}

type userService struct {
	userRepo repo.UserRepo
}

func NewUserService(userRepo repo.UserRepo) UserService {
	return &userService{
		userRepo: userRepo,
	}
}

func (u *userService) Hello(name string) string {
	return u.userRepo.Hello(name)
}

func (u *userService) PostHello(name string, email string) string {
	return u.userRepo.PostHello(name, email)
}
