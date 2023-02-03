// Package slack ...
package slack

import (
	"fmt"

	"github.com/parnurzeal/gorequest"
)

// Field ...
type Field struct {
	Title string `json:"title"`
	Value string `json:"value"`
	Short bool   `json:"short"`
}

// Message ...
type Message struct {
	Color      *string   `json:"color"`
	Title      *string   `json:"title"`
	Text       *string   `json:"text"`
	Fields     []*Field  `json:"fields"`
	Timestamp  *int64    `json:"ts"`
	MarkdownIn *[]string `json:"mrkdwn_in"`
}

// Payload ...
type Payload struct {
	ServiceName string    `json:"username,omitempty"`
	IconEmoji   string    `json:"icon_emoji,omitempty"`
	Channel     string    `json:"channel,omitempty"`
	Text        string    `json:"text,omitempty"`
	Messages    []Message `json:"attachments,omitempty"`
}

// AddField ...
func (message *Message) AddField(field Field) *Message {
	message.Fields = append(message.Fields, &field)
	return message
}

// Client ...
type Client struct {
	WebhookURL string
}

// Send the slack notification
func (c Client) Send(payload Payload) error {
	if c.WebhookURL == "" {
		return fmt.Errorf("webhook url cannot be blank")
	}
	request := gorequest.New()
	resp, _, err := request.
		Post(c.WebhookURL).
		Send(payload).
		End()

	if len(err) > 1 {
		return err[0]
	}
	if resp.StatusCode >= 400 {
		return fmt.Errorf("error sending alert: %v", resp.Status)
	}

	return nil
}

/*
//usage
func main() {
	webhookUrl := "https://hooks.slack.com/services/AA/BB/CC"
	client := &Client{WebhookURL: webhookUrl}

	message1 := Message{}
	message1.AddField(Field{Title: "Author", Value: "Payout"}).AddField(Field{Title: "Alert", Value: "Testing"})
	payload := Payload{
		Text:      ":robot_face:",
		ServiceName:  "payout",
		Channel:   "#kappa-pay-operations",
		IconEmoji: ":rotating_light:",
		Messages:  []Message{message1},
	}
	err := client.Send(payload)
	if err != nil {
		fmt.Printf("error: %s\n", err)
	}
}
*/
