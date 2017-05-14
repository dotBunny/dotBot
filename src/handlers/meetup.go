package handlers

import (
	"github.com/bwmarrin/discordgo"
	"github.com/mmcdole/gofeed"
)

func HandleMeetup(session *discordgo.Session, message *discordgo.MessageCreate) {

	fp := gofeed.NewParser()
	feed, _ := fp.ParseURL("https://www.meetup.com/PTBOGameDev/events/rss/")

	_, _ = session.ChannelMessageSend(message.ChannelID, feed.Items[0].Title)
	_, _ = session.ChannelMessageSendEmbed(message.ChannelID, &discordgo.MessageEmbed{
		Type:        "rich",
		Title:       feed.Items[0].Title,
		URL:         feed.Items[0].GUID,
		Description: feed.Items[0].Description,
	})

	/*<description><![CDATA[<p><img style="float:left; margin-right:4px" src="https://secure.meetupstatic.com/photos/event/1/1/2/9/event_457804393.jpeg" alt="photo" class="photo" />PTBO Game Development</p> <p><p><span>You can show up anytime after 6:30PM, we're just calling 7PM the official start time.</span></p> <br></p> <p>K9J 7B1, ON   - Canada</p> <p>Wednesday, May 17 at 7:00 PM</p> <p>14</p> <p>https://www.meetup.com/PTBOGameDev/events/237855732/</p> ]]></description>*/
}

/*Fields: []*discordgo.MessageEmbedField{
    &discordgo.MessageEmbedField{Name: "Name", Value: g.Name, Inline: true},
    &discordgo.MessageEmbedField{Name: "ID", Value: g.ID, Inline: true},
    &discordgo.MessageEmbedField{Name: "Owner name", Value: guildOwner.Username + "#" + guildOwner.Discriminator, Inline: true},
    &discordgo.MessageEmbedField{Name: "Owner ID", Value: guildOwner.ID, Inline: true},
    &discordgo.MessageEmbedField{Name: "Users", Value: fmt.Sprintf("%v", userCount), Inline: true},
    &discordgo.MessageEmbedField{Name: "Bots", Value: fmt.Sprintf("%v", botCount), Inline: true},
    &discordgo.MessageEmbedField{Name: "Percent", Value: fmt.Sprintf("%v", int(percent)) + "%", Inline: true},
},*/
