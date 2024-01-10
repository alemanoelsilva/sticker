package entities

type Sticker struct {
	ID             uint   `json:"id"`
	Name           string `json:"name"`
	Description    string `json:"description"`
	Status         string `json:"status"`
	IsPublic       bool   `json:"isPublic"`
	IsAutoApproved bool   `json:"isAutoApproved"`
}
