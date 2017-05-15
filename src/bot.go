package main

import (
	"fmt"
	"log"
	"os"
	"path"
	"path/filepath"

	"github.com/BurntSushi/toml"
	"github.com/bwmarrin/discordgo"
)

// Variables used for command line parameters
var (
	Token string
	BotID string
)

// Config is an external config type
type Config struct {
	Token string
}

func main() {
	var config = ReadConfig()
	Token = config.Token

	// Create a new Discord session using the provided bot token.
	dg, err := discordgo.New("Bot " + Token)
	if err != nil {
		fmt.Println("error creating Discord session,", err)
		return
	}

	// Get the account information.
	u, err := dg.User("@me")
	if err != nil {
		fmt.Println("error obtaining account details,", err)
	}

	// Store the account ID for later use.
	BotID = u.ID

	// Register messageCreate as a callback for the messageCreate events.
	dg.AddHandler(messageHandler)

	// Open the websocket and begin listening.
	err = dg.Open()
	if err != nil {
		fmt.Println("error opening connection,", err)
		return
	}

	fmt.Println("Bot is now running.  Press CTRL-C to exit.")
	// Simple way to keep program running until CTRL-C is pressed.
	<-make(chan struct{})
	return
}

// ReadConfig gets the local config file
func ReadConfig() Config {

	dir, pathError := filepath.Abs(filepath.Dir(os.Args[0]))
	if pathError != nil {
		log.Fatal("Unable to determine path of application")
	}

	configPath := path.Join(dir, "config.toml")

	_, err := os.Stat(configPath)

	if err != nil {
		log.Fatal("Config file is missing: ", configPath)
	}

	var config Config
	if _, err := toml.DecodeFile(configPath, &config); err != nil {
		log.Fatal(err)
	}

	return config
}
