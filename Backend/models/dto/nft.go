package dto

type IDName struct {
	ID   uint   `json:"id"`
	Name string `json:"name"`
}

type NFTQuery struct {
	Keyword *string `json:"keyword"`
	GradeID *uint `json:"gradeId"`
	Subjects *[]uint `json:"gradeIds"`
	TopicIds *[]uint `json:"topicIds"`
	page uint `json:"page"`
}

