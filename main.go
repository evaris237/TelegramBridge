package main

import (
	"fmt"
	"log"
	"os"
	"strings"

	"github.com/bwmarrin/discordgo"
)

func composer(status, event, actor, repo, workflow, link string) string {
	var text string

	// choose icon based on the build status
	icons := map[string]string{
		"failure":   "❗️❗️❗️",
		"cancelled": "❕❕❕",
		"success":   "✅✅✅",
	}

	replacer := strings.NewReplacer("_", "\\_", "-", "\\-", ".", "\\.")

	// removing symbols to avoid markdown parser error
	event = replacer.Replace(event)

	repo = replacer.Replace(repo)

	actor = replacer.Replace(actor)

	// Message text composing
	text = icons[strings.ToLower(status)] + "  **" + strings.ToUpper(event) + "**\n"
	text += "was made at " + repo + " \nby " + actor + "\n"
	text += "Check here " + "[" + workflow + "](" + link + ")"

	return text
}

func linkgen(repo, event string) string {
	context := map[string]string{
		"issue_comment":               "issues",
		"issues":                      "issues",
		"pull_request":                "pulls",
		"pull_request_review_comment": "pulls",
		"push":                        "commits",
		"project_card":                "projects",
	}

	event = context[strings.ToLower(event)]

	// generates link based on the triggered event
	return fmt.Sprintf("https://github.com/%s/%s/", repo, event)
}

func sendToDiscord(discord *discordgo.Session, channelID, message string) error {
	_, err := discord.ChannelMessageSend(channelID, message)
	if err != nil {
		return fmt.Errorf("unable to send message to Discord channel: %v", err)
	}
	return nil
}

func main() {

	var (
		// inputs provided by Github Actions runtime
		// should be defined in the action.yml
		token   = os.Getenv("INPUT_TOKEN")
		channel = os.Getenv("INPUT_CHANNEL")
		status  = os.Getenv("INPUT_STATUS")
		event   = os.Getenv("INPUT_EVENT")
		actor   = os.Getenv("INPUT_ACTOR")

		// github environment context
		workflow = os.Getenv("GITHUB_WORKFLOW")
		repo     = os.Getenv("GITHUB_REPOSITORY")
		// commit   = os.Getenv("GITHUB_SHA")
	)

	// Create a new Discord session using bot token
	discord, err := discordgo.New("Bot " + token)
	if err != nil {
		log.Fatalf("unable to create Discord session: %v", err)
	}

	// Open a websocket connection to Discord
	err = discord.Open()
	if err != nil {
		log.Fatalf("unable to open Discord websocket connection: %v", err)
	}
	defer discord.Close()

	// link to the commit
	link := linkgen(repo, event)

	// Prepare message to send
	msg := composer(status, event, actor, repo, workflow, link)

	// Send message to Discord channel
	err = sendToDiscord(discord, channel, msg)
	if err != nil {
		log.Fatalf("unable to send message to Discord channel: %v", err)
	}
}
