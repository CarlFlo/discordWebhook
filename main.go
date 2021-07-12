package discordWebhook

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/CarlFlo/malm"
)

type Webhooks struct {
	webhooksUrls []string
}

func (wh *Webhooks) Add(webhooks string) {
	wh.webhooksUrls = append(wh.webhooksUrls, webhooks)
}

func (wh *Webhooks) Send() {

	values := map[string]string{"content": "This is a very cool and rad test"}
	jsonValue, _ := json.Marshal(values)

	for _, url := range wh.webhooksUrls {
		resp, err := http.Post(url, "application/json", bytes.NewBuffer(jsonValue))
		if err != nil {
			panic(err)
		}
		defer resp.Body.Close()

		// Handle response
		if resp.StatusCode != 200 {
			malm.Error(fmt.Sprintf("Response code %d", resp.StatusCode))
		}

	}
}
