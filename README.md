# Discord Webhook

This module aims to make it both simple and easy to interact with discord through a webhook.


## Usage

Here is some example code to get started

Each of the variable in the **embed** is optional. 
Meaning that it is fine to only have a *title* and *description*.

For images so are only png files tested.
Additionally, the URL's must start with **https://** or else Discord wont accept the message.

```go
var webhook discordWebhook.Webhooks

// You're able to add as many webhooks as you like
webhook.AddWebhook("https://discord.com/api/webhooks/...")
webhook.AddWebhook("https://discord.com/api/webhooks/...")

// Below so are you able to add a message and embed. 
// It is aesthetically pleasing to only pick one.

// The content will look like a normal discord message. The best for simple messages
webhook.SetContent("Text here will look like a normal discord message")

// Adds an embed to the message. Most sutable for more 'complex' messages
webhook.AddEmbed(discordWebhook.Embed{
	Title:       "Here is the title",
	Description: "Here is a description for the embed",
	Color:       0x58b9ff, // Hex color
	Fields: []discordWebhook.EmbedField{{
		Name:  "Field 1",
		Value: "Some text for the first field",
	}, {
		Name:  "Field 2",
		Value: "Some text for the second field",
	}},
	Author: discordWebhook.EmbedAuthor{
		Name:     "Author name",
		Url:      "URL that will make the name clickable",
		Icon_url: "Author image icon URL",
	}, Image: discordWebhook.EmbedImage{
		Url: "URL to an image",
	},
	Timestamp: discordWebhook.GetCurrentTimestamp(), // Use this method to generate a timestamp for the current time
	Footer: discordWebhook.EmbedFooter{
		Text:     "Footer text",
		Icon_url: "URL to footer icon",
	},
})

// Allows you to set the username of the webhook message bot
webhook.SetUsername("Example username")

// Allows you to set the avatar of the webhook bot
webhook.SetAvatar("URL to image")

webhook.Send()
```

## GitHub Copilot

This project was made with the help of [GitHub Copilot](https://copilot.github.com/).

## Todo

- [X] Basic functionality
- [X] Support for embeds
- [ ] Support for multiple embedded images
- [ ] Support for files

## Issues

* Setting the color does not work. It defaults to black