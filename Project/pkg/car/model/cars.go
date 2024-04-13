package model

import (
	"context"
	"database/sql"
	"fmt"
	"log"
	"time"

	//"github.com/codev0/inft3212-6/pkg/abr-plus/validator"
)

type Car struct {
	Id             string `json:"id"`
	CreatedAt      string `json:"createdAt"`
	UpdatedAt      string `json:"updatedAt"`
	Title          string `json:"title"`
	Description    string `json:"description"`
	Year uint   `json:"year"`
}

type CarModel struct {
	DB       *sql.DB
	InfoLog  *log.Logger
	ErrorLog *log.Logger
}

func (m CarModel) GetAll(title string, from, to int, filters Filters) ([]*Car, Metadata, error) {

	// Retrieve all car items from the database.
	query := fmt.Sprintf(
		`
		SELECT count(*) OVER(), id, created_at, updated_at, title, description, ye_ar
		FROM cars
		WHERE (LOWER(title) = LOWER($1) OR $1 = '')
		AND (ye_ar >= $2 OR $2 = 0)
		AND (ye_ar <= $3 OR $3 = 0)
		ORDER BY %s %s, id ASC
		LIMIT $4 OFFSET $5
		`,
		filters.sortColumn(), filters.sortDirection())

	// Create a context with a 3-second timeout.
	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()

	// Organize our four placeholder parameter values in a slice.
	args := []interface{}{title, from, to, filters.limit(), filters.offset()}

	// log.Println(query, title, from, to, filters.limit(), filters.offset())
	// Use QueryContext to execute the query. This returns a sql.Rows result set containing
	// the result.
	rows, err := m.DB.QueryContext(ctx, query, args...)
	if err != nil {
		return nil, Metadata{}, err
	}

	// Importantly, defer a call to rows.Close() to ensure that the result set is closed
	// before GetAll returns.
	defer func() {
		if err := rows.Close(); err != nil {
			m.ErrorLog.Println(err)
		}
	}()

	// Declare a totalRecords variable
	totalRecords := 0

	var cars []*Car
	for rows.Next() {
		var car Car
		err := rows.Scan(&totalRecords, &car.Id, &car.CreatedAt, &car.UpdatedAt, &car.Title, &car.Description, &car.Year)
		if err != nil {
			return nil, Metadata{}, err
		}

		// Add the Movie struct to the slice
		cars = append(cars, &car)
	}

	// When the rows.Next() loop has finished, call rows.Err() to retrieve any error
	// that was encountered during the iteration.
	if err = rows.Err(); err != nil {
		return nil, Metadata{}, err
	}

	// Generate a Metadata struct, passing in the total record count and pagination parameters
	// from the client.
	metadata := calculateMetadata(totalRecords, filters.Page, filters.PageSize)

	// If everything went OK, then return the slice of the movies and metadata.
	return cars, metadata, nil
}

func (m CarModel) Insert(car *Car) error {
	// Insert a new car item into the database.
	query := `
		INSERT INTO cars (title, description, ye_ar) 
		VALUES ($1, $2, $3) 
		RETURNING id, created_at, updated_at
		`
	args := []interface{}{car.Title, car.Description, car.Year}
	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()

	return m.DB.QueryRowContext(ctx, query, args...).Scan(&car.Id, &car.CreatedAt, &car.UpdatedAt)
}

func (m CarModel) Get(id int) (*Car, error) {
	// Return an error if the ID is less than 1.
	if id < 1 {
		return nil, ErrRecordNotFound
	}
	// Retrieve a specific car item based on its ID.
	query := `
		SELECT id, created_at, updated_at, title, description, ye_ar
		FROM cars
		WHERE id = $1
		`
	var car Car
	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()

	row := m.DB.QueryRowContext(ctx, query, id)
	err := row.Scan(&car.Id, &car.CreatedAt, &car.UpdatedAt, &car.Title, &car.Description, &car.Year)
	if err != nil {
		return nil, fmt.Errorf("cannot retrive car with id: %v, %w", id, err)
	}
	return &car, nil
}

func (m CarModel) Update(car *Car) error {
	// Update a specific menu item in the database.
	query := `
		UPDATE cars
		SET title = $1, description = $2, ye_ar = $3, updated_at = CURRENT_TIMESTAMP
		WHERE id = $4 AND updated_at = $5
		RETURNING updated_at
		`
	args := []interface{}{car.Title, car.Description, car.Year, car.Id, car.UpdatedAt}
	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()

	return m.DB.QueryRowContext(ctx, query, args...).Scan(&car.UpdatedAt)
}

func (m CarModel) Delete(id int) error {
	// Return an error if the ID is less than 1.
	if id < 1 {
		return ErrRecordNotFound
	}

	// Delete a specific cars item from the database.
	query := `
		DELETE FROM cars
		WHERE id = $1
		`
	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()

	_, err := m.DB.ExecContext(ctx, query, id)
	return err
}

func ValidateCar(v *validator.Validator, car *Car) {
	// Check if the title field is empty.
	v.Check(car.Title != "", "title", "must be provided")
	// Check if the title field is not more than 100 characters.
	v.Check(len(car.Title) <= 100, "title", "must not be more than 100 bytes long")
	// Check if the description field is not more than 1000 characters.
	v.Check(len(car.Description) <= 1000, "description", "must not be more than 1000 bytes long")
	// Check if the nutrition value is not more than 10000.
	v.Check(car.Year<= 10000, "year", "must not be more than 10000")
}