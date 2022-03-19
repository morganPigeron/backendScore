package controllers

import (
	"backend-score/core/score"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
)

type Api struct {
	service score.GameScoreService
}

func NewApi() *Api {
	return new(Api)
}

func (api *Api) SetService(service score.GameScoreService) {
	api.service = service
}

func (api *Api) Start() {

	router := gin.Default()

	router.GET("/scores", func(c *gin.Context) {
		result, err := api.service.GetAll()
		if err != nil {
			log.Fatal("getAll failed")
		}
		c.JSON(http.StatusOK, result)
	})

	router.POST("/scores", func(c *gin.Context) {
		var score score.Score
		if err := c.ShouldBindJSON(&score); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}
		err := api.service.Save(score)
		if err != nil {
			log.Fatal(err)
		}
		c.JSON(http.StatusOK, score)
	})

	router.Run(":8080")
}
