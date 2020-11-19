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
	log.Println("Starting... "+ system_name + " " + system_version)

	server.SetupApp()

}