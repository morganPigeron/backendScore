package presenter

import (
	"backend-score/repository"
	"net/http"

	"github.com/gin-gonic/gin"
)

type Api struct {
	Repository *repository.MangoDB
}

func NewApi() *Api {
	return new(Api)
}

func (api *Api) SetRepository(r *repository.MangoDB) {
	api.Repository = r
}

func (api *Api) Start() {
	router := gin.Default()
	router.GET("/scores", func(c *gin.Context) {
		result := api.Repository.GetAll()
		c.JSON(http.StatusOK, result)
	})
	router.Run(":8080")
}
