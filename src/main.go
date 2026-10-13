package main

import (
	"fmt"
	"os"
	"os/signal"

	"github.com/bwmarrin/discordgo"
    "github.com/joho/godotenv"
)

var session *discordgo.Session

func init() {
	// read .env file and search for the bot token
	err := godotenv.Load();
  	if err != nil {
    	fmt.Println("Error loading .env file", err)
		return
  	}
	token := os.Getenv("TOKEN");

	// Create a new Discord session
	session, err = discordgo.New("Bot " + token);
	if err != nil {
		fmt.Println("error creating Discord session ", err)
		return
	}
}

func main() {
	// start the discord bot
	session.AddHandler(func(s *discordgo.Session, r *discordgo.Ready) {
		fmt.Printf("Logged in as %s", r.User.String())
	})

	err := session.Open()
	if err != nil {
		fmt.Printf("could not open session: %s", err)
	}

	sigch := make(chan os.Signal, 1)
	signal.Notify(sigch, os.Interrupt)
	<-sigch

	err = session.Close()
	if err != nil {
		fmt.Printf("could not close session gracefully: %s", err)
	}
}
