package repository

import (
	"context"
	"log"

	"github.com/leenawatH/pic-pick-factory/entity"

	"cloud.google.com/go/firestore"
)

type PeronalRepository interface {
	//Save(item *entity.Item) (*entity.Item, error)
	FindAllPersonalTitle() ([]entity.Item, error)
	AddPersonalTitle(s string) error
}

type personalRepository struct {
	client *firestore.Client
}

func NewPeronalRepository(client *firestore.Client) PeronalRepository {
	return &personalRepository{
		client: client,
	}
}

func (p *personalRepository) FindAllPersonalTitle() ([]entity.Item, error) {
	ctx := context.Background()

	docs, _ := p.client.Collection("personal").Documents(ctx).GetAll()

	var results []entity.Item
	for _, doc := range docs {
		item := entity.Item{
			Title: doc.Data()["title"].(string),
		}
		results = append(results, item)
	}
	return results, nil
}

func (p *personalRepository) AddPersonalTitle(s string) error {

	ctx := context.Background()
	ref := p.client.Collection("personal").NewDoc()

	_, err := ref.Set(ctx, map[string]interface{}{
		"title": s,
	})
	if err != nil {
		// Handle any errors in an appropriate way, such as returning them.
		log.Printf("An error has occurred: %s", err)
	}

	return err
}
