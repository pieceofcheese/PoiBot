package main

import (
	"flag"
	"fmt"
	"time"
	"os"

	"github.com/bwmarrin/discordgo"
)

var (
	Token string
	BotId string
	PoiImage string = "/home/linhai/go/assets/poi.jpg"
)

func init() {

	flag.StringVar(&Token, "t", "", "Bot Token")
	flag.Parse()

}

func main() {

	// CONNECTION SETUP

	// Create a new Discord session using provided bot token.
	// dg is the discordGo bot object, err is the error out
	dg, err := discordgo.New("Bot " + Token)
	if err != nil {
		fmt.Println("error creating Discord session,", err)
		return
	}

	// Get Account info
	u, err := dg.User("@me")
	if err != nil {
		fmt.Println("error obtaining account details,", err)
	}
	BotId = u.ID

	// EVENT HANDLER ADDING

	//Add messageCreate as an event to be handled
	dg.AddHandler(messageCreate)

	// OPEN CONNECTION and start receiving

	// Open socket and begin listening
	err = dg.Open()
	if err != nil {
		fmt.Println("error opening connection,", err)
		return
	}

	fmt.Println("Bot is now running. Press CTRL-C to exit.")
	//simple way to keep running program until CTRL-C
	<-make(chan struct{})

	// CLOSE CONNECTION

	return
}

//Handles any messageCreated events
func messageCreate(s *discordgo.Session, m *discordgo.MessageCreate) {

	// Ignore all messages created by bot
	// AVOIDS INFINITE CALLS
	if m.Author.ID == BotId {
		return
	}

	if m.Content == "poi" {
		/*_, err := s.ChannelMessageSend(m.ChannelID, "Poi!")
		if err != nil {
			fmt.Println("ERR: Failed to send message,", err)
		}*/

		image, err := os.Open(PoiImage)

		if err != nil {
			s.ChannelMessageSend(m.ChannelID, "I couldn't find the image poi :(")
			fmt.Println("%s\n", PoiImage)
			return
		}

		_, err = s.ChannelFileSendWithMessage(m.ChannelID, "Poi!", "poi0.jpg", image)
		if err != nil {
			fmt.Println("ERR: Failed to send message,", err)
			return
		}
	}

	// Received a message
	fmt.Printf("%20s %20s %20s > %s\n", m.ChannelID, time.Now().Format(time.Stamp), m.Author.Username, m.Content)
}
