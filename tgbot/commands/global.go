package commands

import (
	"fmt"
	"strings"
	"time"

	"github.com/Syfaro/telegram-bot-api"
)

// GlobalCommand will handle a /global command from a chat, and give back the global scoreboard.
type GlobalCommand struct {
	bot *tgbotapi.BotAPI
}

// SetBotAPI is used to make the bot api available for the handler.
func (gc *GlobalCommand) SetBotAPI(bot *tgbotapi.BotAPI) {
	gc.bot = bot
}

// IsCommandMatch will check if the message string contains an help command.
func (gc *GlobalCommand) IsCommandMatch(update *tgbotapi.Update) bool {
	return strings.HasPrefix(strings.ToLower(update.Message.Text), "/global")
}

// PreProcessText does nothing to the message as its not used.
func (gc *GlobalCommand) PreProcessText(update *tgbotapi.Update) error {
	return nil
}

// Execute will run the help command towards the chat where the command was posted.
func (gc *GlobalCommand) Execute(update *tgbotapi.Update) error {
	hour, min, sec := time.Now().Clock()
	timestring := fmt.Sprintf("%d:%d:%d", hour, min, sec)
	msg := tgbotapi.NewMessage(update.Message.Chat.ID, timestring)
	_, err := gc.bot.SendMessage(msg)

	return err
}
