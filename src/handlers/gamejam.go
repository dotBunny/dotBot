//http://ptbogamejam.com/data/latest.json

package handlers

import (
	"bytes"

	Helpers "../helpers"
	"github.com/bwmarrin/discordgo"
)

// GameJamLatest JSON Data Structure
type GameJamLatest struct {
	Title         string  `json:"title"`
	Link          string  `json:"link"`
	ICS           string  `json:"ics"`
	DateLong      string  `json:"date_long"`
	DateShort     string  `json:"date_short"`
	DateCountdown string  `json:"date_countdown"`
	Countdown     string  `json:"countdown"`
	HoursWork     int     `json:"hours_work"`
	HoursEDU      int     `json:"hours_edu"`
	Prizes        int     `json:"prizes"`
	Message       string  `json:"message"`
	MapLat        float64 `json:"map-lat"`
	MapLong       float64 `json:"map-long"`
	MapZoom       int     `json:"map-zoom"`
	MarkerLat     float64 `json:"marker-lat"`
	MarkerLong    float64 `json:"marker-long"`
	Location      string  `json:"location"`
	AddressOne    string  `json:"address-1"`
	AddressTwo    string  `json:"address-2"`
	Registration  string  `json:"registration"`
	Spots         int     `json:"spots"`
	LatestVideo   string  `json:"latest_video"`
}

// HandleGameJam responds in the channel with the next PTBO Game Jam information
func HandleGameJam(session *discordgo.Session, message *discordgo.MessageCreate) {

	var data GameJamLatest
	Helpers.GetJSON("http://ptbogamejam.com/data/latest.json", &data)

	var buffer bytes.Buffer
	buffer.WriteString("http://ptbogamejam.com")
	buffer.WriteString(data.Link)
	eventLink := buffer.String()

	_, _ = session.ChannelMessageSendEmbed(message.ChannelID, &discordgo.MessageEmbed{
		Type:        "rich",
		Title:       data.Title,
		URL:         eventLink,
		Description: "The PTBO Game Jam is a bi-annual not-for-profit event managed by industry professionals hosted in beautiful Peterborough Ontario, Canada. It is a gathering of people from all walks of life; from seasoned game developers and students just entering the industry, to hobbyists and enthusiasts all eager to engage and see what they can produce in a short timeframe.",
		Color:       52479,
		Fields: []*discordgo.MessageEmbedField{
			&discordgo.MessageEmbedField{
				Name:   "Next Event",
				Value:  data.DateLong,
				Inline: true},
			&discordgo.MessageEmbedField{
				Name:   "Location",
				Value:  Helpers.Clean(data.Location),
				Inline: true}},
		Thumbnail: &discordgo.MessageEmbedThumbnail{
			URL: Helpers.GetIconGameJam()},
	})
}
