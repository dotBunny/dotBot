package commands

import (
	"github.com/bwmarrin/discordgo"
)

//https://docs.unity3d.com/2017.1/Documentation/ScriptReference/

// Untiy responds in the channel with a stub of info on the command + link to more info
func Unity(session *discordgo.Session, message *discordgo.MessageCreate) {

	// Grab Command
	//contentSplit := strings.Split(message.Content, ",")
	//searchString := contentSplit[0]
	//_, _ = session.ChannelMessageSend(message.ChannelID, "Pong!")
}
