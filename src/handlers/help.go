package handlers

import (
	"github.com/bwmarrin/discordgo"
)

// HandleHelp flags a staff members attention that someone needs help
func HandleHelp(session *discordgo.Session, message *discordgo.MessageCreate) {
	_, _ = session.ChannelMessageSend(message.Author.ID, "test")
}
