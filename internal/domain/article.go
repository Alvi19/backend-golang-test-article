package domain

import "time"

type ArticleStatus string

const (
	StatusPublish ArticleStatus = "Publish"
	StatusDraft   ArticleStatus = "Draft"
	StatusTrash   ArticleStatus = "Trash"
)

type Article struct {
	ID          uint          `gorm:"primaryKey;autoIncrement" json:"id"`
	Title       string        `gorm:"type:varchar(200);not null" json:"title"`
	Content     string        `gorm:"type:text;not null" json:"content"`
	Category    string        `gorm:"type:varchar(100);not null" json:"category"`
	CreatedDate time.Time     `gorm:"type:timestamp;autoCreateTime" json:"created_date"`
	UpdatedDate time.Time     `gorm:"type:timestamp;autoUpdateTime" json:"updated_date"`
	Status      ArticleStatus `gorm:"type:varchar(100);not null;default:'Draft'" json:"status"`
}

func (Article) TableName() string {
	return "articles"
}
