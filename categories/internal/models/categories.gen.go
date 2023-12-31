// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.

package models

import (
	"time"

	"gorm.io/gorm"
)

const TableNameCategory = "categories"

// Category mapped from table <categories>
type Category struct {
	CategoryID  uint64         `gorm:"column:category_id;type:bigint(20) unsigned;primaryKey;autoIncrement:true" json:"category_id"`
	Name        string         `gorm:"column:name;type:varchar(50);not null" json:"name"`
	Slogan      string         `gorm:"column:slogan;type:varchar(150);not null" json:"slogan"`
	Description string         `gorm:"column:description;type:longtext;not null" json:"description"`
	Parent      uint64         `gorm:"column:parent;type:bigint(20) unsigned;not null;index:parent,priority:1" json:"parent"`
	PostsCount  uint64         `gorm:"column:posts_count;type:bigint(20) unsigned;not null" json:"posts_count"`
	CreatedAt   time.Time      `gorm:"column:created_at;type:timestamp;not null;default:current_timestamp()" json:"created_at"`
	DeletedAt   gorm.DeletedAt `gorm:"column:deleted_at;type:timestamp" json:"deleted_at"`
	UpdatedAt   time.Time      `gorm:"column:updated_at;type:timestamp" json:"updated_at"`
}

// TableName Category's table name
func (*Category) TableName() string {
	return TableNameCategory
}
