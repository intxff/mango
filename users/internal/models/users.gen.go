// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.

package models

import (
	"time"

	"gorm.io/gorm"
)

const TableNameUser = "users"

// User mapped from table <users>
type User struct {
	UserID         uint64         `gorm:"column:user_id;type:bigint(20) unsigned;primaryKey;autoIncrement:true" json:"user_id"`
	Email          string         `gorm:"column:email;type:varchar(100);not null;uniqueIndex:email,priority:1" json:"email"`
	Password       string         `gorm:"column:password;type:varchar(64);not null" json:"password"`
	Name           string         `gorm:"column:name;type:varchar(50);not null;uniqueIndex:name,priority:1" json:"name"`
	Slogan         string         `gorm:"column:slogan;type:varchar(150);not null" json:"slogan"`
	Role           string         `gorm:"column:role;type:varchar(20);not null;default:reader" json:"role"`
	FollowersCount int32          `gorm:"column:followers_count;type:int(16);not null" json:"followers_count"`
	FollowingCount int32          `gorm:"column:following_count;type:int(16);not null" json:"following_count"`
	LikeCount      int32          `gorm:"column:like_count;type:int(16);not null" json:"like_count"`
	DislikeCount   int32          `gorm:"column:dislike_count;type:int(16);not null" json:"dislike_count"`
	PostsCount     int32          `gorm:"column:posts_count;type:int(16);not null" json:"posts_count"`
	Online         bool           `gorm:"column:online;type:tinyint(1);not null" json:"online"`
	AccountStatus  string         `gorm:"column:account_status;type:varchar(20);not null;default:inactived" json:"account_status"`
	CreatedAt      time.Time      `gorm:"column:created_at;type:timestamp;not null;default:current_timestamp()" json:"created_at"`
	DeletedAt      gorm.DeletedAt `gorm:"column:deleted_at;type:timestamp" json:"deleted_at"`
	UpdatedAt      time.Time      `gorm:"column:updated_at;type:timestamp" json:"updated_at"`
}

// TableName User's table name
func (*User) TableName() string {
	return TableNameUser
}
