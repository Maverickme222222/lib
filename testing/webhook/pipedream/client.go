// Package pipedream is a 3d party service implementation of the Webhook interface
package pipedream

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/r3labs/sse/v2"

	"github.com/kappapay/backend/lib/golang/src/kappa/testing/webhook"
)

const (
	httpComponentID = "sc_zzir4Q" // id for creating pipedream HTTP sources
	basePath        = "https://api.pipedream.com/v1"
)

// Client webhook implementation of pipedream
type Client struct {
	apiKey    string
	orgID     string
	eventChan map[string]chan *sse.Event // map containing channels used for passing webhook events
}

// New client init
func New(apiKey string, orgID string) (*Client, error) {
	return &Client{
		apiKey:    apiKey,
		orgID:     orgID,
		eventChan: make(map[string]chan *sse.Event),
	}, nil
}

// EnsureWebhook will ensure a webhook exists, create or get
func (c *Client) EnsureWebhook(ctx context.Context, name string, opts ...webhook.EnsureOption) (*webhook.HTTPWebhook, error) {
	// eval options
	option := &webhook.EnsureOptions{}
	for _, opt := range opts {
		opt(option)
	}

	var w *webhook.HTTPWebhook

	// list sources and try to find source by name
	after := ""
	for {
		sourceList, err := c.listSources(after)
		if err != nil {
			return nil, err
		}
		for _, source := range sourceList.Data {
			if source.Name == name {
				w = &webhook.HTTPWebhook{
					ID:   source.ID,
					URL:  source.ConfiguredProps.HTTPInterface.EndpointURL,
					Name: source.Name,
				}
				break
			}
		}
		if sourceList.PaginationInfo.Count == sourceList.PaginationInfo.TotalCount {
			break
		}
		after = sourceList.PaginationInfo.EndCursor
	}

	// otherwise create a source
	if w == nil {
		newWebhook, err := c.createSource(name)
		if err != nil {
			return nil, err
		}
		w = newWebhook
	}

	if option.Read {
		return w, c.initListener(w.ID)
	}

	return w, nil
}

// ReadEvent will read one event at a time, will block
func (c *Client) ReadEvent(ctx context.Context, webhookID string, outValue interface{}) (bool, error) {
	ch, ok := c.eventChan[webhookID]
	if !ok {
		return false, fmt.Errorf("webhook not ensured with read")
	}
	var msg *sse.Event
	select {
	case <-ctx.Done():
		return false, ctx.Err()
	case msg = <-ch:
	}

	if msg.Data == nil || len(msg.Data) <= 1 {
		return false, nil
	}

	var event webhookEvent
	err := json.Unmarshal(msg.Data, &event)
	if err != nil {
		return false, nil
	}

	return true, json.Unmarshal([]byte(event.Event.BodyRaw), &outValue)
}

func (c *Client) createSource(name string) (*webhook.HTTPWebhook, error) {
	// json request body
	data := map[string]interface{}{
		"component_id": httpComponentID,
		"name":         name,
		"configured_props": map[string]interface{}{
			"emitBodyOnly":   false,
			"resStatusCode":  "200",
			"resContentType": "application/json",
			"resBody":        "{ \"success\": true }",
		},
	}
	if c.orgID != "" {
		data["org_id"] = c.orgID
	}
	jsonData, err := json.Marshal(data)
	if err != nil {
		return nil, err
	}

	client := http.Client{}
	req, err := http.NewRequest("POST", basePath+"/sources", bytes.NewBuffer(jsonData))
	if err != nil {
		return nil, err
	}
	// set auth headers
	req.Header = http.Header{
		"Content-Type":  {"application/json"},
		"Authorization": {"Bearer " + c.apiKey},
	}

	res, err := client.Do(req)
	if err != nil {
		return nil, err
	}

	var s sourceCreateResponse
	err = json.NewDecoder(res.Body).Decode(&s)
	if err != nil {
		return nil, err
	}

	return &webhook.HTTPWebhook{
		ID:   s.Data.ID,
		URL:  s.Data.ConfiguredProps.HTTPInterface.EndpointURL,
		Name: s.Data.Name,
	}, nil

}

func (c *Client) listSources(after string) (*sourceListResponse, error) {
	client := http.Client{}
	uriPath := basePath + "/users/me/sources"
	if c.orgID != "" {
		uriPath = basePath + fmt.Sprintf("/orgs/%v/sources", c.orgID)
	}
	req, err := http.NewRequest("GET", uriPath, nil)
	if err != nil {
		return nil, err
	}

	// set auth headers
	req.Header = http.Header{
		"Content-Type":  {"application/json"},
		"Authorization": {"Bearer " + c.apiKey},
	}
	// set query params
	q := req.URL.Query()
	q.Add("limit", "100")
	q.Add("after", after)
	req.URL.RawQuery = q.Encode()

	res, err := client.Do(req)
	if err != nil {
		return nil, err
	}

	var list sourceListResponse
	err = json.NewDecoder(res.Body).Decode(&list)
	if err != nil {
		return nil, err
	}

	return &list, nil
}

func (c *Client) initListener(webhookID string) error {
	if _, ok := c.eventChan[webhookID]; ok {
		return nil
	}
	events := make(chan *sse.Event, 100)
	c.eventChan[webhookID] = events

	client := sse.NewClient(fmt.Sprintf("https://api.pipedream.com/sources/%v/sse", webhookID))
	client.Headers["Authorization"] = "Bearer " + c.apiKey

	return client.SubscribeChan("messages", events)
}

var _ webhook.Webhook = &Client{}
