package models

import (
	"time"

	"gorm.io/gorm"
)

type NFTCategory string

const (
	CategoryHot        NFTCategory = "hot"
	CategoryBestseller NFTCategory = "bestseller"
	CategoryNewest     NFTCategory = "newest"
)

type NFT struct {
	gorm.Model
	TokenID          string        `gorm:"unique;not null"`         // 合约里的ID
	ContractAddress  string        `gorm:"not null"`                // 合约地址
	OwnerID          uint          `gorm:"index;not null"`          // 所有者ID
	CreatorID        uint          `gorm:"index;not null"`          // 作者ID
	Grades           []Grade       `gorm:"many2many:nft_grades;"`   // 所属年级
	Subjects         []Subject     `gorm:"many2many:nft_subjects;"` // 所属学科
	Topics           []Topic       `gorm:"many2many:nft_topics;"`   // 所属主题
	Categories       []NFTCategory `gorm:"type:varchar(20)[]"`      // 所属分类
	MetadataURI      string        `gorm:"not null"`                // IPFS的存储数据，名称-摘要-标题等
	SummaryFeature   []byte        `gorm:"type:blob;not null"`      // 摘要特征值
	ContentFeature   []byte        `gorm:"type:blob;not null"`      // 正文特征值
	Owner            User          `gorm:"foreignKey:OwnerID"`
	Creator          User          `gorm:"foreignKey:CreatorID"`
	LikeCount        uint          `gorm:"default:0"` // 点赞次数
	ViewCount        uint          `gorm:"default:0"` // 浏览次数
	TransactionCount uint          `gorm:"default:0"` // 交易次数
}

type Grade struct {
	gorm.Model
	Name        string `gorm:"uniqueIndex;not null"`
	NFTs        []NFT  `gorm:"many2many:nft_grades;"`
	TotalVisits uint   `gorm:"default:0"`
}

type Subject struct {
	gorm.Model
	Name        string `gorm:"uniqueIndex;not null"`
	NFTs        []NFT  `gorm:"many2many:nft_subjects;"`
	TotalVisits uint   `gorm:"default:0"`
}

type Topic struct {
	gorm.Model
	Name        string `gorm:"uniqueIndex;not null"`
	NFTs        []NFT  `gorm:"many2many:nft_topics;"`
	Users       []User `gorm:"many2many:user_topics;"`
	TotalVisits uint   `gorm:"default:0"`
}

type UserTopicVisit struct {
	gorm.Model
	UserID        uint `gorm:"index:idx_user_topic,unique"`
	TopicID       uint `gorm:"index:idx_user_topic,unique"`
	VisitCount    uint `gorm:"default:1"`
	LastVisitTime time.Time
}

type NFTWithSimilarity struct {
	*NFT
	Similarity float32
}
