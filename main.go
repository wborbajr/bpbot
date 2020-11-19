package main

import (
	"log"

	"github.com/joho/godotenv"

	"github.com/wborbajr/bpbot/server"
)

const (
	system_name 	= "Telegram BOT"
	system_version 	= ".:: V0.0.1.1 ::."
)

func init() {
	err := godotenv.Load()

	if err != nil {
		panic("Error loading .env file: ")
	}
}

func main() {

	// Configure Logging
	// LOG_FILE_LOCATION := os.Getenv("LOG_FILE_LOCATION")
	// if LOG_FILE_LOCATION != "" {
	// 	log.SetOutput(&lumberjack.Logger{
	// 		Filename:   LOG_FILE_LOCATION,
	// 		MaxSize:    500, // megabytes
	// 		MaxBackups: 3,
	// 		MaxAge:     28,   //days
	// 		Compress:   true, // disabled by default
	// 	})
	// }

	log.Println("Starting... "+ system_name + " " + system_version)

	server.SetupApp()

}