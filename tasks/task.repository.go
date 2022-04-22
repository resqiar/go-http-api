package tasks

import "gorm.io/gorm"

type ITaskRepository interface {
	FindAll() ([]Task, error)
	Create(task Task) error
}

type repository struct {
	db *gorm.DB
}

func TaskRepository(db *gorm.DB) *repository {
	return &repository{db}
}

func (rep *repository) FindAll() ([]Task, error) {
	var result []Task
	err := rep.db.Find(&result).Error
	return result, err
}

func (rep *repository) Create(task Task) error {
	err := rep.db.Create(&task).Save(&task).Error
	return err
}
