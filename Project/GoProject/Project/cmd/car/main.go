package main

import (
	"database/sql"
	"flag"
	"github.com/codev0/inft3212-6/pkg/abr-plus/model"
	"github.com/gorilla/mux"
	"log"
	"net/http"

	_ "github.com/lib/pq"
)

type config struct {
	port string
	env  string
	db   struct {
		dsn string
	}
}

type application struct {
	config config
	models model.Models
}

func main() {
	var cfg config
	flag.StringVar(&cfg.port, "port", ":8081", "API server port")
	flag.StringVar(&cfg.env, "env", "development", "Environment (development|staging|production)")
	flag.StringVar(&cfg.db.dsn, "db-dsn", "postgresql://postgres:Almira_24@localhost/mydatabase?sslmode=disable", "PostgreSQL DSN")
	flag.Parse()

	// Connect to DB
	db, err := openDB(cfg)
	if err != nil {
		log.Fatal(err)
		return
	}
	defer db.Close()

	app := &application{
		config: cfg,
		models: model.NewModels(db),
	}

	app.run()
}

func (app *application) run() {
	r := mux.NewRouter()

	v1 := r.PathPrefix("/api/v1").Subrouter()

	// Car Singleton
	// Create a new car
	v1.HandleFunc("/cars", app.createCarHandler).Methods("POST")
	// Get a specific car
	v1.HandleFunc("/cars/{carId:[0-9]+}", app.getCarHandler).Methods("GET")
	// Update a specific car
	v1.HandleFunc("/cars/{carId:[0-9]+}", app.updateCarHandler).Methods("PUT")
	// Delete a specific car
	v1.HandleFunc("/cars/{carId:[0-9]+}", app.deleteCarHandler).Methods("DELETE")

	log.Printf("Starting server on %s\n", app.config.port)
	err := http.ListenAndServe(app.config.port, r)
	log.Fatal(err)
}

func openDB(cfg config) (*sql.DB, error) {
	// Use sql.Open() to create an empty connection pool, using the DSN from the config // struct.
	db, err := sql.Open("postgres", cfg.db.dsn)
	if err != nil {
		return nil, err
	}
	return db, nil
}
