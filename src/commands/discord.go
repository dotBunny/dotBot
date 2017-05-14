package commands

import (
	"github.com/bwmarrin/discordgo"
)

// Discord responds with the appropriate Discord invite link
func Discord(session *discordgo.Session, message *discordgo.MessageCreate) {

	_, _ = session.ChannelMessageSend(message.ChannelID, "Pong!")
}
