package weather

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

const apiKey = "46da0e08ad7c0b06cd00d5ddfb588070"

type WeatherStruct struct {
	Name string `json:"name"`
	Main struct {
		Temp float64 `json:"temp"`
	} `json:"main"`
	Weather []struct {
		Description string `json:"description"`
	} `json:"weather"`
}

type City struct {
	Name string `json:"name"`
}

func Weather(c string) string {
	var city City
	err := json.Unmarshal([]byte(c), &city)
	if err != nil {
		return "Введите корректные данные"
	}
	url := fmt.Sprintf("http://api.openweathermap.org/data/2.5/weather?q=%s&units=metric&lang=ru&appid=%s", city.Name, apiKey)
	resp, err := http.Get(url)
	if err != nil {
		return "Не удается получить данные о погоде"
	}

	defer resp.Body.Close()
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return "Ошибка при обработке погодных данных"
	}
	var w WeatherStruct
	err = json.Unmarshal(body, &w)
	if err != nil {
		return "Не найдено данных о погоде для указанного города"
	}
	if len(w.Weather) == 0 {
		return "Не найдено данных о погоде для указанного города"
	}
	temp := fmt.Sprintf("%.0f", w.Main.Temp)
	WeatherCity := w.Name + " - " + w.Weather[0].Description + " " + temp + " градусов"
	return WeatherCity
}
