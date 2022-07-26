package pkg

import (
	"crawl/util"
	"github.com/gtuk/discordwebhook"
	"log"
)

const UserBotDiscord = "Bot News Go"

func BotPushNewGoToDiscord(title string, href string, image string) {
	var config util.Config
	var embeds []discordwebhook.Embed

	var user = UserBotDiscord
	embed := discordwebhook.Embed{}
	embed.Title = &title
	embed.Url = &href

	var thumbnail discordwebhook.Thumbnail
	thumbnail.Url = &image
	embed.Thumbnail = &thumbnail
	embeds = append(embeds, embed)
	message := discordwebhook.Message{
		Username: &user,
		Embeds:   &embeds,
	}
	err := discordwebhook.SendMessage(config.URL_WEBHOOK_DISCORD, message)
	if err != nil {
		log.Fatal(err)
	}
}
