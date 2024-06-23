package service

import (
	"log"

	"github.com/leenawatH/pic-pick-factory/entity"
	"github.com/leenawatH/pic-pick-factory/repository"
)

type CommissionedService interface {
	GetAllCommissionedTitle() ([]entity.Item, error)
	AddCommissionedTitle(s string) (err error)
}

type commissionedService struct {
	repository repository.CommissionedRepository
}

func NewCommissionedService(commissionedRepository repository.CommissionedRepository) *commissionedService {
	return &commissionedService{
		repository: commissionedRepository,
	}
}

func (p *commissionedService) GetAllCommissionedTitle() (items []entity.Item, err error) {
	items, err = p.repository.FindAllCommissionedTitle()
	if err != nil {
		log.Fatalf("err to service : %v", err)
		return
	}

	return items, nil
}

func (p *commissionedService) AddCommissionedTitle(s string) (err error) {
	_ = p.repository.AddCommissionedTitle(s)

	return
}
