package model

type Sticker struct {
	ID             int    `gorm:"primaryKey"`
	Name           string `gorm:"name,type:varchar(255)"`
	Description    string `gorm:"description,type:varchar(255)"`
	Category       string `gorm:"category,size:30"`
	Frequency      string `gorm:"frequency,size:30"`
	Status         string `gorm:"status,size:30"`
	IsPublic       bool   `gorm:"is_public"`
	IsAutoApproval bool   `gorm:"is_auto_approval"`
	User           User   `gorm:"foreignkey:UserId"`
	UserId         int    `gorm:"user_id"`
}
