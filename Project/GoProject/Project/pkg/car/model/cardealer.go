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
	Country     string `json:"cousine"`
}

type CardealerModel struct {
	DB       *sql.DB
	InfoLog  *log.Logger
	ErrorLog *log.Logger
}

var cardealers = []Cardealer{
	{ 
		Id:      "1", 
		Title:   "Porshe Center Almaty", 
		Address: "Кульдинский тракт 12/1", 
		Country: "Germany", 
	   }, 
	   { 
		Id:      "2", 
		Title:   "Exeed", 
		Address: "Халлулина 196Б", 
		Country: "Chinese", 
	   }, 
	   { 
		Id:      "3", 
		Title:   "Cadillac Almaty", 
		Address: "Суюнбая 243/2", 
		Country: "USA", 
	   }, 
	   { 
		Id:      "4", 
		Title:   "Toyota Center Almaty", 
		Address: "Суюнбая 151", 
		Country: "Japanese", 
	   }, 
	   { 
		Id:      "5", 
		Title:   "Bentley Almaty", 
		Address: "Суюнбая 100", 
		Country: "Great Britain", 
	   },
}

func GetCardealers() []Cardealer {
	return cardealers
}

func GetCardealer(id string) (*Cardealer, error) {
	for _, r := range cardealers {
		if r.Id == id {
			return &r, nil
		}
	}
	return nil, errors.New("Cardealer not found")
}
