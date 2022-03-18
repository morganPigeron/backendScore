package main

import (
	"backend-score/model"
	"backend-score/presenter"
	"backend-score/repository"
	"math/rand"
	"time"
)

func main() {

	// generate data
	p1 := model.Score{
		Player: "player1",
		Date:   time.Now().String(),
		Points: rand.Intn(100),
	}
	p2 := model.Score{
		Player: "player2",
		Date:   time.Now().String(),
		Points: rand.Intn(100),
	}
	p3 := model.Score{
		Player: "player3",
		Date:   time.Now().String(),
		Points: rand.Intn(100),
	}

	players := []interface{}{p1, p2, p3}

	db := repository.NewMangoDB()
	defer db.Disconnect()
	db.InsertMany(players)

	api := presenter.NewApi()
	api.SetRepository(db)
	api.Start()

}
