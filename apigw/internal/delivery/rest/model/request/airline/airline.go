package airline

type GetByID struct {
	ID string `json:"id" validate:"required,uuid"`
}

type GetByCode struct {
	Code string `json:"code" validate:"required,max=2"`
}

type Create struct {
	Code     string `json:"code" validate:"required,min=2,max=2"`
	Name     string `json:"name" validate:"required,max=255"`
	Logo     string `json:"logo" validate:"required,url,max=255"`
	IsActive *bool  `json:"is_active" validate:"required,boolean"`
}

type Update struct {
	ID         string `json:"id" validate:"required,uuid"`
	Code       string `json:"code" validate:"required,min=2,max=2"`
	Name       string `json:"name" validate:"required,max=255"`
	Logo       string `json:"logo" validate:"required,url,max=255"`
	IsActive   *bool  `json:"is_active" validate:"required,boolean"`
	Precedence int64  `json:"precedence" validate:"required,min=0"`
}

type Delete struct {
	ID string `json:"id" validate:"required,uuid"`
}
