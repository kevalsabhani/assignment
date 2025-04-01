package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"time"
)

// Serve starts the HTTP server and listens for incoming requests.
func (app *application) serve() error {
	srv := &http.Server{
		Addr:         fmt.Sprintf(":%d", app.config.port),
		Handler:      app.routes(),
		ErrorLog:     log.New(os.Stdout, "", log.Ldate|log.Ltime|log.Lshortfile),
		IdleTimeout:  time.Minute,
		ReadTimeout:  10 * time.Second,
		WriteTimeout: 30 * time.Second,
	}

	app.logger.Printf("Starting server on %s", srv.Addr)

	err := srv.ListenAndServe()
	if err != nil {
		return fmt.Errorf("server failed to start: %w", err)
	}

	return nil
}
