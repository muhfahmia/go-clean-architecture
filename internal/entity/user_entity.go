package entity

type UserEntity struct {
    UserID uint64 `gorm:"primaryKey;autoIncrement;index:user_idx"`
    // Add other fields here
    Timestamp
}
