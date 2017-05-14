package handlers

import (
	Helpers "../helpers"
	"github.com/bwmarrin/discordgo"
	"github.com/mmcdole/gofeed"
)

// HandleMeetup responds in the channel with the next PTBO Game Dev Meetup information
func HandleMeetup(session *discordgo.Session, message *discordgo.MessageCreate) {

	fp := gofeed.NewParser()
	feed, _ := fp.ParseURL("https://www.meetup.com/PTBOGameDev/events/rss/")

	foundImage, imageElement := Helpers.GetElement("img", feed.Items[0].Description, 0)
	foundTimeBlock, timeElement := Helpers.GetElementContent("p", feed.Items[0].Description, 4)

	imageSource := "https://secure.meetupstatic.com/photos/event/1/1/2/9/event_457804393.jpeg"
	nextEvent := "Unknown"

	if foundTimeBlock {
		nextEvent = timeElement
	}

	if foundImage {
		_, imageSource = Helpers.GetAttribute("src", imageElement)
	}

	_, _ = session.ChannelMessageSendEmbed(message.ChannelID, &discordgo.MessageEmbed{
		Type:        "rich",
		Title:       feed.Items[0].Title,
		URL:         feed.Items[0].GUID,
		Description: "A gathering of people from all walks of life; from seasoned game developers and students just entering the industry, to hobbyists and enthusiasts.  We gather once a month to discuss everything related to video game design and development, engage with industry, and do that \"networking\" thing.",
		Color:       14496563, // #dd3333 as decimal
		Fields: []*discordgo.MessageEmbedField{
			&discordgo.MessageEmbedField{
				Name:   "Next Event",
				Value:  nextEvent,
				Inline: true},
			&discordgo.MessageEmbedField{
				Name:   "Location",
				Value:  "KTTC\nFleming College\nRoom D1111",
				Inline: true}},
		Thumbnail: &discordgo.MessageEmbedThumbnail{
			URL: imageSource},
	})
}
