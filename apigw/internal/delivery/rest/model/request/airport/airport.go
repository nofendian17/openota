package airport

type GetByID struct {
	ID string `json:"id" validate:"required,uuid"`
}

type GetByCode struct {
	Code string `json:"code" validate:"required,max=2"`
}

type Create struct {
	Code       string  `json:"code" validate:"required,min=3,max=3"`
	Name       string  `json:"name" validate:"required,max=255"`
	CityID     string  `json:"city_id" validate:"required,uuid"`
	Latitude   float64 `json:"latitude" validate:"required,latitude"`
	Longitude  float64 `json:"longitude" validate:"required,longitude"`
	IsDomestic *bool   `json:"is_domestic" validate:"required,boolean"`
	Timezone   string  `json:"timezone" validate:"required,timezone"`
	IsActive   *bool   `json:"is_active" validate:"required,boolean"`
}

type Update struct {
	ID         string  `json:"id" validate:"required,uuid"`
	Code       string  `json:"code" validate:"required,min=3,max=3"`
	Name       string  `json:"name" validate:"required,max=255"`
	CityID     string  `json:"city_id" validate:"required,uuid"`
	Latitude   float64 `json:"latitude" validate:"required,latitude"`
	Longitude  float64 `json:"longitude" validate:"required,longitude"`
	IsDomestic *bool   `json:"is_domestic" validate:"required,boolean"`
	Timezone   string  `json:"timezone" validate:"required,timezone"`
	IsActive   *bool   `json:"is_active" validate:"required,boolean"`
	Precedence int64   `json:"precedence" validate:"required,min=0"`
}

type Delete struct {
	ID string `json:"id" validate:"required,uuid"`
}
