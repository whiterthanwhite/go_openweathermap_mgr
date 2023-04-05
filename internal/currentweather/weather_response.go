package currentweather

type CurrentWeatherResponse struct {
	Coord struct {
		Lon float32 `json:"lon"`
		Lat float32 `json:"lat"`
	} `json:"coord"`
	Weather []struct {
		ID          int    `json:"id"`
		Main        string `json:"main"`
		Description string `json:"description"`
		Icon        string `json:"icon"`
	} `json:"weather"`
	Base string `json:"base"`
	Main struct {
		Temp      float32 `json:"temp"`
		FeelsLike float32 `json:"feels_like"`
		TempMin   float32 `json:"temp_min"`
		TempMax   float32 `json:"temp_max"`
		Pressure  int     `json:"pressure"`
		Humidity  int     `json:"humidity"`
		SeaLevel  int     `json:"sea_level,omitempty"`
		GrndLevel int     `json:"grnd_level,omitempty"`
	} `json:"main"`
	Visibility int `json:"visibility"`
	Wind       struct {
		Speed float32 `json:"speed"`
		Deg   int     `json:"deg"`
		Gust  float32 `json:"gust,omitempty"`
	} `json:"wind"`
	Rain struct {
		Hour1 float32 `json:"1h,omitempty"`
		Hour3 float32 `json:"3h,omitempty"`
	} `json:"rain,omitempty"`
	Snow struct {
		Hour1 float32 `json:"1h,omitempty"`
		Hour3 float32 `json:"3h,omitempty"`
	} `json:"snow,omitempty"`
	Clouds struct {
		All int `json:"all"`
	} `json:"clouds"`
	Dt  int `json:"dt"`
	Sys struct {
		Type    int    `json:"type"`
		ID      int    `json:"id"`
		Country string `json:"country"`
		Sunrise int    `json:"sunrise"`
		Sunset  int    `json:"sunset"`
	} `json:"sys"`
	Timezone int    `json:"timezone"`
	ID       int    `json:"id"`
	Name     string `json:"name"`
	Cod      int    `json:"cod"`
}
