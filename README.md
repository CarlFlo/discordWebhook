# Discord Webhook

This module aims to make it easy to interact with discord through a webhook.


## Usage

```go
    var webhook discordWebhook.Webhooks

    // You're able to add as many webhooks as you like
	webhook.AddWebhook("https://discord.com/api/webhooks/...")
	webhook.AddWebhook("https://discord.com/api/webhooks/...")

    // Below so are you able to add a message and embed.

	webhook.SetContent("Text here will look like a normal discord message")

    // Adds an embed to the message.
	webhook.AddEmbed(discordWebhook.Embed{
		Title:       "Here is the title",
		Description: "Here is a description for the embed",
		Color:       0x58b9ff,
		Fields: []discordWebhook.Field{{
			Name:  "Field 1",
			Value: "Some text for the first field",
		}, {
			Name:  "Field 2",
			Value: "Some text for the second field",
		}},
		Author: discordWebhook.Author{
			Name:     "Test Author",
			Url:      "https://www.author_image.png",
			Icon_url: "https://www.author_image_icon.png",
		}, Image: discordWebhook.Image{
			Url: "https://www.image.png",
		},
		Timestamp: discordWebhook.GetCurrentTimestamp(),
		Footer: discordWebhook.Footer{
			Text:     "Footer text",
			Icon_url: "https://www.tiny_footer_icon.png",
		},
	})

	webhook.Send()

```


## GitHub Copilot

This project was made with the help of [GitHub Copilot](https://copilot.github.com/).

## Issues

* Setting the color does not work. It defaults to black