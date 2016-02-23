package commands

import (
	"fmt"
	"strings"

	"github.com/go-telegram-bot-api/telegram-bot-api"
)

// EchoCommand will handle a /echo command from a chat. It process the incomming
// message and searches for certain triggers. These triggers will give another
// response. Standard responce is to send back the same message.
type EchoCommand struct {
	bot *tgbotapi.BotAPI
}

// SetBotAPI is used to make the bot api available for the handler.
func (ec *EchoCommand) SetBotAPI(bot *tgbotapi.BotAPI) {
	ec.bot = bot
}

// IsCommandMatch will check if the message string contains an echo command.
func (ec *EchoCommand) IsCommandMatch(update *tgbotapi.Update) bool {
	return strings.HasPrefix(strings.ToLower(update.Message.Text), "/echo")
}

// PreProcessText will remove the actuall command text from the message.
func (ec *EchoCommand) PreProcessText(update *tgbotapi.Update) error {
	update.Message.Text = strings.Replace(update.Message.Text, "/echo ", "", 1)
	update.Message.Text = strings.Replace(update.Message.Text, "/Echo ", "", 1)
	update.Message.Text = strings.Replace(update.Message.Text, "/echo@elitetimebot ", "", 1)
	update.Message.Text = strings.Replace(update.Message.Text, "/Echo@elitetimebot ", "", 1)

	return nil
}

// Execute will run the echo command towards the chat where the command was posted.
func (ec *EchoCommand) Execute(update *tgbotapi.Update) error {
	echomsg := update.Message.Text

	if (strings.Contains(strings.ToLower(echomsg), "i am") ||
		strings.Contains(strings.ToLower(echomsg), "im") ||
		strings.Contains(strings.ToLower(echomsg), "i'm")) &&
		(strings.Contains(strings.ToLower(echomsg), "noob") ||
			strings.Contains(strings.ToLower(echomsg), "weiner") ||
			strings.Contains(strings.ToLower(echomsg), "wiener") ||
			strings.Contains(strings.ToLower(echomsg), "dumb") ||
			strings.Contains(strings.ToLower(echomsg), "stupid")) {

		echomsg = fmt.Sprintf("Thats really not a good thing to say about yourself, @%s.", update.Message.From.UserName)
		msg := tgbotapi.NewMessage(update.Message.Chat.ID, echomsg)
		msg.ReplyToMessageID = update.Message.MessageID
		_, err := ec.bot.Send(msg)
		return err
	}

	msg := tgbotapi.NewMessage(update.Message.Chat.ID, echomsg)
	_, err := ec.bot.Send(msg)

	return err
}
