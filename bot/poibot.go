package main

import (
	"flag"
	"fmt"
	"time"
	"os"
	"regexp"
	"io/ioutil"
	"log"

	"github.com/bwmarrin/discordgo"
)

var (
	Token string
	BotId string

	POIINSTRUCTIONS = "poi_instructions: Instructions Poi!\npoi: POI!\nlolicon: For terrible people Poi!\ngo sleep: I'll tell people to sleep!\nfuck: POI!\nmistakes were made: mistakes poi\npoi_lewd: For lewd people poi.\npoi_tired: I'm tired poi.\npoi_memes: Memes poi."


	// Where the directories for images should be placed
	// images should be in directories with their relevant commands
	// assets/instruction/image
	// e.g. /assets/poi/poi.jpg
	assetRoot string = "/home/linhai/go/assets/"

	// read from a file for the regexes?
	assets map[string][]string = make(map[string][]string)
	
	imagePaths [18]string

	lewd [4]string
	lewdct int = 0

	tired [2]string
	tiredct int = 0

	meme [3]string
	memect int = 0

	lolicon [2]string
	loliconct int = 0

	fuck [2]string
	fuckct int = 0

)

func init() {

	dirs, rootDirErr := ioutil.ReadDir(assetRoot)

	if rootDirErr != nil {
		log.Fatal(rootDirErr)
	}

	for _, d := range dirs {
		if d.IsDir() {
			files, _ := ioutil.ReadDir(assetRoot + d.Name())
			var tempArr []string
			for _, f := range files {
				tempArr = append(tempArr, f.Name())
				fmt.Println("Loaded %s", assetRoot + d.Name() + "/" + f.Name())
			}
			assets[d.Name()] = tempArr
		}
	}

	flag.StringVar(&Token, "t", "", "Bot Token")
	flag.Parse()

	// setup globals
	imagePaths[0] = "poi.jpg"
        imagePaths[1] = "exodias_d.jpg"
        imagePaths[2] = "fuck.gif"
        imagePaths[3] = "go_sleep.png"
        imagePaths[4] = "i_came.jpg"
        imagePaths[5] = "last_night.jpg"
        imagePaths[6] = "lewd_meme.png"
        imagePaths[7] = "lolicon_rights.png"
        imagePaths[8] = "meme_oniichan.jpg"
        imagePaths[9] = "mistakes_were_made.jpg"
        imagePaths[10] = "monopoly.jpg"
        imagePaths[11] = "no_peeking.jpg"
        imagePaths[12] = "oniichan.jpg"
        imagePaths[13] = "ride.jpg"
        imagePaths[14] = "tired.jpg"
        imagePaths[15] = "traps_gay.png"
        imagePaths[16] = "watching.jpg"
	imagePaths[17] = "fuck.jpg"

	lewd[0] = imagePaths[1]
        lewd[1] = imagePaths[4]
        lewd[2] = imagePaths[13]
        lewd[3] = imagePaths[16]

	tired[0] = imagePaths[5]
        tired[1] = imagePaths[14]

	meme[0] = "communism.jpg"
        meme[1] = imagePaths[6]
        meme[2] = imagePaths[8]

	lolicon[0] = imagePaths[7]
        lolicon[1] = imagePaths[10]

	fuck[0] = imagePaths[2]
	fuck[1] = imagePaths[17]

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

	instructionRegex, _ := regexp.Compile("^poi_instructions$")
	poiRegex, _ := regexp.Compile("(^[pP][oO][iI]( [pP][oO][iI])*$)|(^[cC][aA][nN][cC][eE][rR]$)")
	whatIsPoiRegex, _ := regexp.Compile("[wW][hH][aA][tT] [iI][sS] [pP][oO][iI]\\??")
//	lewdRegex, _ := regexp.Compile("^poi_lewd$")
//	fuckRegex, _ := regexp.Compile("^[fF][uU][cC][kK]$")
//	tiredRegex, _ := regexp.Compile("^poi_tired$")
//	sleepRegex, _ := regexp.Compile("^[gG][oO] [sS][lL][eE]{2}[pP][!]?$")
//	memeRegex, _ := regexp.Compile("^poi_memes$")
//	mistakesRegex, _ := regexp.Compile("^mistakes were made$")
//	loliconRegex, _ := regexp.Compile("^lolicon$")
	

	// Ignore all messages created by bot
	// AVOIDS INFINITE CALLS
	if m.Author.ID == BotId {
		return
	}

	// Received a message
        fmt.Printf("%20s %20s %20s > %s\n", m.ChannelID, time.Now().Format(time.Stamp), m.Author.Username, m.Content)

	// User wants to know about poi
	if whatIsPoiRegex.MatchString(m.Content) {
                _, err := s.ChannelMessageSend(m.ChannelID, "Poi is poi.")
                if err != nil {
                        fmt.Println("ERR: failed to send message")
                }
		return
        }

	// User wants instructions
	if instructionRegex.MatchString(m.Content) {
		_, err := s.ChannelMessageSend(m.ChannelID, POIINSTRUCTIONS)
                if err != nil {
                        fmt.Println("ERR: failed to send message")
                }
                return
	}

	// POI
	if poiRegex.MatchString(m.Content) {
		fmt.Println("POI!")
		singleMessage(s, m, imagePaths[0])
		return
	}
/*
	// lewds
	if lewdRegex.MatchString(m.Content) {
		manyMessage(s, m, lewd[:], &lewdct)
		return
	}

	// fuck
	if fuckRegex.MatchString(m.Content) {
		manyMessage(s, m, fuck[:], &fuckct) 
                return
	}

	// tired
	if tiredRegex.MatchString(m.Content) {
		manyMessage(s, m, tired[:], &tiredct)
                return
        }

	// go sleep
	if sleepRegex.MatchString(m.Content) {
		singleMessage(s, m, imagePaths[3])
                return
	}

	// meme
        if memeRegex.MatchString(m.Content) {
		manyMessage(s, m, meme[:], &memect)
                return
        }

	// lolicon
        if loliconRegex.MatchString(m.Content) {
		manyMessage(s, m, lolicon[:], &loliconct)
                return
        }

	// mistakes
	if mistakesRegex.MatchString(m.Content) {
		singleMessage(s, m, imagePaths[9])
                return
        }
*/
}

func manyMessage(s *discordgo.Session, m *discordgo.MessageCreate, imgs []string, ct *int) bool {
/*
    if(*ct+1 >= len(imgs)) {
        *ct = 0
    } else {
       *ct += 1
    }

    image, err := os.Open(assetRoot + imgs[*ct])

    if err != nil {
        s.ChannelMessageSend(m.ChannelID, "I couldn't find the image poi :(")
        fmt.Println("%s\n", assetRoot + imgs[*ct])
        return false
    }

    _, err = s.ChannelFileSendWithMessage(m.ChannelID, "", imgs[*ct], image)
    if err != nil {
        fmt.Println("ERR: Failed to send message,", err)
        return false
    }
*/
    return true


}

func singleMessage(s *discordgo.Session, m *discordgo.MessageCreate, img string) bool {

    image, err := os.Open(assetRoot + img)

    if err != nil {
        s.ChannelMessageSend(m.ChannelID, "I couldn't find the image poi :(")
        fmt.Println("%s\n", assetRoot + img)
        return false
    }

    _, err = s.ChannelFileSendWithMessage(m.ChannelID, "", img, image)
    if err != nil {
        fmt.Println("ERR: Failed to send message,", err)
        return false
    }

    return true

}
