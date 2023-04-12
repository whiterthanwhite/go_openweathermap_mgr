package database

import (
	"context"
	"testing"
	"time"

	"github.com/whiterthanwhite/go_openweathermap_mgr/internal/currentweather"
)

func TestInsertWeatherData(t *testing.T) {
	ctx, _ := context.WithTimeout(context.Background(), time.Second*3)
	db := GetInstance()
	if err := db.CreateConnection(ctx); err != nil {
		t.Fatal(err.Error())
	}
	weather := &currentweather.Weather{
		Lon:  10.0,
		Lat:  20.0,
		Temp: 18.5,
	}
	if err := db.InsertWeatherData(ctx, weather); err != nil {
		t.Fatal(err.Error())
	}
	<-ctx.Done()
}

func TestGetLastWeatherId(t *testing.T) {
	ctx, _ := context.WithTimeout(context.Background(), time.Second*3)
	db := GetInstance()
	if err := db.CreateConnection(ctx); err != nil {
		t.Fatal(err.Error())
	}
	id, err := db.GetLastWeatherId(ctx)
	if err != nil {
		t.Fatal(err.Error())
	}
	if id != 0 {
		t.Logf("Actual: %d; Expected: %d", id, 0)
		t.Fail()
	}
	<-ctx.Done()
}
