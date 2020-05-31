package model

type Weather struct {
	ID          int     `json:"id"`
	Main        string  `json:"main"`
	Description *string `json:"description"`
	Icon        string  `json:"icon"`
}
