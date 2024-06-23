package handler

import (
	"net/http"

	"github.com/leenawatH/pic-pick-factory/entity"
	"github.com/leenawatH/pic-pick-factory/service"

	"github.com/gin-gonic/gin"
)

type loginHandler struct {
	service service.LoginService
}

func NewLoginHandler(loginService service.LoginService) *loginHandler {
	return &loginHandler{
		service: loginService,
	}
}

func (p loginHandler) Login(g *gin.Context) {
	var data entity.Account

	err := g.BindJSON(&data)

	if err != nil {
		g.JSON(http.StatusBadRequest, gin.H{
			"message": "Invalid request body.",
		})
		return
	} else {
		_ = p.service.Login(data.TokenId)
		g.JSON(http.StatusOK, gin.H{
			"message": "success",
		})
	}

	return
}
