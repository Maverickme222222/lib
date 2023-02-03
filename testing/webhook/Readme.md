# Webhook Integration Testing
This package is intended for testing webhook flows. The idea here is to have a way to create webhooks by code, and then listen/subscribe to those webhook events.

The interface gives you two methods:
```go
type Webhook interface {
    EnsureWebhook(ctx context.Context, name string, opts ...EnsureOption) (*HTTPWebhook, error)
    ReadEvent(ctx context.Context, webhookId string, outValue interface{}) (bool, error)
}
```
- EnsureWebhook will ensure that webhook with the given name exists. Name here is like a unique tag.
  - you can pass EnsureOptions to this function, atm only EnsureWithRead is supported (it will activate listener on the events)
- ReadEvent is a blocking function which waits until a new webhook events is triggered, and will unmarshall into outValue.

HTTPWebhook model contains info about a created webhook.
```go
type HTTPWebhook struct {
	Id   string // the id of the webhook
	URL  string // webhook url, this url can be used when setting up webhooks with vendors
	Name string // unique name
}
```

***

### Pipedream Implementation
Pipedream is a 3d party service which can be used to setup workflows.
In our case we only use the basic functionality needed for webhooks.
- Source: event sources are like workflow triggers. In our code we only create HTTP sources, which will give us back some URL to point to.
- Server-Sent Events (SSE): is basically a stream of events where we can subscribe and get events in realtime. When a source is created, we are able to subscribe via SSE to it and receive updates.

***

##### Example usage
We want to test out Sila webhook events:
1. Create webhook in Pipedream, either via dashboard or via code. Make sure to use a code friendly name `(sila-webhook-test)`
2. Take the new webhook url and use it to setup a webhook in Sila dashboard.
3. Example code below will listen to events coming to that webhook:
```go
type testModel struct {
	Id int `json:"id"`
}

func main() {
	webhookName := "sila-webhook-integration-test"
    // timeout used to control execution time
    ctx, cancel := context.WithTimeout(context.Background(), 1*time.Minute)
	defer cancel()
	client, err := pipedream.New(__apikey__, __orgkey__)
	checkErr(err)
	// create or get the webhook source
	w, err := client.EnsureWebhook(ctx, webhookName, webhook.EnsureWithRead())
	checkErr(err)
	fmt.Printf("ensured webhook %#v\n", w)
	for {
		var v testModel
		ok, err := client.ReadEvent(ctx, w.ID, &v)
		if err == context.Canceled {
			fmt.Println("timeout reached")
			return
		}
		checkErr(err)
		if ok {
			fmt.Printf("Got message %#v\n", v)
		}
	}
}
```
