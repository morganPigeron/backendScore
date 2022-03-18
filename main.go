package main

import (
	"backend-score/presenter"
	"backend-score/repository"
)

func main() {

	db := repository.NewMangoDB()
	defer db.Disconnect()

	api := presenter.NewApi()
	api.SetRepository(db)
	api.Start()

}
