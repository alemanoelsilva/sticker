package model

type Sticker struct {
	ID             int    `db:"id"`
	Name           string `db:"name"`
	Description    string `db:"description"`
	Category       string `db:"category"`
	Frequency      string `db:"frequency"`
	Status         string `db:"status"`
	IsPublic       bool   `db:"is_public"`
	IsAutoApproved bool   `db:"is_auto_approved"`
}
