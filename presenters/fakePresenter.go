package presenters

import (
	"backend-score/core/score"
)

type fakePresenter struct {
	scores []score.Score
}

func NewFakePresenter() *fakePresenter {
	var s []score.Score
	return &fakePresenter{
		scores: s,
	}
}

func (p *fakePresenter) Save(score []score.Score) error {
	p.scores = append(p.scores, score...)
	return nil
}

func (p *fakePresenter) GetAll() ([]score.Score, error) {
	return p.scores, nil
}
