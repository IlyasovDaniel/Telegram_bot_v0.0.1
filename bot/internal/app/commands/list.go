package commands

import (
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

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

func init() {
	registeredCommands["list"] = (*Commander).List
}
