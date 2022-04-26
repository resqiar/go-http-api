package dtos

type UpdateAnswerInput struct {
	ID         int    `json:"id" binding:"required,number"`
	AnswerText string `json:"answer"`
}
