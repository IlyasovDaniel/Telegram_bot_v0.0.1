package commands

import (
	"log"

	"github.com/IlyasovDaniel/Telegram_bot_v0.0.1/internal/service/product"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

type Commander struct {
	bot            *tgbotapi.BotAPI
	productService *product.Service
}

func NewCommander(	
	bot *tgbotapi.BotAPI,
	productService *product.Service,
) *Commander {
	return &Commander{
		bot:            bot,
		productService: productService,
	}
}

func (c *Commander) Help(inputMessage *tgbotapi.Message) {
	msg := tgbotapi.NewMessage(inputMessage.Chat.ID,
		"/help - help\n"+
			"/list - list products")
	c.bot.Send(msg)
}
func (c *Commander) List(inputMessage *tgbotapi.Message) {
	outputMgsText := "Here all the products: \n\n"

	products := c.productService.List()
	for _, p := range products {
		outputMgsText += p.Title
		outputMgsText += "\n"
	}

	msg := tgbotapi.NewMessage(inputMessage.Chat.ID, outputMgsText)
	c.bot.Send(msg)
}

func (c *Commander) Default(inputMessage *tgbotapi.Message) {
	log.Printf("[%s] %s", inputMessage.From.UserName, inputMessage.Text)

	msg := tgbotapi.NewMessage(inputMessage.Chat.ID, "You wrote: "+inputMessage.Text)

	c.bot.Send(msg)
}
