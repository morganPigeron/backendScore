package main

import (
	"backend-score/controllers"
	"backend-score/core/score"
	"backend-score/presenters"
)

func main() {
	db := presenters.NewMangoRepository()
	service := score.NewGameScoreService(db)

	api := controllers.NewApi()
	api.SetService(service)
	api.Start()
}
