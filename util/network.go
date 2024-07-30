package util

import (
	"OneBotAssistant/config"
	"github.com/go-resty/resty/v2"
	"time"
)

func getRestyClient(headers map[string]string, client *resty.Client) *resty.Client {
	if client == nil {
		client = resty.New()
	} else {
		return client
	}

	if config.CONFIG.Proxy != "" {
		client.SetProxy(config.CONFIG.Proxy)
	}

	client.SetHeaders(headers)
	client.SetTimeout(10 * time.Second)
	return client
}

func RequestGET(link string, headers map[string]string, client *resty.Client) ([]byte, error) {
	client = getRestyClient(headers, client)
	response, err := client.R().Get(link)
	if err != nil {
		return []byte(""), err
	}
	return response.Body(), nil
}

func RequestPOST(link, data string, headers map[string]string, client *resty.Client) ([]byte, error) {
	client = getRestyClient(headers, client)

	response, err := client.R().SetBody(data).Post(link)
	if err != nil {
		return []byte(""), err
	}
	return response.Body(), nil
}
