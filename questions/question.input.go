package questions

type QuestionInput struct {
	Title string `json:"question" binding:"required"`
	Desc  string `json:"desc"`
}
