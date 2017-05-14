package helpers

var (
	// GameDevMeetupIconURI - PTBO Game Dev Thumbnail
	GameDevMeetupIconURI = func() string { return "http://ptbogamejam.com/files/bot/ptbogamedev-icon.jpg" }

	// GameJamIconURI - PTBO Game Jam Thumbnail
	GameJamIconURI = func() string { return "http://ptbogamejam.com/files/bot/ptbogamejam-icon.jpg" }
)

// GetDiscordInviteURI returns the invite URI for Discord
func GetDiscordInviteURI() (url string) {
	return "http://discord.me/ptbogamejam"
}
