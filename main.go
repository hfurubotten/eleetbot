package main

import (
	"log"
	"os"
	"os/signal"

	"github.com/hfurubotten/eleetbot/tgbot"
)

func main() {
	// Starting bot
	bot, err := tgbot.NewTelegramBot(tgbot.EliteTimeBotToken)
	if err != nil {
		log.Fatal("Couldn't open the Telegram Bot. Error: " + err.Error())
	}

	err = bot.Start()
	if err != nil {
		log.Fatal("Couldn't open the Telegram Bot. Error: " + err.Error())
	}

	// Starting the maintenace
	//maintenance.Start(bot)

	// Prevent main from returning immediately. Wait for interrupt.
	signalChan := make(chan os.Signal, 1)
	signal.Notify(signalChan, os.Kill, os.Interrupt)
	<-signalChan
	log.Println("Application closed by user.")

}
