// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.

package models

const TableNamePostsTag = "posts_tags"

// PostsTag mapped from table <posts_tags>
type PostsTag struct {
	PostID uint64 `gorm:"column:post_id;type:bigint(20) unsigned;not null;index:post_id,priority:1" json:"post_id"`
	TagID  uint64 `gorm:"column:tag_id;type:bigint(20) unsigned;not null;index:tag_id,priority:1" json:"tag_id"`
}

// TableName PostsTag's table name
func (*PostsTag) TableName() string {
	return TableNamePostsTag
}
