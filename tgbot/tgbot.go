package tgbot

import (
	"log"
	"time"

	"github.com/go-telegram-bot-api/telegram-bot-api"
	"github.com/hfurubotten/eleetbot/tgbot/commands"
	"github.com/hfurubotten/eleetbot/tgbot/items"
)

// TelegramBot is the main TG chat object for a bot.
type TelegramBot struct {
	token string
	Bot   *tgbotapi.BotAPI
}

// NewTelegramBot creates a new TelegramBot bot object.
func NewTelegramBot(bottoken string) (*TelegramBot, error) {
	tgbot := new(TelegramBot)

	tgbot.token = bottoken

	bot, err := tgbotapi.NewBotAPI(bottoken)
	if err != nil {
		return nil, err
	}

	tgbot.Bot = bot
	bot.Debug = false

	log.Printf("Authorized on account %s", bot.Self.UserName)
	return tgbot, nil
}

// Start will create the connection to TG and start a new tread for processing
// chat messages.
func (tg *TelegramBot) Start() error {
	u := tgbotapi.NewUpdate(0)
	u.Timeout = 60

	updates, err := tg.Bot.GetUpdatesChan(u)
	if err != nil {
		return err
	}

	go func() {
		//defer func() {
		//	if r := recover(); r != nil {
		//		log.Println("Recovered from panic: ", r)
		//		tg.Start()
		//	}
		//}()

		for _, handler := range commands.StandardCommandHandlers {
			handler.SetBotAPI(tg.Bot)
		}

		for update := range updates {
			log.Printf("[%s] in %s: %s", update.Message.From.UserName, update.Message.Chat.Title, update.Message.Text)

			chat, err := items.NewChat(update.Message.Chat)
			if err != nil {
				log.Println(err)
			}

			chat.LastSeen = time.Now()

			for _, handler := range commands.StandardCommandHandlers {
				if handler.IsCommandMatch(&update) {
					err = handler.PreProcessText(&update)
					if err != nil {
						log.Println("Command not executed due to failed pre processing. \nError:",
							err,
							"\nCommand text:",
							update.Message.Text)
						continue
					}

					err = handler.Execute(&update)
					if err != nil {
						log.Println("Failed to execute command handler. \nError:",
							err,
							"\nCommand text: ",
							update.Message.Text)
					}

					break
				}
			}

		}
	}()

	log.Println("Telegram bot finished starting up.")
	return nil
}
