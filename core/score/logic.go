package score

// repo <- service -> controller

type gameScoreService struct {
	repository GameScoreRepository
}

func NewGameScoreService(r GameScoreRepository) GameScoreService {
	return &gameScoreService{
		repository: r,
	}
}

func (s *gameScoreService) Save(score Score) error {
	err := s.repository.Save([]Score{score})
	return err
}

func (s *gameScoreService) GetAll() ([]Score, error) {
	scores, err := s.repository.GetAll()
	return scores, err
}
