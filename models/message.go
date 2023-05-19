package models

import "time"

// representation of a message sent by a user within a channel between
// two (or maybe more) users
type Message struct {
	ID           string    // unique id of the message
	Userid       string    // id of user author of the message
	Content      string    // actual message
	CreationDate time.Time // client-side time of the message sent
	Channelid    string    // the channel where the message was sent
}

// representation of the channel itself that allows two users to communicate
type Channel struct {
	ID      string
	Members []string
	Open    bool
}
