package models

import "time"

// to represent records of stuff the user have done
type LifeRecord struct {
	Institution string   `json:"institution"` // where it happened
	Title       string   `json:"title"`       // title of experience
	Description string   `json:"description"` // description of experience
	Attachments []string `json:"attachments"` // array with files user appeneded (prefer to use only pics, audio, video)
}

// this is the represetantion of user within the system
type User struct {
	ID             string          `json:"id"`             // ObjectID of mongodb
	Name           string          `json:"name"`           // name to display
	TypeName       string          `json:"typename"`       // alias or real name
	AccountType    string          `json:"accountType"`    // which method used to create account: google, facebook, email/password
	Username       string          `json:"username"`       // codename for user when created by other methods
	Status         string          `json:"status"`         // status of account, if is just created, verified, deactivated, banned, etc.
	LastConnection time.Time       `json:"lastConnection"` // last time saw a message
	Tags           map[string]bool `json:"tags"`           // tags associated to user to cluster it or clasify it
	Email          string          `json:"email"`          // registered email
	Password       string          `json:"password"`       // hashed and salted password
	Age            int8            `json:"age"`            // age of user
	InterestedIn   []string        `json:"interestedIn"`   // gender that is looking for
	Relationship   []string        `json:"relationship"`   // type of relationship
	Pictures       []string        `json:"pictures"`       // user pictures associated to profile, stores the file url
	LastLocation   Location        `json:"lastLocation"`   // location of the last login
	Phone          string          `json:"phone"`          // phone number of user, must be unique in db, it should be an optional field
	Weight         float32         `json:"weight"`         // weight of person
	Height         float32         `json:"heigth"`         // height of person
	Religion       string          `json:"religion"`       // religion of person
	Education      []LifeRecord    `json:"education"`      // degrees, educational background, etc
	Work           []LifeRecord    `json:"work"`           // where has he worked at
	AboutMe        string          `json:"aboutMe"`        // is the self description of the user
	Likes          map[string]bool `json:"likes"`          // topics that the user likes/dislikes
	Spectrum       interface{}     `json:"spectrum"`       // filters defined for profile search
}
