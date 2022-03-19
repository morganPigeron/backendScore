package score

type GameScoreService interface {
	Save(score Score) error
	GetAll() ([]Score, error)
}
