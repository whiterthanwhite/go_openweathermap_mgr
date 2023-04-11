package main

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"os"
	"time"

	"github.com/whiterthanwhite/go_openweathermap_mgr/internal/database"
	"github.com/whiterthanwhite/go_openweathermap_mgr/internal/middleware"
	"github.com/whiterthanwhite/go_openweathermap_mgr/internal/serverhandlers"
	sp "github.com/whiterthanwhite/go_openweathermap_mgr/internal/signal_processing"
)

func main() {
	fmt.Printf("server start %v\n", time.Now())

	ctx, cancel := context.WithCancel(context.Background())
	go sp.Processing(cancel)

	mux := http.NewServeMux()
	mux.HandleFunc("/currentweather", serverhandlers.AddWeatherMeasurement())
	s := http.Server{
		Addr:    ":8080",
		Handler: &middleware.Logger{Handler: mux, LogWriter: os.Stdout},
	}

	// Process database connection
	if err := databaseProcessing(ctx); err != nil {
		cancel()
		log.Fatal(err.Error())
	}

	go func() {
		select {
		case <-ctx.Done():
			if err := database.GetInstance().Close(ctx); err != nil {
				log.Println(err.Error())
			}
			s.Close()
		}
	}()

	if err := s.ListenAndServe(); err != nil {
		log.Println(err.Error())
	}
	<-ctx.Done()
	fmt.Printf("server stop %v\n", time.Now())
}

func databaseProcessing(parentCtx context.Context) error {
	dbinstance := database.GetInstance()
	if err := dbinstance.CreateConnection(parentCtx); err != nil {
		return err
	}
	if err := dbinstance.CreateNecessaryTables(parentCtx); err != nil {
		return err
	}
	return nil
}
