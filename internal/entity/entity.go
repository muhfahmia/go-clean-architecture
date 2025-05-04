package entity

import "time"

type Timestamp struct {
	CreatedAt time.Time `gorm:"type:timestamp(0);autoCreateTime"`
	UpdatedAt time.Time `gorm:"type:timestamp(0);autoUpdateTime"`
}
