package main

// imports necessary packages from the standard library
import (
	"flag"
	"fmt"
	"log"
	"net/http"
	"os"
	"time"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"

	"context"

	_ "github.com/lib/pq"
)

const version = "1.0.0"

// The config struct holds configuration values like the server's port number and environment.
type config struct {
	port int
	env  string
	db   struct {
		dsn string
	}
}

// The application struct contains the application's configuration and a logger for logging purposes.
type application struct {
	config config
	logger *log.Logger
	db     *gorm.DB
}

func main() {
	var cfg config
	flag.IntVar(&cfg.port, "port", 4000, "API server port")
	flag.StringVar(&cfg.env, "env", "development", "Environment (development|staging|production)")

	flag.StringVar(&cfg.db.dsn, "db-dsn", "postgres://postgres:postgres@localhost/tms?sslmode=disable", "PostgreSQL Connection")

	flag.Parse()
	logger := log.New(os.Stdout, "", log.Ldate|log.Ltime)

	db, err := openDB(cfg)
	if err != nil {
		logger.Fatal(err)
	}

	logger.Printf("database connection pool established")

	app := &application{
		config: cfg,
		logger: logger,
		db:     db,
	}

	// sets up an HTTP server (srv) with the specified port, the application's route handlers (returned by the app.routes() method),
	// and various timeouts for connection idle, read, and write.
	srv := &http.Server{
		Addr:         fmt.Sprintf(":%d", cfg.port),
		Handler:      app.routes(),
		IdleTimeout:  time.Minute,
		ReadTimeout:  10 * time.Second,
		WriteTimeout: 30 * time.Second,
	}
	logger.Printf("starting %s server on %s", cfg.env, srv.Addr)

	err = srv.ListenAndServe()
	logger.Fatal(err)
}

// The openDB() function returns a sql.DB connection pool.
func openDB(cfg config) (*gorm.DB, error) {

	logger := log.New(os.Stdout, "", log.Ldate|log.Ltime)

	// db, err := sql.Open("postgres", cfg.db.dsn)
	db, err := gorm.Open(postgres.Open(cfg.db.dsn), &gorm.Config{})
	if err != nil {
		return nil, err
	}

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	logger.Printf("Context %s ", ctx)
	if err != nil {
		return nil, err
	}

	return db, nil
}
