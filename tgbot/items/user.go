package items

import (
	"github.com/go-telegram-bot-api/telegram-bot-api"
)

var users = make(map[int]*User)

// User reperesent a user in TG, and stores the relevant information about the user.
type User struct {
	tgbotapi.User
}

// NewUser will create a new User object.
func NewUser(apiUser tgbotapi.User) (*User, error) {
	if chat, ok := users[apiUser.ID]; ok {
		return chat, nil
	}

	val := &User{
		User: apiUser,
	}
	users[apiUser.ID] = val

	return val, nil
}
