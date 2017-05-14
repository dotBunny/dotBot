package commands

import (
	"github.com/bwmarrin/discordgo"
)

// Ping responds in the channel with the standard Pong!
func Ping(session *discordgo.Session, message *discordgo.MessageCreate) {
	_, _ = session.ChannelMessageSend(message.ChannelID, "Pong!")
}
