package models

import "time"

// represents the status of the user on the platform
type UserStatus int

const (
	ACTIVE      UserStatus = iota // status of user that has verified the account
	PENDING                       // for users that created an account but haven't verified authenticity
	DEACTIVATED                   // for banned users, users that deleted the account, etc.
)

// to represent records of stuff the user have done
type LifeRecord struct {
	Institution string   `json:"institution"` // where it happened
	Title       string   `json:"title"`       // title of experience
	Description string   `json:"description"` // description of experience
	Attachments []string `json:"attachments"` // array with files user appeneded (prefer to use only pics, audio, video)
}

// static configurations for result filtering at the moment of fetching profiles and matchmaking
type Spectrum interface{} // TODO: define the structure

// this is the represetantion of user within the system
type User struct {
	ID             string          `json:"id"`               // unique ID of user within the system ObjectID of mongodb
	Name           string          `json:"name"`             // name to display
	Status         UserStatus      `json:"status"`           // status of account: pending verification, verified, deactivated, banned, etc
	LastConnection time.Time       `json:"lastConnection"`   // last time saw a message
	Tags           []string        `json:"tags"`             // one-word adjectives to clasify itself
	Email          string          `json:"email"`            // registered email
	Password       string          `json:"password"`         // hashed and salted password
	Age            int8            `json:"age"`              // age of user
	InterestedIn   []string        `json:"interestedIn"`     // gender that is looking for
	Relationship   []string        `json:"relationship"`     // type of relationship
	Pictures       []string        `json:"pictures"`         // user pictures associated to profile, stores the file url
	LastLocation   Location        `json:"lastLocation"`     // location of the last login
	Phone          string          `json:"phone"`            // phone number of user, must be unique in db, it should be an optional field
	Weight         float32         `json:"weight"`           // weight of person
	Height         float32         `json:"heigth"`           // height of person
	Religion       string          `json:"religion"`         // religion of person
	Education      []LifeRecord    `json:"education"`        // degrees, educational background, etc
	Work           []LifeRecord    `json:"work"`             // where has he worked at
	AboutMe        string          `json:"aboutMe"`          // is the self description of the user
	Preferences    map[string]bool `json:"preferences"`      // topics that the user likes/dislikes, key is concet, value if like or dislike
	Spectrum       Spectrum        `json:"spectrum"`         // filters defined for profile search
	IsMatchedWith  interface{}     `json:"isMatchedWith"`    // a relationship with another user on graph
	HasUnmatched   interface{}     `json:"hasUnmatchedWith"` // another type of relationship within the graph with another user
	HasBlocked     interface{}     `json:"hasBlocked"`       // another type of relationship within the graph with another user, the non-relationship
}
