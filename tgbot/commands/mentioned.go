package commands

import (
	"strings"

	"github.com/go-telegram-bot-api/telegram-bot-api"
)

// MentionedCommand will handle a message where the bot is mentioned.
type MentionedCommand struct {
	bot *tgbotapi.BotAPI
}

// SetBotAPI is used to make the bot api available for the handler.
func (mc *MentionedCommand) SetBotAPI(bot *tgbotapi.BotAPI) {
	mc.bot = bot
}

// IsCommandMatch will check if the message string contains 1337.
func (mc *MentionedCommand) IsCommandMatch(update *tgbotapi.Update) bool {
	return strings.Contains(update.Message.Text, "@elitetimebot")
}

// PreProcessText will remove the actuall command text from the message.
func (mc *MentionedCommand) PreProcessText(update *tgbotapi.Update) error {
	return nil
}

// Execute will run the echo command towards the chat where the command was posted.
func (mc *MentionedCommand) Execute(update *tgbotapi.Update) error {
	msg := tgbotapi.NewMessage(update.Message.Chat.ID, mentionedText)
	_, err := mc.bot.Send(msg)

	return err
}
