package service

import (
	"github.com/leenawatH/pic-pick-factory/repository"
)

type LoginService interface {
	Login(tokenid string) error
}

type loginService struct {
	repository repository.LoginRepository
}

func NewLoginService(loginRepository repository.LoginRepository) *loginService {
	return &loginService{
		repository: loginRepository,
	}
}

func (p *loginService) Login(tokenid string) error {
	_, err := p.repository.Login(tokenid)

	return err
}
