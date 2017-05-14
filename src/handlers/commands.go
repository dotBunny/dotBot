package handlers

import (
	"github.com/bwmarrin/discordgo"
)

// HandleCommands responds to a user privately with information about the bots commands
func HandleCommands(session *discordgo.Session, message *discordgo.MessageCreate) {

	iconPath := "https://dl.dropboxusercontent.com/u/118962/dotBunny/dotBot/icon.png"

	_, _ = session.ChannelMessageSendEmbed(message.ChannelID, &discordgo.MessageEmbed{
		Type:        "rich",
		Title:       "dotBot Commands",
		Description: "All commands respond in the channel issued in and can be sent directly.",
		Color:       1974564,
		Fields: []*discordgo.MessageEmbedField{
			&discordgo.MessageEmbedField{
				Name:   ".meetup",
				Value:  "Responds back with information on the next PTBO Game Dev Meetup",
				Inline: false},
			&discordgo.MessageEmbedField{
				Name:   ".ping",
				Value:  "Responds back with the expected \"Pong!\"",
				Inline: false}},
		Thumbnail: &discordgo.MessageEmbedThumbnail{
			URL: iconPath},
	})
}
