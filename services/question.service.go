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

type questionservice struct {
	repository repositories.IQuestionRepository
}

func QuestionService(rep repositories.IQuestionRepository) *questionservice {
	return &questionservice{rep}
}

func (s *questionservice) FindAll() ([]entities.Question, error) {
	// Call users service to retrieve all questions
	result, err := s.repository.FindAll()
	if err != nil {
		return nil, err
	}
	return result, err
}

func (s *questionservice) Create(questionInput dtos.QuestionInput, authorId int64) error {
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
