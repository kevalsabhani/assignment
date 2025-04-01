package main

import (
	"flag"
	"log"
	"os"

	"github.com/kevalsabhani/assignment/internal/data"
)

type config struct {
	port int
	env  string
}

type application struct {
	config config
	logger *log.Logger
	models data.Models
}

func main() {
	var cfg config

	flag.IntVar(&cfg.port, "port", 4000, "API server port")
	flag.StringVar(&cfg.env, "env", "development", "Environment (development|staging|production)")

	flag.Parse()

	logger := log.New(os.Stdout, "", log.Ldate|log.Ltime|log.Lshortfile)

	app := &application{
		config: cfg,
		logger: logger,
		models: data.NewModels(),
	}

	err := app.serve()
	if err != nil {
		logger.Fatal(err)
	}
}
