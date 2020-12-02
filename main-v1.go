package main

import (
	"log"
	"os"
	"strconv"
	"time"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
	"github.com/joho/godotenv"
)

const (
	systemName 	= "Telegram BOT"
	systemVersion = ".:: V0.0.1.1 ::."
)

var (
	superToken string
	chatID		int64
)

func init() {

	// https://api.telegram.org/bot851768239:AAGFupgI71uMHPbHPTsIvEUS3vpIyLDQXLg/deleteWebhook
	// https://api.telegram.org/bot851768239:AAGFupgI71uMHPbHPTsIvEUS3vpIyLDQXLg/getUpdates

	// 	2020/12/02 16:08:47 getUpdates resp: {"ok":true,"result":[{"update_id":151184257,
	// "message":{"message_id":92,"from":{"id":782816475,"is_bot":false,"first_name":"Waldir Borba Junior","username":"wborbajr","language_code":"en"},"chat":{"id":782816475,"first_name":"Waldir Borba Junior","username":"wborbajr","type":"private"},"date":1606936126,"text":"/start","entities":[{"offset":0,"length":6,"type":"bot_command"}]}}]}

	err := godotenv.Load()

	if err != nil {
		panic("Error loading .env file: ")
	}

	superToken = os.Getenv("BOT_TOKEN")
	chatID, _ = strconv.ParseInt(os.Getenv("BOT_CHATID"), 10, 64)

}

// TeleBot BPBot
type TeleBot struct {
	botAPI *tgbotapi.BotAPI
}

func main() {

	bot, newBotErr := tgbotapi.NewBotAPI(superToken)

	if newBotErr != nil {
		log.Println("😤😤😤 Telegram BOT Not Found")
		return
	}

	log.Printf("Authorized on account Username: %s - Appear name %s", bot.Self.UserName, bot.Self.FirstName)

	bot.Debug = true

	teleBot := TeleBot{
		botAPI: bot,
	}

	go teleBot.updateNewMsg()

	for {
		// teleBot.telegramSendMsg("now time : " + time.Now().Format(time.RFC3339))
		time.Sleep(time.Second * 30)
	}
}

func (t *TeleBot) telegramSendMsg(text string) {

	msg := tgbotapi.NewMessage(chatID, text)
	_, err := t.botAPI.Send(msg)
	if err != nil {
		log.Println("⚡︎⚡︎⚡︎ Failed to send Telegram message")
		return
	}
}

func (t *TeleBot) replyKeyboardMarkup() {

	// msg := tgbotapi.NewMessage(chatID, "Please select the action you want to perform")
	// msg.ReplyMarkup = tgbotapi.ReplyKeyboardMarkup{
	// 	Keyboard: [][]tgbotapi.KeyboardButton{
	// 		tgbotapi.NewKeyboardButtonRow(
	// 			tgbotapi.NewKeyboardButton(showMeStatus),
	// 			tgbotapi.NewKeyboardButton(adjustRTP),
	// 		),
	// 	},
	// }

	// t.botAPI.Send(msg)
}

func (t *TeleBot) updateNewMsg() {
	// bot, newBotErr := tgbotapi.NewBotAPI(superToken)
	// if newBotErr != nil {
	// 	log.Println("🐔🐔🐔 Telegram BOT 找不到")
	// 	return
	// }

	u := tgbotapi.NewUpdate(0)
	u.Timeout = 10
	updates, err := t.botAPI.GetUpdatesChan(u)

	if err != nil {
		log.Println("😰😰😰 Telegram - Update message failed")
		return
	}
	for update := range updates {

		if update.CallbackQuery != nil {
			cbd := update.CallbackQuery.Data

			switch cbd {
			case "/start":
				log.Println("data :", update.CallbackQuery.Data)
				t.telegramSendMsg("input username plz.")
			}
		}

		if update.Message != nil {

			cmd := update.Message.Command()

			switch cmd {
				case "login":
					t.showLoginKeyboard()

			}

			log.Printf("[super - %s] %s", update.Message.From.UserName, update.Message.Text)
		}

	}
}

func (t *TeleBot) showLoginKeyboard() {

	println("showLoginKeyboard")

	bt := tgbotapi.NewInlineKeyboardMarkup(
		tgbotapi.NewInlineKeyboardRow(
			tgbotapi.NewInlineKeyboardButtonData("Inser Username", "inserUsername"),
		),
	)

	msg := tgbotapi.NewMessage(chatID, "Please select action")
	msg.ReplyMarkup = bt
	t.botAPI.Send(msg)
}