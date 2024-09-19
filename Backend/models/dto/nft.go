package dto

type IDName struct {
	ID   uint   `json:"id"`
	Name string `json:"name"`
}

type NFTQuery struct {
	Keyword  *string `json:"keyword"`
	GradeIDs *[]uint `json:"gradeIds"`
	Subjects *[]uint `json:"subjectIds"`
	TopicIds *[]uint `json:"topicIds"`
	Page     *uint   `json:"page"`
	PageSize *uint   `json:"pagesize"`
}

type CreateNFT struct {
	TokenID         string       `json:"tokenId" binding:"required"`
	ContractAddress string       `json:"contractAddress" binding:"required"`
	MetadataURI     string       `json:"metadataURI" binding:"required"`
	SummaryFeature  [512]float32 `json:"summaryFeature" binding:"required"`
	ContentFeature  [512]float32 `json:"contentFeature" binding:"required"`
	Grade           string       `json:"grade" binging:"required"`
	Subject         string       `json:"subject" binging:"required"`
	Topic           string       `json:"topic" binging:"required"`
	Price           float64      `json:"price" binging:"required"`
}

type OrderMessage struct {
	NFTId uint    `json:"nftId" binding:"required"`
	Price float64 `json:"price" binding:"required"`
}

type SubjectsAndGrades struct {
	Subjectids []*uint `json:"subjectIds"`
	Gradeids   []*uint `json:"gradeIds"`
}
