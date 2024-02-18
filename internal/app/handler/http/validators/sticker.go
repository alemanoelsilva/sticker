package validators

type Sticker struct {
	Name           string `validate:"required"`
	Description    string `validate:"required"`
	Category       string `validate:"required"`
	Frequency      string `validate:"required"`
	Status         string `validate:"required"`
	IsPublic       string `validate:"oneof=true false"`
	IsAutoApproval string `validate:"oneof=true false"`
}
