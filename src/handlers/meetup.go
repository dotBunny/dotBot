package handlers

import (
	"strings"

	"fmt"

	"github.com/bwmarrin/discordgo"
	"github.com/mmcdole/gofeed"
	"golang.org/x/net/html"
)

func HandleMeetup(session *discordgo.Session, message *discordgo.MessageCreate) {

	fp := gofeed.NewParser()
	feed, _ := fp.ParseURL("https://www.meetup.com/PTBOGameDev/events/rss/")

	parsedHTML := html.NewTokenizer(strings.NewReader(feed.Items[0].Description))

	imageURL := ""

	currentToken := parsedHTML.Next()
	for currentToken != html.ErrorToken {
		switch {
		case currentToken == html.StartTagToken:
			tag := parsedHTML.Token()
			fmt.Printf(tag.Data)
			isImage := tag.Data == "img"
			if isImage {
				fmt.Printf("FOUND IMAGE")
				for _, a := range tag.Attr {
					if a.Key == "src" {
						fmt.Printf("FOUND SRC")
						imageURL = a.Val
						continue
					}
				}
			}
		}
		currentToken = parsedHTML.Next()
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
				Value:  "Wednesday, May 17 at 7:00 PM",
				Inline: true},
			&discordgo.MessageEmbedField{
				Name:   "Location",
				Value:  "KTTC\nFleming College\nRoom D1111",
				Inline: true}},
		Thumbnail: &discordgo.MessageEmbedThumbnail{
			URL: imageURL},
	})

	/*<description>
	<![CDATA[
		<p>
			<img style="float:left; margin-right:4px" src="https://secure.meetupstatic.com/photos/event/1/1/2/9/event_457804393.jpeg" alt="photo" class="photo" />
			PTBO Game Development
		</p>
		<p><p>
		<span>You can show up anytime after 6:30PM, we're just calling 7PM the official start time.</span>
		</p> <br></p>
		<p>K9J 7B1, ON   - Canada</p>
		<p>Wednesday, May 17 at 7:00 PM</p> <p>14</p> <p>https://www.meetup.com/PTBOGameDev/events/237855732/</p> ]]></description>*/
}
