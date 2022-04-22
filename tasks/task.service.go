package tasks

type IService interface {
	FindAll() ([]Task, error)
	Create(taskInput TaskInput, authorId int64) error
}

type service struct {
	repository ITaskRepository
}

func TaskService(rep ITaskRepository) *service {
	return &service{rep}
}

func (s *service) FindAll() ([]Task, error) {
	// Call users service to retrieve all tasks
	result, err := s.repository.FindAll()
	if err != nil {
		return nil, err
	}
	return result, err
}

func (s *service) Create(taskInput TaskInput, authorId int64) error {
	// Bind task input (DTO) into Task obj
	t := Task{
		Task:     taskInput.Task,
		Desc:     taskInput.Desc,
		IsDone:   taskInput.IsDone,
		AuthorID: authorId,
	}

	// Call task repository to create a new user and save them to database.
	err := s.repository.Create(t)
	return err
}
