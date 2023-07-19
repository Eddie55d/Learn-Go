package models

import (
	"errors"
	"fmt"
	"strconv"
	"time"
)

// error Get Ice Cream
var ErrSqlNoRow = errors.New("sql: no rows in result set")

// Ice Cream
type IceCream struct {
	IceCreamID        uint      `json:"icecream_id"`
	Title             string    `json:"title"`
	Composition       string    `json:"composition"`
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

// errors POST ice cream
var (
	ErrTitleEmpty           = errors.New("title field is empty")
	ErrCompositionEmpty     = errors.New("composition field is empty")
	ErrPriceEmpty           = errors.New("price field is zero or negative")
	ErrPrice                = errors.New("price field value is not a digit")
	ErrDateOfManufacture    = errors.New("dateOfManufacture field is empty or invalid")
	ErrExpirationDate       = errors.New("expirationDate field is empty or invalid")
	ErrExpirationDateBefore = errors.New("date of manufacture must be before the expiration date")
)

// Validation POST ice cream
func (i *IceCreamPost) Validate() error {
	if i.Title == "" {
		return ErrTitleEmpty
	}
	if i.Composition == "" {
		return ErrCompositionEmpty
	}

	dateManufacture, err := time.Parse("2006-01-02", i.DateOfManufacture)
	if err != nil {
		fmt.Println(dateManufacture)
		return ErrDateOfManufacture
	}
	if dateManufacture.IsZero() {
		return ErrDateOfManufacture
	}

	expirationDate, err := time.Parse("2006-01-02", i.ExpirationDate)
	if err != nil {
		return ErrExpirationDate
	}
	if expirationDate.IsZero() {
		return ErrExpirationDate
	}

	dateBefore := dateManufacture.Before(expirationDate)
	if !dateBefore {
		return ErrExpirationDateBefore
	}

	pr, err := strconv.ParseFloat(i.Price, 64)
	if err != nil {
		return ErrPrice
	}
	if pr == 0 || pr < 0 {
		return ErrPriceEmpty
	}

	return nil
}

// Get All Ice Creams params
type QueryParams struct {
	Limit, Offset                          string
	ManufactureDate, ExpirationDate, Price bool
}

// Update validation Ice Cream
func (updateIceCream *IceCreamPost) ValidUpdate(iceCream *IceCream) (*IceCreamPost, error) {

	if updateIceCream.Title == "" {
		updateIceCream.Title = iceCream.Title
	}

	if updateIceCream.Composition == "" {
		updateIceCream.Composition = iceCream.Composition
	}

	if updateIceCream.DateOfManufacture == "" {
		updateIceCream.DateOfManufacture = iceCream.DateOfManufacture.Format("2006-01-02")
	}

	if updateIceCream.ExpirationDate == "" {
		updateIceCream.ExpirationDate = iceCream.ExpirationDate.Format("2006-01-02")
	}

	if updateIceCream.Price == "" {
		updateIceCream.Price = iceCream.Price
	}

	err := updateIceCream.Validate()
	if err != nil {
		return nil, err
	}

	newIceCreamValidUpdate := IceCreamPost{
		Title:             updateIceCream.Title,
		Composition:       updateIceCream.Composition,
		DateOfManufacture: updateIceCream.DateOfManufacture,
		ExpirationDate:    updateIceCream.ExpirationDate,
		Price:             updateIceCream.Price,
	}

	return &newIceCreamValidUpdate, err
}
