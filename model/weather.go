package model

type Weather struct {
	Id          int     `json:"id" gorm:"primary_key"`
	Main        string  `json:"main"`
	Description *string `json:"desciption"`
	Icon        string  `json:"icon"`
}

func (weather Weather) GetPrimaryKey() int {
	return weather.Id
}
