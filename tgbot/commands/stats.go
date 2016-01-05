package commands

import (
	"fmt"
	"log"
	"strings"
	"time"

	"github.com/Syfaro/telegram-bot-api"
	"github.com/hfurubotten/eleetbot/tgbot/items"
)

const (
	Day   = iota
	Week  = iota
	Month = iota
)

// StatsCommand will handle a /stats command from a chat. It process the incomming
// message and searches for certain triggers, and send back the appropiate stast.
type StatsCommand struct {
	bot       *tgbotapi.BotAPI
	fullstats bool
	period    int
}

// SetBotAPI is used to make the bot api available for the handler.
func (sc *StatsCommand) SetBotAPI(bot *tgbotapi.BotAPI) {
	sc.bot = bot
}

// IsCommandMatch will check if the message string contains 1337.
func (sc *StatsCommand) IsCommandMatch(update *tgbotapi.Update) bool {
	return strings.HasPrefix(strings.ToLower(update.Message.Text), "/stats")
}

// PreProcessText will remove the actuall command text from the message.
func (sc *StatsCommand) PreProcessText(update *tgbotapi.Update) error {
	sc.fullstats = strings.Contains(strings.ToLower(update.Message.Text), "full") ||
		strings.Contains(strings.ToLower(update.Message.Text), "all")

	if strings.Contains(strings.ToLower(update.Message.Text), "week") {
		sc.period = Week
	} else if strings.Contains(strings.ToLower(update.Message.Text), "month") {
		sc.period = Month
	} else {
		sc.period = Day
	}

	return nil
}

// Execute will run the echo command towards the chat where the command was posted.
func (sc *StatsCommand) Execute(update *tgbotapi.Update) error {
	chat, err := items.NewChat(update.Message.Chat)
	if err != nil {
		return err
	}

	var top []int
	var scoreboard map[int]int64
	out := ""
	if sc.period == Week {
		top = chat.TopPlayersWeek()
		scoreboard = chat.ScoreBoardWeek
		out = "Elite timemasters this week"
	} else if sc.period == Month {
		top = chat.TopPlayersMonth()
		scoreboard = chat.ScoreBoardMonth
		out = "Elite timemasters this month"
	} else {
		out = "Elite timemasters today"
		if sc.fullstats {
			top = chat.TopPlayers()
			scoreboard = chat.ScoreBoard
		} else {
			top = chat.Top5Players()
			scoreboard = chat.ScoreBoard
		}
	}

	if len(top) == 0 {
		msg := tgbotapi.NewMessage(update.Message.Chat.ID, "No one have been elite enough...")
		_, err = sc.bot.SendMessage(msg)

		return err
	}

	for i := 0; i < len(top); i++ {
		user, err := items.NewUser(tgbotapi.User{
			ID: top[i],
		})
		if err != nil {
			log.Println(err)
			continue
		}

		score := time.Duration(scoreboard[user.ID])
		if score.Seconds() < 1 {
			out += fmt.Sprintf("\n%s: %d ms", user.UserName, (score.Nanoseconds() / 1000))
		} else {
			out += fmt.Sprintf("\n%s: %s", user.UserName, score)
		}
	}

	msg := tgbotapi.NewMessage(update.Message.Chat.ID, out)
	_, err = sc.bot.SendMessage(msg)

	return err
}
