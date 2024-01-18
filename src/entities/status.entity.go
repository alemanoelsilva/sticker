package entities

// EnumCategory represents the category enumeration.
type StatusType string

const (
	ACTIVE   StatusType = "ACTIVE"
	INACTIVE StatusType = "INACTIVE"
	DRAFT    StatusType = "DRAFT"
	HOLYDAY  StatusType = "HOLYDAY"
)
