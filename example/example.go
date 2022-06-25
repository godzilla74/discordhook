package main

import (
	"discordhook"
	"log"
)

var (
	username    = "DiscordHook Example"
	message     = "Hello this is example content of string"
	description = "A little description for you guys :)"
	url         = "webhook_url"
)

func main() {
	simple()
	embed()
	embedWithFields()
}

func simple() {
	message := discordhook.Message{
		Username: &username,
		Content:  &message,
	}

	err := discordhook.SendMessage(url, message)
	if err != nil {
		log.Fatal(err)
	}
}

func embed() {
	embeds := make([]discordhook.Embed, 0)
	embeds = append(embeds, discordhook.Embed{
		Title:       &message,
		Description: &description,
	})

	message := discordhook.Message{
		Username: &username,
		Content:  &message,
		Embeds:   &embeds,
	}

	err := discordhook.SendMessage(url, message)
	if err != nil {
		log.Fatal(err)
	}
}

func embedWithFields() {
	var (
		ccy       = "BTC"
		value     = "0.0002"
		inline    = true
		typeName  = "type"
		valueType = "join"
	)

	fields := make([]discordhook.Field, 0)
	fields = append(fields, discordhook.Field{
		Name:   &ccy,
		Value:  &value,
		Inline: &inline,
	})
	fields = append(fields, discordhook.Field{
		Name:   &typeName,
		Value:  &valueType,
		Inline: &inline,
	})
	fields = append(fields, discordhook.Field{
		Name:  &ccy,
		Value: &value,
	})

	embeds := make([]discordhook.Embed, 0)
	embeds = append(embeds, discordhook.Embed{
		Title:       &message,
		Description: &description,
		Fields:      &fields,
	})

	message := discordhook.Message{
		Username: &username,
		Content:  &message,
		Embeds:   &embeds,
	}

	err := discordhook.SendMessage(url, message)
	if err != nil {
		log.Fatal(err)
	}
}
