package commands

import "github.com/go-telegram-bot-api/telegram-bot-api"

// Command is an interface used to describe a handler for a chat command.
type Command interface {
	SetBotAPI(bot *tgbotapi.BotAPI)
	IsCommandMatch(update *tgbotapi.Update) bool
	PreProcessText(update *tgbotapi.Update) error
	Execute(update *tgbotapi.Update) error
}
