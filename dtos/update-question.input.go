package dtos

type UpdateQuestionInput struct {
	ID    int    `json:"id" binding:"required,number"`
	Title string `json:"title"`
	Desc  string `json:"desc"`
}
