# Weather Telegram Bot

A simple Telegram bot written in Go that provides current weather information based on the user's location. The bot uses the [OpenWeatherMap API](https://openweathermap.org/api) and the [go-telegram-bot-api](https://github.com/go-telegram-bot-api/telegram-bot-api) library. Built with [Cobra](https://github.com/spf13/cobra) for CLI structure.

## Features
- Users send their location to the bot and receive current weather data (temperature and description).
- If a user sends a message without a location, the bot prompts for location sharing.
- Uses environment variables for API keys.

## Requirements
- Go 1.18+
- Telegram Bot Token ([How to get one](https://core.telegram.org/bots#6-botfather))
- OpenWeatherMap API Key ([Get one here](https://home.openweathermap.org/api_keys))

## Setup
1. **Clone the repository:**
   ```sh
   git clone <your-repo-url>
   cd tgbot
   ```

2. **Install dependencies:**
   ```sh
   go mod tidy
   ```

3. **Create a `.env` file** in the project root with the following content:
   ```env
   TELE_TOKEN=your-telegram-bot-token
   OPENWEATHER_API_KEY=your-openweather-api-key
   ```

4. **Build and run the bot:**
   ```sh
   go run main.go start
   ```
   Or build the binary:
   ```sh
   go build -o weather-bot
   ./weather-bot start
   ```

## Usage
- Start a chat with your bot on Telegram.
- Send your location using the "Send Location 📍" button.
- Receive the current weather for your location.

## CLI Commands
- `start` — Starts the Telegram bot.
- `version` — Shows the application version.

## Project Structure
- `main.go` — Entry point.
- `cmd/` — Cobra commands (`handler.go`, `root.go`, `version.go`).

## License
See [LICENSE](LICENSE).
https://t.me/golang_weather_7d38a6ca_bot
---
*Made for DEVOPS101 Week 2*
