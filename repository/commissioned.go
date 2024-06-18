package repository

import (
	"context"
	"log"

	"github.com/leenawatH/pic-pick-factory/entity"

	"cloud.google.com/go/firestore"
)

type CommissionedRepository interface {
	//Save(item *entity.Item) (*entity.Item, error)
	FindAllCommissionedTitle() ([]entity.Item, error)
	AddCommissionedTitle(s string) error
}

type commissionedRepository struct {
	client *firestore.Client
}

func NewCommissionedRepository(client *firestore.Client) CommissionedRepository {
	return &commissionedRepository{
		client: client,
	}
}

// Find All Title
func (p *commissionedRepository) FindAllCommissionedTitle() ([]entity.Item, error) {
	ctx := context.Background()

	docs, _ := p.client.Collection("commissioned").Documents(ctx).GetAll()

	var results []entity.Item
	for _, doc := range docs {
		item := entity.Item{
			Title: doc.Data()["title"].(string),
		}
		results = append(results, item)
	}
	return results, nil
}

// Add new Title
func (p *commissionedRepository) AddCommissionedTitle(s string) error {

	ctx := context.Background()
	ref := p.client.Collection("commissioned").NewDoc()

	_, err := ref.Set(ctx, map[string]interface{}{
		"title": s,
	})
	if err != nil {
		// Handle any errors in an appropriate way, such as returning them.
		log.Printf("An error has occurred: %s", err)
	}

	return err
}
