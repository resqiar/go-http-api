package questions

import "gorm.io/gorm"

type IQuestionRepository interface {
	FindAll() ([]Question, error)
	Create(question Question) error
}

type repository struct {
	db *gorm.DB
}

func QuestionRepository(db *gorm.DB) *repository {
	return &repository{db}
}

func (rep *repository) FindAll() ([]Question, error) {
	var result []Question
	err := rep.db.Find(&result).Error
	return result, err
}

func (rep *repository) Create(question Question) error {
	err := rep.db.Create(&question).Save(&question).Error
	return err
}
