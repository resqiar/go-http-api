package dtos

type AnswerInput struct {
	AnswerText string `json:"answer" binding:"required, string"`
	QuestionID int64  `json:"question_id" binding:"required, number"`
}
