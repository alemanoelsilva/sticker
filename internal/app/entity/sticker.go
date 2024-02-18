package entity

type CategoryType string

const (
	FINANCIAL CategoryType = "FINANCIAL"
	FITNESS   CategoryType = "FITNESS"
	HEALTH    CategoryType = "HEALTH"
	BUSINESS  CategoryType = "BUSINESS"
)

type FrequencyType string

const (
	DAILY   FrequencyType = "DAILY"
	WEEKLY  FrequencyType = "WEEKLY"
	MONTHLY FrequencyType = "MONTHLY"
)

type StatusType string

const (
	ACTIVE   StatusType = "ACTIVE"
	INACTIVE StatusType = "INACTIVE"
	DRAFT    StatusType = "DRAFT"
	ON_HOLD  StatusType = "ON_HOLD"
)

type Sticker struct {
	ID             int           `json:"id"`
	Name           string        `json:"name"`
	Description    string        `json:"description"`
	Category       CategoryType  `json:"category"`
	Frequency      FrequencyType `json:"frequency"`
	Status         StatusType    `json:"status"`
	IsPublic       bool          `json:"isPublic"`
	IsAutoApproved bool          `json:"isAutoApproved"`
}
