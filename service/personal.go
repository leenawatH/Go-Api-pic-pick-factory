package service

import (
	"log"

	"github.com/leenawatH/pic-pick-factory/entity"
	"github.com/leenawatH/pic-pick-factory/repository"
)

type personalService struct {
	repository repository.PeronalRepository
}

type PersonalService interface {
	GetAllPersonalTitle() ([]entity.Item, error)

	AddPersonalTitle(s string) error
}

func NewPeronalService(personalRepository repository.PeronalRepository) *personalService {
	return &personalService{
		repository: personalRepository,
	}
}

func (p *personalService) GetAllPersonalTitle() (items []entity.Item, err error) {
	items, err = p.repository.FindAllPersonalTitle()
	if err != nil {
		log.Fatalf("err to service : %v", err)
		return
	}

	return items, nil
}

func (p *personalService) AddPersonalTitle(s string) (err error) {
	_ = p.repository.AddPersonalTitle(s)

	return
}
