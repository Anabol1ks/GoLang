package main

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"

	"github.com/gin-gonic/gin"
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

func weather(c *gin.Context) {
	var city City
	if err := c.BindJSON(&city); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Неправильный формат ввода"})
		return
	}
	url := fmt.Sprintf("http://api.openweathermap.org/data/2.5/weather?q=%s&units=metric&lang=ru&appid=%s", city.Name, apiKey)
	resp, err := http.Get(url)
	if err != nil {
		c.JSON(http.StatusServiceUnavailable, gin.H{"error": "Не удается получить данные о погоде"})
		return
	}

	defer resp.Body.Close()
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Ошибка при обработке погодных данных"})
		return
	}
	var w WeatherStruct
	err = json.Unmarshal(body, &w)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Не найдено данных о погоде для указанного города"})
		return
	}
	if len(w.Weather) == 0 {
		c.JSON(http.StatusNotFound, gin.H{"error": "Не найдено данных о погоде для указанного города"})
		return
	}
	temp := fmt.Sprintf("%.0f", w.Main.Temp)
	WeatherCity := w.Name + " - " + w.Weather[0].Description + " " + temp + " градусов"
	c.JSON(http.StatusOK, WeatherCity)
}

func main() {
	router := gin.Default()
	router.POST("/weather", weather)
	router.Run(":2002")
}
