package main

import (
	"strings"

	Handlers "./handlers"

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
	case ".commands":
		Handlers.HandleCommands(session, message)
	case ".help":
		Handlers.HandleHelp(session, message)
	case ".ping":
		Handlers.HandlePing(session, message)
	case ".meetup":
		Handlers.HandleMeetup(session, message)
	default:
		return
	}
}
