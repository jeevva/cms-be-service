package handlers

import (
	"jeevva/cms-be-service/internal/models"
	"jeevva/cms-be-service/internal/repositories"
	"net/http"

	"github.com/gin-gonic/gin"
)

type HandlerUser struct {
	*repositories.RepoUser
}

func NewUser(r *repositories.RepoUser) *HandlerUser{
	return &HandlerUser{r}
}

func (h *HandlerUser) PostData(ctx *gin.Context){
	var user models.User
	if err := ctx.ShouldBind(&user); err != nil{
		ctx.AbortWithError(http.StatusBadGateway, err)
		return
	}

	response, err := h.CreateUser(&user)
	if err != nil{
		ctx.AbortWithError(http.StatusBadGateway, err)
		return
	}

	ctx.JSON(200, response)
}