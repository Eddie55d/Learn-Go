package models

import "time"

type IceCream struct {
	Icecream_id       uint      `json:"icecream_id"`
	Title             string    `json:"title"`
	Сomposition       string    `json:"composition"`
	DateOfManufacture time.Time `json:"date_of_manufacture"`
	ExpirationDate    time.Time `json:"expiration_date"`
	Price             string    `json:"price"`
}

type IceCreamPost struct {
	Title             string `json:"title"`
	Сomposition       string `json:"composition"`
	DateOfManufacture string `json:"date_of_manufacture"`
	ExpirationDate    string `json:"expiration_date"`
	Price             string `json:"price"`
}

type IceCreamUp struct {
	Title             string `json:"title"`
	Сomposition       string `json:"composition"`
	DateOfManufacture string `json:"date_of_manufacture"`
	ExpirationDate    string `json:"expiration_date"`
	Price             string `json:"price"`
}
