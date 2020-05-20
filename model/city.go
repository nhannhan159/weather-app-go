package model

type City struct {
	ID       int     `json:"id"`
	Name     string  `json:"name"`
	FindName *string `json:"findName"`
	Country  *string `json:"country"`
	Lat      float64 `json:"lat"`
	Lon      float64 `json:"lon"`
}
