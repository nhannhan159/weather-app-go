package model

type Weather struct {
	Id          int     `json:"id" gorm:"primary_key"`
	Main        string  `json:"main"`
	Description *string `json:"desciption"`
	Icon        string  `json:"icon"`
}

func (this Weather) GetPrimaryKey() int {
	return this.Id
}
