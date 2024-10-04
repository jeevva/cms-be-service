package routers

import (
	"jeevva/cms-be-service/internal/handlers"
	"jeevva/cms-be-service/internal/repositories"

	"github.com/gin-gonic/gin"
	"github.com/jmoiron/sqlx"
)

func user(g *gin.Engine, d *sqlx.DB) {
	route := g.Group("/user")

	repo := repositories.NewUser(d)
	handler := handlers.NewUser(repo)
	
	route.POST("/", handler.PostData)
}