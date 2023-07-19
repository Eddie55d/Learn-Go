package icecreamrepo

import (
	"ice-cream-app/internal/models"
	"time"

	"fmt"

	"testing"

	"github.com/DATA-DOG/go-sqlmock"
)

func TestIcecreamRepoGetId(t *testing.T) {
	db, mock, err := sqlmock.New(sqlmock.QueryMatcherOption(sqlmock.QueryMatcherEqual))
	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}
	defer db.Close()

	fmt.Println("Test GetIceCreamByID success")

	var (
		iceCreamID          uint = 1
		iceCreamTitle            = "ice1"
		iceCreamComposition      = "some composition"
		iceCreamDateStart        = time.Now()
		iceCreamDateEnd          = time.Now()
		iceCreamPrice            = "178"
	)

	expectedRows := mock.NewRows([]string{
		"icecream_id",
		"title",
		"composition",
		"date_of_manufacture",
		"expiration_date",
		"price",
	}).AddRow(iceCreamID, iceCreamTitle, iceCreamComposition, iceCreamDateStart, iceCreamDateEnd, iceCreamPrice)

	mock.ExpectQuery(`SELECT icecream_id, 
														title, 
														composition, 
														date_of_manufacture, 
														expiration_date, 
														price 
														FROM icecreams 
														WHERE icecream_id = $1 AND is_deleted IS NULL`).
		WithArgs("1").WillReturnRows(expectedRows)

	var repo IceCreamDbStore

	repo.db = db

	res, err := repo.GetIceCreamByID("1")
	if err != nil {
		t.Error(err)
		return
	}

	if res.IceCreamID != iceCreamID {
		t.Errorf("Expected: %d, got: %d", iceCreamID, res.IceCreamID)
		return
	}
	if res.Title != iceCreamTitle {
		t.Errorf("Expected: %s, got: %s", iceCreamTitle, res.Title)
		return
	}
	if res.Composition != iceCreamComposition {
		t.Errorf("Expected: %s, got: %s", iceCreamComposition, res.Composition)
		return
	}
	if res.DateOfManufacture != iceCreamDateStart {
		t.Errorf("Expected: %s, got: %s", iceCreamDateStart, res.DateOfManufacture)
		return
	}
	if res.ExpirationDate != iceCreamDateEnd {
		t.Errorf("Expected: %s, got: %s", iceCreamDateEnd, res.ExpirationDate)
		return
	}
	if res.Price != iceCreamPrice {
		t.Errorf("Expected: %s, got: %s", iceCreamPrice, res.Price)
		return
	}

}

func TestIcecreamRepoGetIdErr(t *testing.T) {
	db, mock, err := sqlmock.New(sqlmock.QueryMatcherOption(sqlmock.QueryMatcherEqual))
	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}
	defer db.Close()

	fmt.Println("Test GetIceCreamByID error")

	expectedRows := mock.NewRows([]string{
		"icecream_id",
		"title",
		"composition",
		"date_of_manufacture",
		"expiration_date",
		"price",
	}).RowError(0, fmt.Errorf("sql: no rows in result set"))

	mock.ExpectQuery(`SELECT icecream_id, 
													 title, 
													 composition, 
					   							 date_of_manufacture, 
													 expiration_date, 
													 price 
													 FROM icecreams 
													 WHERE icecream_id = $1 AND is_deleted IS NULL`).
		WithArgs("1").WillReturnRows(expectedRows)

	var repo IceCreamDbStore

	repo.db = db

	_, err = repo.GetIceCreamByID("1")

	if err == nil {
		t.Errorf("Expected: error models.ErrSqlNoRow, got: error nil")
		return
	}
	if err != nil && err != models.ErrSqlNoRow {
		t.Errorf("Expected: error models.ErrSqlNoRow, got: %v", err)
		return
	}

}

func TestIcecreamRepoDeleteIceCream(t *testing.T) {
	db, mock, err := sqlmock.New(sqlmock.QueryMatcherOption(sqlmock.QueryMatcherEqual))
	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}
	defer db.Close()

	fmt.Println("Test DeleteIceCream success")

	var (
		lastInsertID int64 = 1
		affected     int64 = 1
	)

	result := sqlmock.NewResult(lastInsertID, affected)

	mock.ExpectExec(`UPDATE icecreams SET is_deleted = $2 WHERE icecream_id = $1 AND is_deleted IS NULL`).WillReturnResult(result)

	var repo IceCreamDbStore

	repo.db = db

	err = repo.DeleteIceCream("1")
	if err != nil && err != models.ErrSqlNoRow {
		t.Error(err)
		return
	}

	if err == models.ErrSqlNoRow {
		t.Errorf("Expected: %d row deleted, got: %d row deleted", affected, 0)
		return
	}

}

func TestIcecreamRepoDeleteIceCreamErr(t *testing.T) {
	db, mock, err := sqlmock.New(sqlmock.QueryMatcherOption(sqlmock.QueryMatcherEqual))
	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}
	defer db.Close()

	fmt.Println("Test DeleteIceCream error")

	var (
		lastInsertID int64 = 0
		affected     int64 = 0
	)

	result := sqlmock.NewResult(lastInsertID, affected)
	mock.ExpectExec(`UPDATE icecreams SET is_deleted = $2 WHERE icecream_id = $1 AND is_deleted IS NULL`).WillReturnResult(result)

	var repo IceCreamDbStore

	repo.db = db

	err = repo.DeleteIceCream("1")
	if err == nil {
		t.Errorf("Expected: error models.ErrSqlNoRow, got: error nil")
		return
	}
	if err != nil && err != models.ErrSqlNoRow {
		t.Errorf("Expected: error models.ErrSqlNoRow, got: %v", err)
		return
	}

}
