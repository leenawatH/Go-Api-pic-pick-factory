package repository

import (
	"context"
	"log"

	firebase "firebase.google.com/go"
	"firebase.google.com/go/auth"
)

type LoginRepository interface {
	Login(tokenid string) (*auth.UserRecord, error)
}

type loginRepository struct {
	app *firebase.App
}

func NewLoginRepository(app *firebase.App) *loginRepository {
	return &loginRepository{
		app: app,
	}
}

func (p *loginRepository) Login(idtoken string) (*auth.UserRecord, error) {

	ctx := context.Background()
	client, err := p.app.Auth(ctx)
	if err != nil {
		log.Printf("Error get Auth client: %v\n", err)
		return nil, err
	}

	token, err := client.VerifyIDToken(ctx, idtoken)

	if err != nil {
		log.Printf("Error verifying ID token: %v\n", err)
		return nil, err
	}

	user, err := client.GetUser(ctx, token.UID)

	if err != nil {
		log.Printf("Error getting user record: %v\n", err)
		return nil, err
	}

	return user, err
}
