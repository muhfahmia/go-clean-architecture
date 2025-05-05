package entity

import (
	"time"

	"github.com/muhfahmia/pkg/enum"
)

type UserEntity struct {
	UserID      uint64          `gorm:"primaryKey;autoIncrement;index:uid_idx"`
	Name        string          `gorm:"type:varchar(150);"`
	Username    string          `gorm:"type:varchar(50);"`
	Password    string          `gorm:"type:varchar(100);"`
	Email       string          `gorm:"type:varchar(100);"`
	Msisdn      string          `gorm:"type:varchar(20);"`
	LastLoginAt time.Time       `gorm:"type:timestamp(0);precision:6"`
	Status      enum.UserStatus `gorm:"tyoe:user_status;default:'active'"`
	Timestamp
}
