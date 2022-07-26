package pkg

import (
	"crawl/util"
	"github.com/gtuk/discordwebhook"
	"log"
)

const UserBotDiscord = "Bot News Go"

func BotPushNewGoToDiscord(config util.Config, title string, href string, image string) {
	var embeds []discordwebhook.Embed

	var user = UserBotDiscord
	embed := discordwebhook.Embed{}
	embed.Title = &title
	embed.Url = &href

	var thumbnail discordwebhook.Image
	thumbnail.Url = &image
	embed.Image = &thumbnail
	embeds = append(embeds, embed)
	message := discordwebhook.Message{
		Username: &user,
		Embeds:   &embeds,
	}
	err := discordwebhook.SendMessage(config.URL_WEBHOOK_DISCORD, message)
	if err != nil {
		log.Println(err)
	}
}
