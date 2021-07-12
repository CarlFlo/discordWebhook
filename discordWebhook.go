package discordWebhook

import (
	"bytes"
	"encoding/json"
	"io/ioutil"
	"net/http"
	"net/url"
	"time"

	"github.com/CarlFlo/malm"
)

type Webhooks struct {
	webhooksUrls  []*url.URL
	webhookParams WebhookParams
}

type WebhookParams struct {
	Content string  `json:"content,omitempty"`
	Embeds  []Embed `json:"embeds,omitempty"`
}

type Embed struct {
	Title       string  `json:"title,omitempty"`
	Description string  `json:"description,omitempty"`
	Url         string  `json:"url,omitempty"`
	Color       int     `json:"clor,omitempty"`
	Author      Author  `json:"author,omitempty"`
	Fields      []Field `json:"fields,omitempty"`
	Image       Image   `json:"image,omitempty"`
	Timestamp   string  `json:"timestamp,omitempty"`
	Footer      Footer  `json:"footer,omitempty"`
}

type Author struct {
	Name     string `json:"name,omitempty"`
	Url      string `json:"url,omitempty"`
	Icon_url string `json:"icon_url,omitempty"`
}

type Field struct {
	Name  string `json:"name,omitempty"`
	Value string `json:"value,omitempty"`
}

type Image struct {
	Url string `json:"url,omitempty"`
}

type Footer struct {
	Text     string `json:"text,omitempty"`
	Icon_url string `json:"icon_url,omitempty"`
}

func (wh *Webhooks) AddWebhook(webhooks string) {

	url, err := url.Parse(webhooks)
	if err != nil {
		malm.Error(err.Error())
	}

	wh.webhooksUrls = append(wh.webhooksUrls, url)
}

func (wh *Webhooks) SetContent(newContent string) {
	wh.webhookParams.Content = newContent
}

func (wh *Webhooks) AddEmbed(e Embed) {
	wh.webhookParams.Embeds = append(wh.webhookParams.Embeds, e)
}

func GetCurrentTimestamp() string {
	return time.Now().Format("2006-01-02 15:04:05")
}

func (wh *Webhooks) Send() {

	// Prepares the content of the message.
	jsonContent, _ := json.Marshal(wh.webhookParams)
	buffer := bytes.NewBuffer(jsonContent)

	for _, url := range wh.webhooksUrls {

		req := &http.Request{
			Method: http.MethodPost,
			URL:    url,
			Header: http.Header{
				"Content-Type": []string{"application/json"},
			},
			Body: ioutil.NopCloser(buffer),
		}

		resp, err := http.DefaultClient.Do(req)
		if err != nil {
			malm.Error(err.Error())
		}
		defer resp.Body.Close()

		// If statuscode is 400
		if resp.StatusCode == 400 {
			malm.Error("400 Bad Request")
		}

	}
}
