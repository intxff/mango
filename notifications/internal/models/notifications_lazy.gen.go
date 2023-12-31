// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.

package models

import (
	"time"

	"gorm.io/gorm"
)

const TableNameNotificationsLazy = "notifications_lazy"

// NotificationsLazy mapped from table <notifications_lazy>
type NotificationsLazy struct {
	NotificationType string         `gorm:"column:notification_type;type:varchar(20);not null;index:notification_type,priority:1" json:"notification_type"`
	ObjectID         uint64         `gorm:"column:object_id;type:bigint(20) unsigned;not null;index:notification_type,priority:2" json:"object_id"`
	LastScanTime     time.Time      `gorm:"column:last_scan_time;type:timestamp;not null" json:"last_scan_time"`
	CreatedAt        time.Time      `gorm:"column:created_at;type:timestamp;not null;default:current_timestamp()" json:"created_at"`
	DeletedAt        gorm.DeletedAt `gorm:"column:deleted_at;type:timestamp" json:"deleted_at"`
	UpdatedAt        time.Time      `gorm:"column:updated_at;type:timestamp" json:"updated_at"`
}

// TableName NotificationsLazy's table name
func (*NotificationsLazy) TableName() string {
	return TableNameNotificationsLazy
}
