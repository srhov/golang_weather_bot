package main

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"github.com/joho/godotenv"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	bot, err := tgbotapi.NewBotAPI(os.Getenv("TELEGRAM_TOKEN"))
	if err != nil {
		log.Panic(err)
	}

	bot.Debug = true
	log.Printf("Authorized on account %s", bot.Self.UserName)

	u := tgbotapi.NewUpdate(0)
	u.Timeout = 60

	updates := bot.GetUpdatesChan(u)

	for update := range updates {
		if update.Message == nil {
			continue
		}

		if update.Message.Location != nil {
			lat := update.Message.Location.Latitude
			lon := update.Message.Location.Longitude

			weather := getWeather(lat, lon)
			msg := tgbotapi.NewMessage(update.Message.Chat.ID, weather)
			bot.Send(msg)
		} else {
			msg := tgbotapi.NewMessage(update.Message.Chat.ID, "Please send your location 📍")
			msg.ReplyMarkup = locationRequestKeyboard()
			bot.Send(msg)
		}
	}
}

func getWeather(lat, lon float64) string {
	apiKey := os.Getenv("OPENWEATHER_API_KEY")
	url := fmt.Sprintf("https://api.openweathermap.org/data/2.5/weather?lat=%f&lon=%f&units=metric&appid=%s", lat, lon, apiKey)

	resp, err := http.Get(url)
	if err != nil {
		return "Error getting weather data 😕"
	}
	defer resp.Body.Close()

	body, _ := io.ReadAll(resp.Body)
	var data map[string]interface{}
	json.Unmarshal(body, &data)

	if data["main"] == nil {
		return "Couldn't get weather info."
	}

	main := data["main"].(map[string]interface{})
	weather := data["weather"].([]interface{})[0].(map[string]interface{})

	temp := main["temp"].(float64)
	desc := weather["description"].(string)

	return fmt.Sprintf("🌡️ %.1f°C\n🌤️ %s", temp, desc)
}

func locationRequestKeyboard() tgbotapi.ReplyKeyboardMarkup {
	button := tgbotapi.NewKeyboardButtonLocation("Send Location 📍")
	keyboard := tgbotapi.NewReplyKeyboard(
		tgbotapi.NewKeyboardButtonRow(button),
	)
	return tgbotapi.ReplyKeyboardMarkup{
		Keyboard:        keyboard.Keyboard,
		ResizeKeyboard:  true,
		OneTimeKeyboard: true,
	}
}
