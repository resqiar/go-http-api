package answers

type IAnswerService interface {
	FindAll() ([]Answer, error)
	Create(answerInput AnswerInput, authorId int64) error
}

type service struct {
	repository IAnswerRepository
}

func AnswerService(rep IAnswerRepository) *service {
	return &service{rep}
}

func (s *service) FindAll() ([]Answer, error) {
	// Call answer service to retrieve all questions
	result, err := s.repository.FindAll()
	if err != nil {
		return nil, err
	}
	return result, err
}

func (s *service) Create(answerInput AnswerInput, authorId int64) error {
	// Bind answer input (DTO) into Answer obj
	t := Answer{
		AnswerText: answerInput.AnswerText,
		AuthorID:   authorId,
		QuestionID: answerInput.QuestionID,
	}

	// Call answer repository to create a new answer and save them to database.
	err := s.repository.Create(t)
	return err
}
