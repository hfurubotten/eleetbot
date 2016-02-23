package commands

import (
	"fmt"

	"github.com/go-telegram-bot-api/telegram-bot-api"
)

// NewMemberCommand will handle a new member joining a chat.
type NewMemberCommand struct {
	bot *tgbotapi.BotAPI
}

// SetBotAPI is used to make the bot api available for the handler.
func (nmc *NewMemberCommand) SetBotAPI(bot *tgbotapi.BotAPI) {
	nmc.bot = bot
}

// IsCommandMatch will check if the new memeber field has a userID.
func (nmc *NewMemberCommand) IsCommandMatch(update *tgbotapi.Update) bool {
	return update.Message.NewChatParticipant.ID > 0
}

// PreProcessText does nothing to the message as its not used.
func (nmc *NewMemberCommand) PreProcessText(update *tgbotapi.Update) error {
	return nil
}

// Execute will run the echo command towards the chat where the command was posted.
func (nmc *NewMemberCommand) Execute(update *tgbotapi.Update) error {
	var output string
	if update.Message.NewChatParticipant.UserName == "" {
		output = fmt.Sprintf("Hey, %s. Welcome to %s.", update.Message.NewChatParticipant.FirstName, update.Message.Chat.Title)
	} else if update.Message.NewChatParticipant.UserName == "elitetimebot" {
		output = "Hey! A new chat for me to interact with, I'm excited! Type /help to see what I can do."
	} else if update.Message.NewChatParticipant.UserName == "MrHeine" {
		output = "OMG! OMG! OMG! My owner has returned to me!! \\o/ {.-.} <3"
	} else {
		output = fmt.Sprintf("Hey, @%s. Welcome to %s.", update.Message.NewChatParticipant.UserName, update.Message.Chat.Title)
	}
	msg := tgbotapi.NewMessage(update.Message.Chat.ID, output)
	_, err := nmc.bot.Send(msg)
	return err
}
