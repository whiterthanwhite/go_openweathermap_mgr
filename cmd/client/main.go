package main

import (
	"flag"
	"fmt"
	"os"

	"github.com/whiterthanwhite/go_openweathermap_mgr/internal/currentweather"
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
	cr, err := c.GetByGeoCoordinates(float32(*lat), float32(*lon))
	if err != nil {
		panic(err.Error())
	}
	fmt.Printf("Coordinates info:\ntimezone: %d\nname: %s\ncod: %d\ntemp: %f\n", cr.Timezone, cr.Name,
		cr.Cod, cr.Main.Temp)
}
