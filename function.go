package joke

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"
	"os"

	"github.com/afritzler/joke-skill/pkg/types"
)

const (
	randomJokeAPI = "https://official-joke-api-l2zakgd5zq-ez.a.run.app/random_joke"
)

// RandomJoke returns a random joke
func RandomJoke(w http.ResponseWriter, r *http.Request) {
	var replies []interface{}
	url := getenv("JOKEAPI", randomJokeAPI)

	res, err := http.Get(url)
	if err != nil {
		log.Printf("failed to get url %s: %v", url, err)
		replies = append(replies, generateTextMessage(types.RequestErrorMessage, 0))
		returnWithReply(w, replies)
	}

	defer res.Body.Close()
	body, readErr := ioutil.ReadAll(res.Body)
	if readErr != nil {
		log.Printf("failed to read body: %v", err)
		replies = append(replies, generateTextMessage(types.RequestErrorMessage, 0))
		returnWithReply(w, replies)
	}

	response := types.RandomJokeResponse{}
	err = json.Unmarshal([]byte(body), &response)
	if err != nil {
		log.Printf("failed to unmarshal body: %v", err)
		replies = append(replies, generateTextMessage(types.RequestErrorMessage, 0))
		returnWithReply(w, replies)
	} else {
		replies = append(replies, generateTextMessage(response.Setup, 0))
		replies = append(replies, generateTextMessage(response.Punchline, 1))
	}
	returnWithReply(w, replies)
}

func returnWithReply(w http.ResponseWriter, replies []interface{}) {
	output, err := json.Marshal(types.Replies{Replies: replies})
	if err != nil {
		log.Printf("failed to marshal replies: %+v\n", err)
		panic("something went wrong here with marshalling")
	}
	w.Header().Set("content-type", "application/json")
	w.Write(output)
	return
}

func getenv(key, fallback string) string {
	value := os.Getenv(key)
	if len(value) == 0 {
		return fallback
	}
	return value
}

func generateTextMessage(text string, delay int) types.TextMessage {
	return types.TextMessage{
		Type:    types.TextType,
		Content: text,
		Delay:   delay,
	}
}
