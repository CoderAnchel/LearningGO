package main

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"os"
	"time"
)

type Weather struct {
	Location struct {
		Name    string `"json:name"`
		Country string `"json:country"`
	} `"json:location"`
	Current struct {
		Temp_C    float64 `"json:temp_C"`
		Condition struct {
			Text string `"json:text"`
		} `"json:condition"`
	} `"json:current"`
	Forecast struct {
		Forecastday []struct {
			Hour []struct {
				Time_epoch int64   `"json:time_epoch"`
				Temp_C     float64 `"json:temp_c"`
				Condition  struct {
					Text string `"json:text"`
				} `"json:condition"`
				ChanceOfRain float64 `"json:chance_of_rain"`
			} `"json:hour"`
		} `"json:forecastday"`
	} `"json:forecast"`
}

func getPrevisionDay(day string) int {
	switch day {
	case "today":
		return 0
	case "tomorrow":
		return 1
	}
	return 0
}

func main() {
	var city string = "Barcelona"
	var pDay string

	if len(os.Args) >= 3 {
		city = os.Args[1]
		pDay = os.Args[2]
	}

	res, err := http.Get("http://api.weatherapi.com/v1/forecast.json?key=f49c421e8ec14d70927141947241311&q=" + city + "&days=5&aqi=no&alerts=no")

	if err != nil {
		panic(err)
	}

	defer res.Body.Close()

	if res.StatusCode != 200 {
		fmt.Println("Error with the data fetching")
	}

	body, err := io.ReadAll(res.Body)
	if err != nil {
		panic(err)
	}

	var clima Weather
	err = json.Unmarshal(body, &clima)
	if err != nil {
		panic(err)
	}

	location, current, hours := clima.Location, clima.Current, clima.Forecast.Forecastday[getPrevisionDay(pDay)].Hour

	fmt.Printf("%s, %s: %.0f, %s \n", location.Name, location.Country, current.Temp_C, current.Condition.Text)

	for _, hour := range hours {
		date := time.Unix(hour.Time_epoch, 0)

		if date.Before(time.Now()) {
			continue
		}

		fmt.Printf(
			"%s - %.0fC, %.0f%%,  %s \n",
			date.Format("15:04"),
			hour.Temp_C,
			hour.ChanceOfRain,
			hour.Condition.Text,
		)

	}
}
