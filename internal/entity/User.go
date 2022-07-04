package entity

import (
	"gorm.io/gorm"
	"time"
)

type User struct {
	Id        int       `json:"id" gorm:"primaryKey"`
	Key       string    `json:"key"`
	Value     string    `json:"value"`
	Timestamp time.Time `gorm:"type:time" json:"timestamp,omitempty"`
}
type UserModel struct {
	Db *gorm.DB
}
