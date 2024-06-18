package main

import (
	"log"
	"net/http"
	"os"

	"github.com/leenawatH/pic-pick-factory/handler"
	"github.com/leenawatH/pic-pick-factory/repository"
	"github.com/leenawatH/pic-pick-factory/service"

	"github.com/joho/godotenv"

	"github.com/gin-gonic/gin"

	"context"
	//"fmt"

	firebase "firebase.google.com/go"

	"cloud.google.com/go/firestore"
	"google.golang.org/api/option"
)

const (
	projectId string = "pic-pick-factory"
)

func main() {

	err := godotenv.Load()
	if err != nil {
		log.Fatalf("Error loading .env file: %v", err)
	}

	sdkPath := os.Getenv("FIREBASE_ADMIN_SDK_JSON")
	if sdkPath == "" {
		log.Fatal("FIREBASE_ADMIN_SDK_JSON environment variable is not set")
	}

	ctx := context.Background()
	opt := option.WithCredentialsFile(sdkPath)
	_, err = firebase.NewApp(ctx, nil, opt)
	if err != nil {
		log.Fatalf("err to start firebase : %v", err)
		return
	}

	client, err := firestore.NewClient(ctx, projectId, opt)
	if err != nil {
		log.Fatalf("err to create Firebase client in main : %v", err)
	}

	defer client.Close()

	personalRepository := repository.NewPeronalRepository(client)
	personalService := service.NewPeronalService(personalRepository)
	personalHandler := handler.NewPersonalHandler(personalService)

	commissionedRepository := repository.NewCommissionedRepository(client)
	commissionedService := service.NewCommissionedService(commissionedRepository)
	commissionedHandler := handler.NewCommissionedHandler(commissionedService)

	r := gin.Default()
	//example get /ping
	r.GET("/ping", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "pong",
		})
	})

	//Get All personal
	r.GET("/personal", personalHandler.GetAllPersonalTitle)
	r.GET("/commissioned", commissionedHandler.GetAllCommissionedTitle)

	r.POST("/addpertitle", personalHandler.AddPersonalTitle)
	r.POST("/addcomtitle", commissionedHandler.AddCommissionedTitle)
	r.Run() // listen and serve on 0.0.0.0:8080 (for windows "localhost:8080")

}
