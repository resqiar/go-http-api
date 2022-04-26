package services

import (
	"http-api/dtos"
	"http-api/entities"
	"http-api/repositories"
)

type IAnswerService interface {
	FindAll() ([]entities.Answer, error)
	FindById(id int64) (entities.Answer, error)
	Create(answerInput dtos.AnswerInput, authorId int64) error
	UpdateAnswer(updateInput dtos.UpdateAnswerInput) error
	SoftDeleteAnswer(deleteInput dtos.DeleteAnswerInput) error
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

func (s *answerService) FindById(id int64) (entities.Answer, error) {
	// Call repo to retrieve specific answer
	result, err := s.repository.FindById(id)
	return result, err
}

func (s *answerService) UpdateAnswer(updateInput dtos.UpdateAnswerInput) error {
	err := s.repository.Update(int64(updateInput.ID), updateInput)
	return err
}

func (s *answerService) SoftDeleteAnswer(deleteInput dtos.DeleteAnswerInput) error {
	// Call repo to SOFT DELETE the given id
	err := s.repository.Delete(int64(deleteInput.ID))
	return err
}
