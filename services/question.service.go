package services

import (
	"http-api/dtos"
	"http-api/entities"
	"http-api/repositories"
)

type IQuestionService interface {
	FindAll() ([]entities.Question, error)
	Create(questionInput dtos.QuestionInput, authorId int64) error
}

type service struct {
	repository repositories.IQuestionRepository
}

func QuestionService(rep repositories.IQuestionRepository) *service {
	return &service{rep}
}

func (s *service) FindAll() ([]entities.Question, error) {
	// Call users service to retrieve all questions
	result, err := s.repository.FindAll()
	if err != nil {
		return nil, err
	}
	return result, err
}

func (s *service) Create(questionInput dtos.QuestionInput, authorId int64) error {
	// Bind question input (DTO) into Question obj
	t := entities.Question{
		Title:    questionInput.Title,
		Desc:     questionInput.Desc,
		AuthorID: authorId,
	}

	// Call question repository to create a new user and save them to database.
	err := s.repository.Create(t)
	return err
}
