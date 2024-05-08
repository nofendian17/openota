package country

type GetByID struct {
	ID string `json:"id" validate:"required,uuid"`
}

type GetByCode struct {
	Code string `json:"code" validate:"required,max=2"`
}

type Create struct {
	Name         string  `json:"name" validate:"required,max=50"`
	Code         string  `json:"code" validate:"required,min=3,max=3"`
	PhoneCode    string  `json:"phone_code" validate:"required,max=5"`
	Capital      string  `json:"capital" validate:"required,max=50"`
	Latitude     float64 `json:"latitude" validate:"required"`
	Longitude    float64 `json:"longitude" validate:"required"`
	CurrencyCode string  `json:"currency" validate:"required,min=3,max=3"`
	IsActive     *bool   `json:"is_active" validate:"required"`
}

type Update struct {
	ID           string  `json:"id" validate:"required,uuid"`
	Name         string  `json:"name" validate:"required,max=50"`
	Code         string  `json:"code" validate:"required,min=3,max=3"`
	PhoneCode    string  `json:"phone_code" validate:"required,max=5"`
	Capital      string  `json:"capital" validate:"required,max=50"`
	Latitude     float64 `json:"latitude" validate:"required"`
	Longitude    float64 `json:"longitude" validate:"required"`
	CurrencyCode string  `json:"currency" validate:"required,min=3,max=3"`
	IsActive     *bool   `json:"is_active" validate:"required"`
}

type Delete struct {
	ID string `json:"id" validate:"required,uuid"`
}