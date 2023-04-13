package main

import (
	"flag"
	"fmt"
	"log"

	"github.com/MrBorisT/birthday_reminder_tg_bot/internal/config"
	"github.com/mymmrac/telego"
	th "github.com/mymmrac/telego/telegohandler"
	tu "github.com/mymmrac/telego/telegoutil"
)

func main() {
	dev := flag.Bool("dev", true, "developer mode")
	flag.Parse()

	var options []telego.BotOption
	if *dev {
		options = append(options, telego.WithDefaultDebugLogger())
	}
	config.Init()
	bot, err := telego.NewBot(config.ConfigData.Token, options...)

	if err != nil {
		log.Fatalln("initializing bot", err)
	}

	updates, _ := bot.UpdatesViaLongPolling(nil)

	defer bot.StopLongPolling()

	bh, _ := th.NewBotHandler(bot, updates)

	defer bh.Stop()

	defer bot.StopLongPolling()

	bh.Handle(func(bot *telego.Bot, update telego.Update) {
		_, _ = bot.SendMessage(tu.Message(
			tu.ID(update.Message.Chat.ID),
			fmt.Sprintf("Hello %s!", update.Message.From.FirstName),
		))
	}, th.CommandEqual("start"))

	bh.Handle(func(bot *telego.Bot, update telego.Update) {
		_, _ = bot.SendMessage(tu.Message(
			tu.ID(update.Message.Chat.ID),
			"Unknown command, use /start",
		))
	}, th.AnyCommand())

	bh.Start()
}
