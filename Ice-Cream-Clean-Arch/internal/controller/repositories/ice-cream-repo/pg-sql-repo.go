package icecreamrepo

import (
	"database/sql"
	"errors"
	"fmt"

	"time"

	"ice-cream-app/internal/models"

	"github.com/sirupsen/logrus"
)

type IceCreamDbStore struct {
	db *sql.DB
}

func NewIceCreamDbStore(db *sql.DB) *IceCreamDbStore {

	return &IceCreamDbStore{db: db}
}

// Get All Ice Creams
func (icecream *IceCreamDbStore) GetIceCreams(params models.QueryParams) ([]models.IceCream, error) {

	const query = "SELECT icecream_id, title, composition, date_of_manufacture, expiration_date, price FROM icecreams WHERE is_deleted IS NULL ORDER BY"

	qp := params

	var selectRow string

	switch {
	case qp.ManufactureDate:
		selectRow = query + " date_of_manufacture"
	case qp.ExpirationDate:
		selectRow = query + " expiration_date"
	case qp.Price:
		selectRow = query + " price"
	default:
		selectRow = query + " icecream_id"
	}

	if qp.Limit != "" {
		selectRow += fmt.Sprintf(` LIMIT %s`, qp.Limit)
	}

	if qp.Offset != "" {
		selectRow += fmt.Sprintf(` OFFSET %s`, qp.Offset)
	}

	rows, err := icecream.db.Query(selectRow)
	if err != nil {
		logrus.Error(err)
		return nil, err
	}
	defer rows.Close()

	iceCreams := []models.IceCream{}

	for rows.Next() {
		icecream := models.IceCream{}
		err := rows.Scan(
			&icecream.IceCreamID,
			&icecream.Title,
			&icecream.Composition,
			&icecream.DateOfManufacture,
			&icecream.ExpirationDate,
			&icecream.Price)
		if err != nil {
			logrus.Error(err)
			return nil, err
		}
		iceCreams = append(iceCreams, icecream)
	}

	return iceCreams, err
}

// Get Ice Cream By ID
func (icecream *IceCreamDbStore) GetIceCreamByID(paramID string) (models.IceCream, error) {

	id := paramID
	var iceCream models.IceCream

	err := icecream.db.QueryRow(`SELECT icecream_id, 
																			title, 
																			composition, 
																			date_of_manufacture, 
																			expiration_date, 
																			price 
																			FROM icecreams 
																			WHERE icecream_id = $1 AND is_deleted IS NULL`, id).
		Scan(
			&iceCream.IceCreamID,
			&iceCream.Title,
			&iceCream.Composition,
			&iceCream.DateOfManufacture,
			&iceCream.ExpirationDate,
			&iceCream.Price,
		)

	switch {
	case err != nil && !errors.Is(err, sql.ErrNoRows):
		return iceCream, err
	case errors.Is(err, sql.ErrNoRows):
		return iceCream, models.ErrSqlNoRow
	default:
		return iceCream, err
	}
}

// Delete Ice Cream
func (icecream *IceCreamDbStore) DeleteIceCream(paramID string) error {

	id := paramID
	deleteTime := time.Now()

	res, err := icecream.db.Exec("UPDATE icecreams SET is_deleted = $2 WHERE icecream_id = $1 AND is_deleted IS NULL",
		id, deleteTime)
	if err != nil {
		return err
	}

	deleteRow, err := res.RowsAffected()
	if err != nil && deleteRow != 0 {
		return err
	}
	if deleteRow == 0 {
		return models.ErrSqlNoRow
	}

	return nil
}

// Post Ice Cream
func (icecream *IceCreamDbStore) PostIceCream(iceCream models.IceCreamPost) (int64, error) {
	var newIceCream = iceCream

	const insertQuery = `
		INSERT INTO icecreams (title, 
		                       composition, 
		                       date_of_manufacture, 
		                       expiration_date, 
		                       price) 
		VALUES ($1, $2, $3, $4, $5) 
		RETURNING icecream_id`

	var icecreamId int64

	err := icecream.db.QueryRow(
		insertQuery,
		newIceCream.Title,             // $1, title
		newIceCream.Composition,       // $2, composition
		newIceCream.DateOfManufacture, // $3, date_of_manufacture
		newIceCream.ExpirationDate,    // $4, expiration_date
		newIceCream.Price,             // $5, price
	).Scan(&icecreamId)

	switch {
	case err != nil && !errors.Is(err, sql.ErrNoRows):
		return 0, err
	case errors.Is(err, sql.ErrNoRows):
		return 0, models.ErrSqlNoRow
	default:
		return icecreamId, err
	}

}

// Update Ice Cream
func (icecream *IceCreamDbStore) PutIceCream(paramID int, iceCream models.IceCreamPost) error {
	const insertQuery = `
	UPDATE icecreams SET 
	title=$1, 
	composition=$2, 
	date_of_manufacture=$3, 
	expiration_date=$4, 
	price=$5 
	WHERE icecream_id=$6`

	id := paramID
	var newIceCream = iceCream

	res, err := icecream.db.Exec(
		insertQuery,
		newIceCream.Title,             // $1, title
		newIceCream.Composition,       // $2, composition
		newIceCream.DateOfManufacture, // $3, date_of_manufacture
		newIceCream.ExpirationDate,    // $4, expiration_date
		newIceCream.Price,             // $5, price
		id,                            // $6, icecreamID
	)
	if err != nil {
		return err
	}

	updateRow, err := res.RowsAffected()
	if err != nil && updateRow != 0 {
		return err
	}
	if updateRow == 0 {
		return models.ErrSqlNoRow
	}

	return nil
}
