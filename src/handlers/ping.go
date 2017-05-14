package handlers

import (
	"github.com/bwmarrin/discordgo"
)

// HandlePing responds in the channel with the standard Pong!
func HandlePing(session *discordgo.Session, message *discordgo.MessageCreate) {
	_, _ = session.ChannelMessageSend(message.ChannelID, "Pong!")
}
