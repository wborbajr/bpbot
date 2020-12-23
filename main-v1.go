package main

import (
    "database/sql"
    "fmt"
    "log"

    _ "github.com/go-sql-driver/mysql"
    "github.com/go-telegram-bot-api/telegram-bot-api"

)

const (
 host     = "127.1.0.1"
 port  = "3306"
    database  = "test"
    user      = "root"
    password  = "Cpd00mjh"

    botKey      = "1492425752:AAFY7N1QUey4hSpeDQ0s6oJsZ-eJKWiY3Wo"


)

type City struct {
    Id              int
    Name            string
    Population      int
}

//Edita do modo que for melhor
type (c City) String() string {
    return c.Name
}

func checkBD(err error) {
    if err != nil {
  fmt.Println("Erro de conexção com o banco de dados!")
  fmt.Println(err)
 } else {
        fmt.Println("Sucesso em criar conexção com o banco de dados.")
 }
}

func checkBot(err error) {
 if err != nil {
  fmt.Println("Erro de conexção com o BOT")
  fmt.Println(err)
 } else {
  fmt.Println("Sucesso em carregar o BOT")
 }
}

func main() {

    // Initialize connection string.
    var connectionString = fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?allowNativePasswords=true", user, password, host, port, database)

    // Initialize connection object.
    db, err := sql.Open("mysql", connectionString)
    defer db.Close()

    err = db.Ping()
 checkBD(err)

    if err != nil {

    } else {
        var connectionBot = fmt.Sprintf("%s", botKey)
        bot, err := tgbotapi.NewBotAPI(connectionBot)

        checkBot(err)

        bot.Debug = false

        fmt.Printf("Authorized on account %s", bot.Self.UserName)

        u := tgbotapi.NewUpdate(0)
        u.Timeout = 60

        updates, err := bot.GetUpdatesChan(u)

        for update := range updates {
            if update.Message == nil {
                continue
            }

            log.Printf("[%s] %s", update.Message.From.UserName, update.Message.Text)

            if update.Message.IsCommand() {
                msg := tgbotapi.NewMessage(update.Message.Chat.ID, "")
                switch update.Message.Command() {
                case "help":
                    msg.Text = "type /sayhi or /status."
                case "sayhi":
                    msg.Text = "Hi :)"
                case "status":
                    msg.Text = "I'm ok."
                case "withArgument":
                    msg.Text = "You supplied the following argument: " + update.Message.CommandArguments()
                case "html":
                    msg.ParseMode = "html"
                    msg.Text = "This will be interpreted as HTML, click <a href=\"https://www.example.com\">here</a>"
                case "cidades":
                    res, err := db.Query("SELECT * FROM cities")
                    defer res.Close()
                    if err != nil {
                        fmt.Println("Erro ao selecionar (CITIES) no banco")
                        log.Fatal(err)
                    } else {
                        msg.Text = "Lista de cidades"
                        for res.Next() {

                            var city City
                            err := res.Scan(&city.Id, &city.Name, &city.Population)

                            if err != nil {
                                fmt.Println("Erro ao carregar a lista de cidades em (CITIES) do banco")
                                log.Fatal(err)
                            } else {

                                msg.Text = city.String()
                                fmt.Printf("%v\n", city)

                            }
                        }
                    }
                default:
                    msg.Text = "Comando não encontado"
                }
                bot.Send(msg)
            }
        }
    }
}