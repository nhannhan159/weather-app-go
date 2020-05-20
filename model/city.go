package model

type City struct {
	Id       int     `json:"id" gorm:"primary_key"`
	Name     string  `json:"name"`
	FindName *string `json:"findName"`
	Country  *string `json:"country"`
	Lat      float64 `json:"lat"`
	Lon      float64 `json:"lon"`
}

func (this City) GetPrimaryKey() int {
	return this.Id
}
