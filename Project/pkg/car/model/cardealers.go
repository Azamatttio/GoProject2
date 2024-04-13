package model

import (
	"database/sql"
	"errors"
	"log"
)

type Cardealer struct {
	Id          string `json:"id"`
	CreatedAt   string `json:"createdAt"`
	UpdatedAt   string `json:"updatedAt"`
	Title       string `json:"title"`
	Description string `json:"description"`
	Address     string `json:"address"`
	Coordinates string `json:"coordinates"`
	Country     string `json:"country"`
}

type CardealerModel struct {
	DB       *sql.DB
	InfoLog  *log.Logger
	ErrorLog *log.Logger
}

func (m CardealerModel) GetAll() ([]*Cardealer, error) {
	// TODO: implement this method
	return nil, errors.New("not implemented")
}

func (m CardealerModel) Insert(car *Cardealer) error {
	// TODO: implement this method
	return errors.New("not implemented")
}

func (m CardealerModel) Get(id int) (*Cardealer, error) {
	// TODO: implement this method
	return nil, errors.New("not implemented")
}

func (m CardealerModel) Update(car *Cardealer) error {
	// TODO: implement this method
	return errors.New("not implemented")
}

func (m CardealerModel) Delete(id int) error {
	// TODO: implement this method
	return errors.New("not implemented")
}