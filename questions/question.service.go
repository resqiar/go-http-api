package questions

type IService interface {
	FindAll() ([]Question, error)
	Create(questionInput QuestionInput, authorId int64) error
}

type service struct {
	repository IQuestionRepository
}

func QuestionService(rep IQuestionRepository) *service {
	return &service{rep}
}

func (s *service) FindAll() ([]Question, error) {
	// Call users service to retrieve all questions
	result, err := s.repository.FindAll()
	if err != nil {
		return nil, err
	}
	return result, err
}

func (s *service) Create(questionInput QuestionInput, authorId int64) error {
	// Bind question input (DTO) into Question obj
	t := Question{
		Title:    questionInput.Title,
		Desc:     questionInput.Desc,
		AuthorID: authorId,
	}

	// Call question repository to create a new user and save them to database.
	err := s.repository.Create(t)
	return err
}
