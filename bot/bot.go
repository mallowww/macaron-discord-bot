package bot

import (
	"fmt"

	"github.com/bwmarrin/discordgo"
	"github.com/mallowww/macaron-bot/config"
)

var (
	BotID      string
	discordBot *discordgo.Session
)

func Start() {
	discordBot, err := discordgo.New("Bot " + config.Token)
	if err != nil {
		fmt.Println(err.Error())
		return
	}

	user, err := discordBot.User("@me")
	if err != nil {
		fmt.Println(err.Error())
	}

	BotID = user.ID

	discordBot.AddHandler(messageHandler)
	err = discordBot.Open()
	if err != nil {
		fmt.Println(err.Error())
	}
	fmt.Sprintln("discord-bot status: running")
}

func messageHandler(session *discordgo.Session, msg *discordgo.MessageCreate) {
	if msg.Author.ID == BotID {
		return
	}

	if msg.Content == "ping" {
		_, err := session.ChannelMessageSend(msg.ChannelID, "halo~")
		fmt.Println(err.Error())
	}
}
