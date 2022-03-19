package score

type GameScoreRepository interface {
	Save([]Score) error
	GetAll() ([]Score, error)
}
