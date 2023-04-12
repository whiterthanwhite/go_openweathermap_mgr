package main

import (
	"context"
	"flag"
	"fmt"
	"os"
	"time"

	"github.com/whiterthanwhite/go_openweathermap_mgr/internal/currentweather"
	sp "github.com/whiterthanwhite/go_openweathermap_mgr/internal/signal_processing"
)

var (
	lat = flag.Float64("lat", 0.0, "latitude")
	lon = flag.Float64("lon", 0.0, "longtitude")
)

func main() {
	key := os.Getenv("WEATHER_API_KEY")
	if key == "" {
		panic("api key was not specified!")
	}
	flag.Parse()
	c := &currentweather.CurrentWeatherData{
		APIKey: key,
	}

	ctx, cancel := context.WithCancel(context.Background())
	go sp.Processing(cancel)
	go requestProcessing(ctx, c)

	<-ctx.Done()
	fmt.Println("exit client")
}

func requestProcessing(parentCtx context.Context, c *currentweather.CurrentWeatherData) {
	ticker := time.NewTicker(time.Second * 5)
	for {
		select {
		case <-ticker.C:
			cwr, err := c.GetByGeoCoordinates(float32(*lat), float32(*lon))
			if err != nil {
				panic(err.Error())
			}
			if err := c.SendDataToServer(cwr); err != nil {
				panic(err.Error())
			}
		case <-parentCtx.Done():
			fmt.Println("stop")
			ticker.Stop()
			return
		}
	}
}
