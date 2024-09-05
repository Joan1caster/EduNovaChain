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
	TokenID          string        `gorm:"unique;not null"`
	ContractAddress  string        `gorm:"not null"`
	OwnerID          uint          `gorm:"index;not null"`
	CreatorID        uint          `gorm:"index;not null"`
	Grades           []Grade       `gorm:"many2many:nft_grades;"`
	Subjects         []Subject     `gorm:"many2many:nft_subjects;"`
	Topics           []Topic       `gorm:"many2many:nft_topics;"`
	Categories       []NFTCategory `gorm:"type:varchar(20)[]"`
	MetadataURI      string        `gorm:"not null"`
	SummaryFeature   []byte        `gorm:"type:blob;not null"`
	ContentFeature   []byte        `gorm:"type:blob;not null"`
	Owner            User          `gorm:"foreignKey:OwnerID"`
	Creator          User          `gorm:"foreignKey:CreatorID"`
	LikeCount        uint          `gorm:"default:0"`
	ViewCount        uint          `gorm:"default:0"`
	TransactionCount uint          `gorm:"default:0"`
}

type Grade struct {
	gorm.Model
	Name        string `gorm:"uniqueIndex;not null"`
	NFTs        []NFT
	TotalVisits uint `gorm:"default:0"`
}

type Subject struct {
	gorm.Model
	Name        string `gorm:"uniqueIndex;not null"`
	NFTs        []NFT
	TotalVisits uint `gorm:"default:0"`
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
