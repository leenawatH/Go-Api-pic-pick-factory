package handler

import (
	"database/sql"
	"errors"
	"net/http"

	"github.com/leenawatH/pic-pick-factory/entity"
	"github.com/leenawatH/pic-pick-factory/service"

	"github.com/gin-gonic/gin"
)

type personalHandler struct {
	service service.PersonalService
}

func NewPersonalHandler(personalService service.PersonalService) *personalHandler {
	return &personalHandler{
		service: personalService,
	}
}

func (p personalHandler) GetAllPersonalTitle(g *gin.Context) {
	popData, err := p.service.GetAllPersonalTitle()
	if err == nil {
		g.JSON(http.StatusOK, popData)

		return
	} else if errors.Is(err, sql.ErrNoRows) {
		g.JSON(http.StatusNotFound, gin.H{
			"message": "Not matching data",
		})

		return
	}
	g.JSON(http.StatusInternalServerError, gin.H{
		"message": "Something went wrong.",
	})

}

func (p personalHandler) AddPersonalTitle(g *gin.Context) {
	var data entity.Item

	err := g.BindJSON(&data)

	if err != nil {
		g.JSON(http.StatusBadRequest, gin.H{
			"message": "Invalid request body.",
		})
		return
	} else {
		_ = p.service.AddPersonalTitle(data.Title)
		g.JSON(http.StatusOK, gin.H{
			"message": data.Title,
		})
	}

	return
}
