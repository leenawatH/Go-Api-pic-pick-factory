package handler

import (
	"database/sql"
	"errors"
	"net/http"

	"github.com/leenawatH/pic-pick-factory/entity"
	"github.com/leenawatH/pic-pick-factory/service"

	"github.com/gin-gonic/gin"
)

type commissionedHandler struct {
	service service.CommissionedService
}

func NewCommissionedHandler(commissionedService service.CommissionedService) *commissionedHandler {
	return &commissionedHandler{
		service: commissionedService,
	}
}

func (p commissionedHandler) GetAllCommissionedTitle(g *gin.Context) {
	popData, err := p.service.GetAllCommissionedTitle()
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

func (p commissionedHandler) AddCommissionedTitle(g *gin.Context) {
	var data entity.Item

	err := g.BindJSON(&data)

	if err != nil {
		g.JSON(http.StatusBadRequest, gin.H{
			"message": "Invalid request body.",
		})
		return
	} else {
		_ = p.service.AddCommissionedTitle(data.Title)
		g.JSON(http.StatusOK, gin.H{
			"message": data.Title,
		})
	}

	return
}
