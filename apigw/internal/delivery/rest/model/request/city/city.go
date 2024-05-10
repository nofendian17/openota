package city

type GetByID struct {
	ID string `json:"id" validate:"required,uuid"`
}

type GetByStateID struct {
	ID string `json:"id" validate:"required,uuid"`
}

type Create struct {
	StateID   string  `json:"state_id"`
	Name      string  `json:"name"`
	Latitude  float64 `json:"latitude"`
	Longitude float64 `json:"longitude"`
	IsActive  *bool   `json:"is_active"`
}

type Update struct {
	ID         string  `json:"id" validate:"required,uuid"`
	StateID    string  `json:"state_id"`
	Name       string  `json:"name"`
	Latitude   float64 `json:"latitude"`
	Longitude  float64 `json:"longitude"`
	IsActive   *bool   `json:"is_active"`
	Precedence int64   `json:"precedence" validate:"required,min=0"`
}

type Delete struct {
	ID string `json:"id" validate:"required,uuid"`
}
