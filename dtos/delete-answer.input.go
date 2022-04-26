package dtos

type DeleteAnswerInput struct {
	ID int `json:"id" binding:"required,number"`
}
