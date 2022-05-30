package bot

import (
	"fmt"
	"github.com/bwmarrin/discordgo"
)

var botid string
var gobot *discordgo.Session


func Start() {
	token := "yo token"
	gobot, err := discordgo.New("Bot " + token)
	if err != nil {
		panic(err)
		return
	}
	u, err := gobot.User("@me")
	if err != nil {
		panic(err)
		return
	}
	botid = u.ID
	gobot.AddHandler(messageHandler)
	err = gobot.Open()
	if err != nil {
		panic(err)
		return
	}
	fmt.Println("Luna is up and running")
}

func messageHandler(s *discordgo.Session, m*discordgo.MessageCreate) {
	if m.Author.ID == botid {
		return
	}

	if m.Content == "!ping" {
		_,_ = s.ChannelMessageSend(m.ChannelID, "Pong!")
	}
}