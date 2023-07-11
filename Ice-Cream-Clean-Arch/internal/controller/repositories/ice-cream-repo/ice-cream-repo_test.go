package icecreamrepo_test

import (
	"database/sql"

	"fmt"

	"testing"

	"github.com/DATA-DOG/go-sqlmock"
)

func TestIcecreamRepoGetId(t *testing.T) {
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}
	defer db.Close()

	fmt.Println("Test #1")

	expectedRows := mock.NewRows([]string{
		"icecream_id",
		"title",
		"composition",
		"date_of_manufacture",
		"expiration_date",
		"price",
	}).AddRow(1, "ice1", "hdwpoicdw", "2023-03-04", "2023-03-05", "178")

	mock.ExpectQuery("SELECT icecream_id, title, composition, date_of_manufacture, expiration_date, price FROM icecreams WHERE icecream_id = ?").WithArgs(1).WillReturnRows(expectedRows)

	var (
		IceCreamID        uint64
		Title             string
		Composition       string
		DateOfManufacture string
		ExpirationDate    string
		Price             string
	)

	err = db.QueryRow(`SELECT 
	icecream_id, 
	title, 
	composition, 
	date_of_manufacture, 
	expiration_date, 
	price 
	FROM icecreams WHERE icecream_id = ?`, 1).
		Scan(
			&IceCreamID,
			&Title,
			&Composition,
			&DateOfManufacture,
			&ExpirationDate,
			&Price)

	if err != nil {
		fmt.Println(err)
		return
	}
	if IceCreamID == 1 {
		fmt.Printf("Get ice cream with id: %d\n", IceCreamID)
		return
	}

}

func TestIcecreamRepoGetIdErr(t *testing.T) {
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}
	defer db.Close()

	fmt.Println("Test #2")

	expectedRows := mock.NewRows([]string{
		"icecream_id",
		"title",
		"composition",
		"date_of_manufacture",
		"expiration_date",
		"price",
	}).RowError(0, fmt.Errorf("ice cream not found"))

	mock.ExpectQuery(
		`SELECT 
			icecream_id, 
			title, 
			composition, 
			date_of_manufacture, 
			expiration_date, 
			price FROM icecreams WHERE icecream_id = ?
	`).WithArgs(2).WillReturnRows(expectedRows)

	var (
		IceCreamID        uint64
		Title             string
		Composition       string
		DateOfManufacture string
		ExpirationDate    string
		Price             string
	)

	err = db.QueryRow(`SELECT 
		icecream_id, 
		title, 
		composition, 
		date_of_manufacture, 
		expiration_date, 
		price 
		FROM icecreams WHERE icecream_id = ?`, 2).
		Scan(
			&IceCreamID,
			&Title,
			&Composition,
			&DateOfManufacture,
			&ExpirationDate,
			&Price)

	if err != nil && err != sql.ErrNoRows {
		fmt.Println(err)
		return
	}

	if err == sql.ErrNoRows {
		fmt.Println("ice cream not found")
		return
	}

}
