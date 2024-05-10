package city

type GetByID struct {
	ID string `json:"id" validate:"required,uuid"`
}

type GetByStateID struct {
	ID string `json:"id" validate:"required,uuid"`
}

type Create struct {
	StateID   string  `json:"state_id" validate:"required,uuid"`
	Name      string  `json:"name" validate:"required"`
	Latitude  float64 `json:"latitude" validate:"required,latitude"`
	Longitude float64 `json:"longitude" validate:"required,longitude"`
	IsActive  *bool   `json:"is_active" validate:"required,boolean"`
}

type Update struct {
	ID         string  `json:"id" validate:"required,uuid"`
	StateID    string  `json:"state_id" validate:"required,uuid"`
	Name       string  `json:"name" validate:"required,max=255"`
	Latitude   float64 `json:"latitude" validate:"required,latitude"`
	Longitude  float64 `json:"longitude" validate:"required,longitude"`
	IsActive   *bool   `json:"is_active" validate:"required,boolean"`
	Precedence int64   `json:"precedence" validate:"required,min=0"`
}

type Delete struct {
	ID string `json:"id" validate:"required,uuid"`
}
