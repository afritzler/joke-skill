package types

const (
	TextType            = "text"
	ButtonsType         = "buttons"
	RequestErrorMessage = "Looks like there was a hick-up in my though process. Could you please try again?"
)

// Replies
type Replies struct {
	Replies []interface{} `json:"replies"`
}

// TextMessage defines a response of type text message.
// Example:
// {
// 	"type": "text",
//	"delay": 2,
//	"content": "MY_TEXT",
// }
type TextMessage struct {
	Type    string `json:"type"`
	Content string `json:"content"`
	Delay   int    `json:"delay,omitempty"`
}

// RandomJokeResponse defines a joke api response
// {"id":196,"type":"general","setup":"What did the traffic light say to the car as it passed?","punchline":"Don't look I'm changing!"}
type RandomJokeResponse struct {
	ID        int    `json:"ID"`
	Type      string `json:"type"`
	Setup     string `json:"setup"`
	Punchline string `json:"punchline"`
}
