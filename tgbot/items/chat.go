package items

import (
	"sync"
	"time"

	"github.com/go-telegram-bot-api/telegram-bot-api"
)

var chats = make(map[int]*Chat)

// Chat reperesent a chat in TG for the server, and stores relevant information about it.
type Chat struct {
	tgbotapi.Chat
	Token    string
	LastSeen time.Time

	ScoreBoard             map[int]int64
	scoreBoardTackingDay   int
	ScoreBoardWeek         map[int]int64
	scoreBoardTackingWeek  int
	ScoreBoardMonth        map[int]int64
	scoreBoardTackingMonth time.Month
	scoreBoardLock         *sync.Mutex
}

// NewChat will create a new chat object from an api data object.
func NewChat(apichat tgbotapi.Chat) (*Chat, error) {
	if chat, ok := chats[apichat.ID]; ok {
		return chat, nil
	}

	val := &Chat{
		Chat:           apichat,
		Token:          "",
		ScoreBoard:     make(map[int]int64),
		scoreBoardLock: &sync.Mutex{},
	}
	chats[apichat.ID] = val

	return val, nil
}

// Update1337 will register a 1337 event for the specific chat.
func (c *Chat) Update1337(user *User, timestamp time.Time) error {
	c.scoreBoardLock.Lock()
	defer c.scoreBoardLock.Unlock()

	c.checkScoreboardReset()

	now := time.Now()

	leet := time.Date(now.Year(), now.Month(), now.Day(), 13, 37, 0, 0, now.Location())

	if leet.After(timestamp) {
		leet = time.Date(now.Year(), now.Month(), now.Day(), 13, 37, 0, 0, now.Location()).AddDate(0, 0, -1)
	}

	score := int64(timestamp.Sub(leet))

	if storedscore, ok := c.ScoreBoard[user.ID]; ok {
		if storedscore > score {
			c.ScoreBoard[user.ID] = score
		}
	} else {
		c.ScoreBoard[user.ID] = score
	}

	if storedscore, ok := c.ScoreBoardWeek[user.ID]; ok {
		if storedscore > score {
			c.ScoreBoardWeek[user.ID] = score
		}
	} else {
		c.ScoreBoardWeek[user.ID] = score
	}

	if storedscore, ok := c.ScoreBoardMonth[user.ID]; ok {
		if storedscore > score {
			c.ScoreBoardMonth[user.ID] = score
		}
	} else {
		c.ScoreBoardMonth[user.ID] = score
	}

	return nil
}

// Top5Players will give the top five players within the current day.
func (c *Chat) Top5Players() []int {
	c.scoreBoardLock.Lock()
	defer c.scoreBoardLock.Unlock()

	sorter := NewScoreSorter(c.ScoreBoard)
	players := sorter.Sorted()

	out := make([]int, 0)
	for i := 0; i < len(players) && i < 5; i++ {
		out = append(out, players[i])
	}

	return out
}

// TopPlayersWeek will give the top players within the current week.
func (c *Chat) TopPlayersWeek() []int {
	c.scoreBoardLock.Lock()
	defer c.scoreBoardLock.Unlock()

	sorter := NewScoreSorter(c.ScoreBoardWeek)
	return sorter.Sorted()

	// top 5 implementation
	//players := sorter.Sorted()

	//out := make([]int, 0)
	//for i := 0; i < len(players) && i < 5; i++ {
	//	out = append(out, players[i])
	//}

	//return out
}

// TopPlayersMonth will give the top players within the current month.
func (c *Chat) TopPlayersMonth() []int {
	c.scoreBoardLock.Lock()
	defer c.scoreBoardLock.Unlock()

	sorter := NewScoreSorter(c.ScoreBoardMonth)
	return sorter.Sorted()

	// top 5 implementation
	//players := sorter.Sorted()

	//out := make([]int, 0)
	//for i := 0; i < len(players) && i < 5; i++ {
	//	out = append(out, players[i])
	//}

	//return out
}

// TopPlayers will give a sorted list of all top players from this chat.
func (c *Chat) TopPlayers() []int {
	c.scoreBoardLock.Lock()
	defer c.scoreBoardLock.Unlock()

	sorter := NewScoreSorter(c.ScoreBoard)
	return sorter.Sorted()
}

func (c *Chat) checkScoreboardReset() {
	if c.scoreBoardTackingDay != time.Now().Day() {
		c.scoreBoardTackingDay = time.Now().Day()
		c.ScoreBoard = make(map[int]int64)
	}

	_, week := time.Now().ISOWeek()
	if c.scoreBoardTackingWeek != week {
		c.scoreBoardTackingWeek = week
		c.ScoreBoardWeek = make(map[int]int64)
	}

	if c.scoreBoardTackingMonth != time.Now().Month() {
		c.scoreBoardTackingMonth = time.Now().Month()
		c.ScoreBoardMonth = make(map[int]int64)
	}
}

// ForEachChat will run a function for each chat registered.
func ForEachChat(handler func(key int, chat *Chat) error) (err error) {
	for id, val := range chats {
		err = handler(id, val)
		if err != nil {
			return err
		}
	}
	return nil
}
