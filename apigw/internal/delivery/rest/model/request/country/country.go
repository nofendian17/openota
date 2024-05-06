package country

type GetByID struct {
	ID string `json:"id" validate:"required,uuid"`
}

type GetByCode struct {
	Code string `json:"code" validate:"required,max=2"`
}
