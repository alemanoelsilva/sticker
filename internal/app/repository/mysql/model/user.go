package model

type User struct {
	ID       int    `gorm:"primaryKey"`
	Name     string `gorm:"name,unique,size:255"`
	Email    string `gorm:"email,size:255"`
	Password string `gorm:"password,size:255"`
}
