package tasks

type IService interface {
	FindAll() ([]Task, error)
	Create(taskInput TaskInput) (int, error)
}

type service struct {
	repository ITaskRepository
}

func TaskService(rep ITaskRepository) *service {
	return &service{rep}
}

func (s *service) FindAll() ([]Task, error) {
	result, err := s.repository.FindAll()
	if err != nil {
		return nil, err
	}
	return result, err
}

func (s *service) Create(taskInput TaskInput) (int, error) {
	t := Task{
		Task:   taskInput.Task,
		Desc:   taskInput.Desc,
		IsDone: taskInput.IsDone,
	}
	result, err := s.repository.Create(t)
	return result, err
}
