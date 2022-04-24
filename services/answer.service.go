package services

import (
	"http-api/dtos"
	"http-api/entities"
	"http-api/repositories"
)

type IAnswerService interface {
	FindAll() ([]entities.Answer, error)
	Create(answerInput dtos.AnswerInput, authorId int64) error
}

type answerService struct {
	repository repositories.IAnswerRepository
}

func AnswerService(rep repositories.IAnswerRepository) *answerService {
	return &answerService{rep}
}

func (s *answerService) FindAll() ([]entities.Answer, error) {
	// Call answer service to retrieve all questions
	result, err := s.repository.FindAll()
	if err != nil {
		return nil, err
	}
	return result, err
}

func (s *answerService) Create(answerInput dtos.AnswerInput, authorId int64) error {
	// Bind answer input (DTO) into Answer obj
	t := entities.Answer{
		AnswerText: answerInput.AnswerText,
		AuthorID:   authorId,
		QuestionID: answerInput.QuestionID,
	}

	// Call answer repository to create a new answer and save them to database.
	err := s.repository.Create(t)
	return err
}
