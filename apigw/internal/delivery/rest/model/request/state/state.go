package state

type GetByID struct {
	ID string `json:"id" validate:"required,uuid"`
}

type GetByCountryID struct {
	ID string `json:"id" validate:"required,uuid"`
}

type Create struct {
	Name      string  `json:"name" validate:"required,max=50"`
	CountryID string  `json:"country_id" validate:"required,uuid"`
	Latitude  float64 `json:"latitude" validate:"required"`
	Longitude float64 `json:"longitude" validate:"required"`
	IsActive  *bool   `json:"is_active" validate:"required"`
}

type Update struct {
	ID         string  `json:"id" validate:"required,uuid"`
	Name       string  `json:"name" validate:"required,max=50"`
	CountryID  string  `json:"country_id" validate:"required,uuid"`
	Latitude   float64 `json:"latitude" validate:"required"`
	Longitude  float64 `json:"longitude" validate:"required"`
	IsActive   *bool   `json:"is_active" validate:"required,boolean"`
	Precedence int64   `json:"precedence" validate:"required,min=0"`
}

type Delete struct {
	ID string `json:"id" validate:"required,uuid"`
}
