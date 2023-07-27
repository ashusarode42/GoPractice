package Models

import "gorm"

type Address struct {
	HouseNo int    `json:"HouseNo"`
	City    string `json:"City"`
	State   string `json:"State"`
	PinCode string `json:"PinCode"`
	UserId  int
	gorm.Model
}
