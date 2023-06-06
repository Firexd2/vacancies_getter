package core

import (
	"github.com/gotd/td/tg"
	"strings"
	"vacancies_getter/structs"
)

const (
	channelID = 1246902558 // cyprus it
	//channelID = 1454499729 // test
)

var patterns = []string{"Python", "python"}

func containsAny(s string, substrings []string) bool {
	for _, v := range substrings {
		if strings.Contains(s, v) {
			return true
		}
	}
	return false
}

func NewMessageHandler(new_vanacies_ch chan structs.NewVacancy, update *tg.UpdateNewChannelMessage) {
	switch message := update.Message.(type) {
	case *tg.Message:
		switch peer := message.PeerID.(type) {
		case *tg.PeerChannel:
			if peer.ChannelID == channelID {
				if !containsAny(message.Message, patterns) {
					return
				}
				//fmt.Println(message)
				new_vanacies_ch <- structs.NewVacancy{message.Message, message.ID}
			}
		}
	}
}
