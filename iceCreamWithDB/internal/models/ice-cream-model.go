package models

import (
	"strconv"
	"time"

	"github.com/sirupsen/logrus"
)

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
	Composition       string `json:"composition"`
	DateOfManufacture string `json:"date_of_manufacture"`
	ExpirationDate    string `json:"expiration_date"`
	Price             string `json:"price"`
}

type IceCreamPostValid struct {
	Title             string    `json:"title"`
	Сomposition       string    `json:"composition"`
	DateOfManufacture time.Time `json:"date_of_manufacture"`
	ExpirationDate    time.Time `json:"expiration_date"`
	Price             float64   `json:"price"`
}

// Validation POST ice cream

func (newIceCream *IceCreamPost) ValidPost() (validIceCream *IceCreamPostValid, msg string, err error) {

	var message string

	if newIceCream.Title == "" {
		message += "title field is empty"
		return nil, message, err
	}

	if newIceCream.Composition == "" {
		message += "composition field is empty"
		return nil, message, err
	}

	if newIceCream.Price == "" {
		message += "price field is empty"
		return nil, message, err
	}

	tm, err := time.Parse("2006-01-02", newIceCream.DateOfManufacture)
	if err != nil {
		logrus.Warn(err)
	}

	if tm.IsZero() {
		message += "The Date Of Manufacture field is incorrect or empty"
		return nil, message, err
	}

	tExp, err := time.Parse("2006-01-02", newIceCream.ExpirationDate)
	if err != nil {
		logrus.Warn(err)
	}

	if tExp.IsZero() {
		message += "The Expiration Date field is incorrect or empty"
		return nil, message, err
	}

	dtST, err := time.Parse("2006-01-02", newIceCream.DateOfManufacture)
	if err != nil {
		message += "something is wrong with the date"
		return nil, message, err
	}

	dtEnd, err := time.Parse("2006-01-02", newIceCream.ExpirationDate)
	if err != nil {
		message += "something is wrong with the date"
		return nil, message, err
	}

	pr, err := strconv.ParseFloat(newIceCream.Price, 64)
	if err != nil {
		message += "something is wrong with the price"
		return nil, message, err
	}

	newIceCreamValid := IceCreamPostValid{
		Title:             newIceCream.Title,
		Сomposition:       newIceCream.Composition,
		DateOfManufacture: dtST,
		ExpirationDate:    dtEnd,
		Price:             pr,
	}

	return &newIceCreamValid, message, err
}

// Validation PUT ice cream

func (upIce *IceCreamPost) ValidUpdate(iceCream *IceCream) (validIceCream *IceCreamPostValid, msg string, err error) {

	message := ""

	if upIce.Title == "" {
		upIce.Title = iceCream.Title
	}

	if upIce.Composition == "" {
		upIce.Composition = iceCream.Сomposition
	}

	if upIce.DateOfManufacture == "" {
		upIce.DateOfManufacture = iceCream.DateOfManufacture.Format("2006-01-02")
	}

	if upIce.ExpirationDate == "" {
		upIce.ExpirationDate = iceCream.ExpirationDate.Format("2006-01-02")
	}

	if upIce.Price == "" {
		upIce.Price = iceCream.Price
	}

	dtST, err := time.Parse("2006-01-02", upIce.DateOfManufacture)
	if err != nil {
		logrus.Warn(err)
		message += "something is wrong with the Date Of Manufacture"
		return nil, message, err
	}

	dtEnd, err := time.Parse("2006-01-02", upIce.ExpirationDate)
	if err != nil {
		logrus.Warn(err)
		message += "something is wrong with the Expiration Date"
		return nil, message, err
	}

	pr, err := strconv.ParseFloat(upIce.Price, 32)
	if err != nil {
		logrus.Warn(err)
		message += "something is wrong with the price"
		return nil, message, err
	}

	newIceCreamValidUpdate := IceCreamPostValid{
		Title:             upIce.Title,
		Сomposition:       upIce.Composition,
		DateOfManufacture: dtST,
		ExpirationDate:    dtEnd,
		Price:             pr,
	}

	return &newIceCreamValidUpdate, message, err
}
