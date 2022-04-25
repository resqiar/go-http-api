package services

import (
	"http-api/dtos"
	"http-api/entities"
	"http-api/repositories"
)

type IQuestionService interface {
	FindAll() ([]entities.Question, error)
	FindById(id int64) (entities.Question, error)
	Create(questionInput dtos.QuestionInput, authorId int64) error
	UpdateQuestion(updateInput dtos.UpdateQuestionInput) error
}

type questionService struct {
	repository repositories.IQuestionRepository
}

func QuestionService(rep repositories.IQuestionRepository) *questionService {
	return &questionService{rep}
}

func (s *questionService) FindAll() ([]entities.Question, error) {
	// Call users service to retrieve all questions
	result, err := s.repository.FindAll()
	if err != nil {
		return nil, err
	}
	return result, err
}

func (s *questionService) Create(questionInput dtos.QuestionInput, authorId int64) error {
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

func (s *questionService) FindById(id int64) (entities.Question, error) {
	// Call repo to retrieve specific question
	result, err := s.repository.FindById(id)
	return result, err
}

func (s *questionService) UpdateQuestion(updateInput dtos.UpdateQuestionInput) error {
	err := s.repository.Update(int64(updateInput.ID), updateInput)
	return err
}
