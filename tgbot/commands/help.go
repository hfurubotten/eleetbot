package commands

import (
	"strings"

	"github.com/Syfaro/telegram-bot-api"
)

// HelpCommand will handle a /help command from a chat, and print back a commandlist.
type HelpCommand struct {
	bot *tgbotapi.BotAPI
}

// SetBotAPI is used to make the bot api available for the handler.
func (hc *HelpCommand) SetBotAPI(bot *tgbotapi.BotAPI) {
	hc.bot = bot
}

// IsCommandMatch will check if the message string contains an help command.
func (hc *HelpCommand) IsCommandMatch(update *tgbotapi.Update) bool {
	return strings.HasPrefix(strings.ToLower(update.Message.Text), "/help")
}

// PreProcessText does nothing to the message as its not used.
func (hc *HelpCommand) PreProcessText(update *tgbotapi.Update) error {
	return nil
}

// Execute will run the help command towards the chat where the command was posted.
func (hc *HelpCommand) Execute(update *tgbotapi.Update) error {
	msg := tgbotapi.NewMessage(update.Message.Chat.ID, commandHelpList)
	_, err := hc.bot.SendMessage(msg)

	return err
}
