package main

import (
	"github.com/joho/godotenv"

	"github.com/wborbajr/bpbot/server"
	bplogger "github.com/wborbajr/bpbot/utils"
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

	bplogger.GeneralLogger.Println("Starting... "+ system_name + " " + system_version)

	server.SetupApp()

}