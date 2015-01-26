// Package gomeh provides a Go interface for Meh.com daily deals
package gomeh

import (
	"crypto/tls"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"time"
)

// Meh contains the details of the meh (deal, poll, and video)
type Meh struct {
	Deal  Deal  `json:"deal"`
	Poll  Poll  `json:"poll"`
	Video Video `json:"video"`
}

// Deal contains deal specific data
type Deal struct {
	Features       string    `json:"features"`
	ID             string    `json:"id"`
	Items          []Item    `json:"items"`
	Photos         []string  `json:"photos"`
	Title          string    `json:"title"`
	SoldOutAt      time.Time `json:"soldOutAt"`
	Specifications string    `json:"specifications"`
	Story          Story     `json:"story"`
	Theme          Theme     `json:"theme"`
	URL            string    `json:"url"`
	Topic          Topic     `json:"topic"`
}

// Topic contains data about the forum topic for the deal including how many
// comments, replies, votes the topic has and the topic's URL.
type Topic struct {
	CommentCount int       `json:"commentCount"`
	CreatedAt    time.Time `json:"createdAt"`
	ID           string    `json:"id"`
	ReplyCount   int       `json:"replyCount"`
	URL          string    `json:"url"`
	VoteCount    int       `json:"voteCount"`
}

// Theme for the deal including the accent color, background color,
// background image, and whether the foreground is light or dark.
type Theme struct {
	AccentColor     string `json:"accentColor"`
	BackgroundColor string `json:"backgroundColor"`
	BackgroundImage string `json:"backgroundImage"`
	Foreground      string `json:"foreground"`
}

// Story contains the deal's story, including the title and the body in Markdown
// format.
type Story struct {
	Title string `json:"title"`
	Body  string `json:"body"`
}

// Item represents the individual products available for purchase and contain
// attributes (key/value pairs such as Color: Georgia Red), condition (New or
// Refurbished), the item's unique identifier, the price, and a photo URL.
type Item struct {
	Attributes []interface{} `json:"attributes"`
	Condition  string        `json:"condition"`
	ID         string        `json:"id"`
	Price      int           `json:"price"`
	Photo      string        `json:"photo"`
}

// Poll contains the deals accompanying user poll
type Poll struct {
	Answers   []Answer  `json:"answers"`
	ID        string    `json:"id"`
	StartDate time.Time `json:"startDate"`
	Title     string    `json:"title"`
	Topic     Topic     `json:"topic"`
}

// Video describes the details of the deals accompanying video
type Video struct {
	ID        string    `json:"id"`
	StartDate time.Time `json:"startDate"`
	Title     string    `json:"title"`
	URL       string    `json:"url"`
	Topic     Topic     `json:"topic"`
}

// Answer describes a polls possible answers
type Answer struct {
	ID        string `json:"id"`
	Text      string `json:"text"`
	VoteCount int    `json:"voteCount"`
}

// String returns the title of the Meh with the price and prefixes [Sold Out] if
// there is a SoldOutAt time set.
func (m Meh) String() (s string) {
	if m.SoldOut() {
		s = fmt.Sprintf("[Sold Out] %v - $%v", m.Deal.Title, m.Deal.Items[0].Price)
	} else {
		s = fmt.Sprintf("%v - $%v", m.Deal.Title, m.Deal.Items[0].Price)
	}
	return
}

// SoldOut returns true if the Meh has sold out.
func (m Meh) SoldOut() bool {
	return !m.Deal.SoldOutAt.IsZero()
}

const mehURL = "https://api.meh.com/1/current.json?apikey="

func callAPI(apikey string) (body []byte, err error) {
	// Configure http transport and client settings for SSL1.1 and AES256 cipher
	tr := &http.Transport{
		TLSClientConfig: &tls.Config{
			CipherSuites: []uint16{tls.TLS_RSA_WITH_AES_256_CBC_SHA},
			MaxVersion:   tls.VersionTLS11,
		},
	}
	client := &http.Client{Transport: tr}
	url := string(mehURL + apikey)

	// Create and execute the GET request
	res, err := client.Get(url)
	if err != nil {
		return
	}

	// Close the connection when done
	defer res.Body.Close()

	// Read out the response body []byte
	body, err = ioutil.ReadAll(res.Body)
	if err != nil {
		return
	}

	return
}

// GetMeh returns the current Meh
func GetMeh(apikey string) (meh Meh, err error) {

	// Get current data from api as []byte
	apiData, err := callAPI(apikey)
	if err != nil {
		return
	}

	// Populate a Meh with the deserialized json
	err = json.Unmarshal(apiData, &meh)
	if err != nil {
		return
	}

	return
}
