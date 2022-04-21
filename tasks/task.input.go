package tasks

type TaskInput struct {
	Task   string `json:"task" binding:"required"`
	Desc   string `json:"desc"`
	IsDone *bool  `json:"isDone" binding:"required"`
}
