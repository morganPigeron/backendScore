package score_test

import (
	"backend-score/core/score"
	"backend-score/presenters"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestPresenter(t *testing.T) {
	presenter := presenters.NewFakePresenter()
	service := score.NewGameScoreService(presenter)
	testScore := score.Score{
		Player: "test",
		Date:   "test",
		Points: 100,
	}
	service.Save(testScore)
	s, _ := service.GetAll()
	assert.Equal(t, s[0], testScore, "they should be equal")
}
