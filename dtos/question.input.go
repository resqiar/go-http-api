package dtos

type QuestionInput struct {
	Title string `json:"title" binding:"required"`
	Desc  string `json:"desc" binding:"required"`
}
