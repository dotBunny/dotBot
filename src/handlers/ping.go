package handlers

import (
	"github.com/bwmarrin/discordgo"
)

func HandlePing(session *discordgo.Session, message *discordgo.MessageCreate) {
	_, _ = session.ChannelMessageSend(message.ChannelID, "Pong!")
}
