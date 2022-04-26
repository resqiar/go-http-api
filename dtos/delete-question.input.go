package dtos

type DeleteQuestionInput struct {
	ID int `json:"id" binding:"required,number"`
}
