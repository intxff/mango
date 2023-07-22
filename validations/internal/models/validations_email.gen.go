// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.

package models

import (
	"time"

	"gorm.io/gorm"
)

const TableNameValidationsEmail = "validations_email"

// ValidationsEmail mapped from table <validations_email>
type ValidationsEmail struct {
	ValidationEmailID uint64         `gorm:"column:validation_email_id;type:bigint(20) unsigned;not null;index:validation_email_id,priority:1" json:"validation_email_id"`
	Code              string         `gorm:"column:code;type:varchar(20);not null" json:"code"`
	CreatedAt         time.Time      `gorm:"column:created_at;type:timestamp;not null;default:current_timestamp()" json:"created_at"`
	DeletedAt         gorm.DeletedAt `gorm:"column:deleted_at;type:timestamp" json:"deleted_at"`
	UpdatedAt         time.Time      `gorm:"column:updated_at;type:timestamp" json:"updated_at"`
}

// TableName ValidationsEmail's table name
func (*ValidationsEmail) TableName() string {
	return TableNameValidationsEmail
}