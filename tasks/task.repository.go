package tasks

import "gorm.io/gorm"

type ITaskRepository interface {
	FindAll() ([]Task, error)
	Create(task Task) (int, error)
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

func (rep *repository) Create(task Task) (int, error) {
	err := rep.db.Create(&task).Error
	return 200, err
}
