package commands

import (
	Helpers "../helpers"

	"github.com/bwmarrin/discordgo"
)

// Discord responds with the appropriate Discord invite link
func Discord(session *discordgo.Session, message *discordgo.MessageCreate) {

	messageChannel, foundError := session.Channel(message.ChannelID)

	if foundError == nil {
		if messageChannel.GuildID == Helpers.InternalGuildID() {
			_, _ = session.ChannelMessageSendEmbed(message.ChannelID, &discordgo.MessageEmbed{
				Type:        "rich",
				Title:       "dotBunny Discord Server",
				URL:         Helpers.InternalDiscordURI(),
				Description: "Join the dotBunny Discord server.",
				Color:       3394815,
				Thumbnail: &discordgo.MessageEmbedThumbnail{
					URL: Helpers.InternalIconURI()},
			})
		} else {

			_, _ = session.ChannelMessageSendEmbed(message.ChannelID, &discordgo.MessageEmbed{
				Type:        "rich",
				Title:       "PTBO Game Jam Discord Server",
				URL:         Helpers.GameJamDiscordURI(),
				Description: "Join the cool local #gamedev's and all the other awesome people.",
				Color:       52479,
				Thumbnail: &discordgo.MessageEmbedThumbnail{
					URL: Helpers.GameJamIconURI()},
			})
		}
	}
}
