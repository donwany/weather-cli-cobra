package utils

import (
	"encoding/json"
	"github.com/joho/godotenv"
	"io/ioutil"
	"log"
	"net/http"
	"os"
)

func GetEnv(key string) string  {
	return os.Getenv(key)
}

func LoadEnv()  {
	err := godotenv.Load(".env")
	if err != nil {
		log.Fatal("Error Loading .env file")
	}
}
// Weather Struct
type WeatherResponse struct {
	Name string `json:"name"`
	Main Main `json:"main"`
}

type Main struct {
	Temp float64 `json:"temp"`
	FeelsLike float64 `json:"feels_like"`
	TempMin float64 `json:"temp_min"`
	TempMax float64 `json:"temp_max"`
	Pressure float64 `json:"pressure"`
	Humidity float64 `json:"humidity"`
}

func getData(CITY string, APIKEY string) WeatherResponse {
	URL := "https://api.openweathermap.org/data/2.5/weather?q=" + CITY + "&appid=" + APIKEY + "&units=metric"
	resBytes := makeRequest(URL)
	return resBytes
}

func makeRequest(URL string) WeatherResponse {
	response, err := http.Get(URL)
	if err != nil {
		log.Fatal(err)
	}
	// ioutil package
	jsonBytes, err := ioutil.ReadAll(response.Body)
	defer func() {
		e := response.Body.Close()
		if e != nil {
			log.Fatal(e)
		}
	}()

	var weatherResponse WeatherResponse
	er := json.Unmarshal(jsonBytes, &weatherResponse)
	if er != nil {
		log.Fatal(er)
	}
	return weatherResponse
}

func GetCity(city string) float64 {
	LoadEnv()
	APIKEY := GetEnv("APIKEY")
	response := getData(city, APIKEY)
	return response.Main.Temp
}
