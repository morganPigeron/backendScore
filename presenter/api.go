package presenter

import (
	"backend-score/model"
	"backend-score/repository"
	"fmt"
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

	router.POST("/scores", func(c *gin.Context) {
		var score model.Score
		if err := c.ShouldBindJSON(&score); err != nil {
			fmt.Println(score)
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}
		api.Repository.InsertMany([]interface{}{score})
		c.JSON(http.StatusOK, score)
	})

	router.Run(":8080")
}
