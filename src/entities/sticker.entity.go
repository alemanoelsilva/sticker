package entities

type Sticker struct {
	ID             uint          `json:"id"`
	Name           string        `json:"name"`
	Description    string        `json:"description"`
	Category       CategoryType  `json:"category"`
	Frequency      FrequencyType `json:"frequency"`
	Status         StatusType    `json:"status"`
	IsPublic       bool          `json:"isPublic"`
	IsAutoApproved bool          `json:"isAutoApproved"`
}
