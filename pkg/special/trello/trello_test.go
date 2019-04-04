package trello_test

import (
	"log"
	"testing"

	"github.com/adlio/trello"
	"github.com/pelletier/go-toml"
	"github.com/stretchr/testify/assert"
)

func loadConfig() (*toml.Tree, error) {
	confPath := "conf.toml"
	t, err := toml.LoadFile(confPath)

	return t, err
}

// refer, https://github.com/adlio/trello
func TestTrello(t *testing.T) {
	assert := assert.New(t)

	tm, err := loadConfig()
	if err != nil {
		assert.Fail("load config failed")
	}

	appKey := tm.Get("trello.appKey").(string)
	token := tm.Get("trello.token").(string)
	username := tm.Get("trello.username").(string)
	boardID := tm.Get("trello.boardID").(string)
	listID := tm.Get("trello.listID").(string)

	client := trello.NewClient(appKey, token)

	m, _ := client.GetMember(username, trello.Defaults())
	boards, err := m.GetBoards(trello.Defaults())

	log.Printf("boards: %v, err: %v", len(boards), err)

	var board trello.Board
	for _, b := range boards {
		log.Printf("board: %v, number list: %v, id: %v", b.Name, len(b.Lists), b.ID)
		if b.ID == boardID {
			board = *b
		}
	}

	lists, _ := board.GetLists(trello.Defaults())

	var currentList trello.List
	for _, l := range lists {
		log.Printf("list: %v, number card: %v, id: %v", l.Name, len(l.Cards), l.ID)

		if l.ID == listID {
			currentList = *l
		}
	}

	cards, _ := currentList.GetCards(trello.Defaults())
	for _, c := range cards {
		log.Printf("card: %v, number label: %v, id: %v", c.Name, len(c.Labels), c.ID)
		if len(c.Labels) > 0 {
			log.Printf("labels: %+v", c.Labels[0])
		}
	}
}
