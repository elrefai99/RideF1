package user

type User struct {
	ID           int64  `gorm:"primaryKey" json:"id"`
	Name         string `gorm:"not null" json:"name"`
	Email        string `gorm:"uniqueIndex;not null" json:"email"`
	PasswordHash string `gorm:"not null" json:"-"`
}
