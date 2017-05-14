package main

import (
	"strings"

	Command "./commands"

	"github.com/bwmarrin/discordgo"
)

func messageHandler(session *discordgo.Session, message *discordgo.MessageCreate) {

	// No need for the bot to process its own commands really
	if message.Author.ID == BotID {
		return
	}

	// Grab Command
	contentSplit := strings.Split(message.Content, ",")
	command := contentSplit[0]

	// Check and Dispatch!
	switch command {
	case ".discord":
		Command.Discord(session, message)
	case ".gamejam":
		Command.GameJam(session, message)
	case ".help":
		Command.Help(session, message)
	case ".meetup":
		Command.Meetup(session, message)
	case ".ping":
		Command.Ping(session, message)
	case ".unity":
		Command.Unity(session, message)
	default:
		return
	}
}
