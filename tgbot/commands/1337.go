package commands

import (
	"fmt"
	"strings"
	"time"

	"github.com/go-telegram-bot-api/telegram-bot-api"
	"github.com/hfurubotten/eleetbot/tgbot/items"
)

// LeetCommand will handle a 1337 message from a chat.
type LeetCommand struct {
	bot *tgbotapi.BotAPI
}

// SetBotAPI is used to make the bot api available for the handler.
func (lc *LeetCommand) SetBotAPI(bot *tgbotapi.BotAPI) {
	lc.bot = bot
}

// IsCommandMatch will check if the message string contains 1337.
func (lc *LeetCommand) IsCommandMatch(update *tgbotapi.Update) bool {
	return strings.HasPrefix(strings.TrimSpace(update.Message.Text), "1337")
}

// PreProcessText will remove the actuall command text from the message.
func (lc *LeetCommand) PreProcessText(update *tgbotapi.Update) error {
	return nil
}

// Execute will run the echo command towards the chat where the command was posted.
func (lc *LeetCommand) Execute(update *tgbotapi.Update) error {
	chat, err := items.NewChat(update.Message.Chat)
	if err != nil {
		return err
	}

	user, err := items.NewUser(update.Message.From)
	if err != nil {
		return err
	}

	fmt.Printf("1337: User %s, TG time: %s, S time: %s \n",
		update.Message.From,
		update.Message.Time().String(),
		time.Now().String())

	return chat.Update1337(user, update.Message.Time())
}
