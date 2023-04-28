package models

type LifeRecord struct {
	Institution string `json:"institution"` // where it happened
	Title       string `json:"title"`       // title of experience
	Description string `json:"description"` // description of experience
}

// this is the represetantion of user within the system
type User struct {
	ID           string          `json:"id"`       // ObjectID of mongodb
	Name         string          `json:"name"`     // name to display
	TypeName     string          `json:"typename"` // alias or real name
	Tags         map[string]bool `json:"tags"`
	Email        string          `json:"email"`        // registered email
	Password     string          `json:"password"`     // hashed and salted password
	Age          int8            `json:"age"`          // age of user
	InterestedIn []string        `json:"interestedIn"` // gender that is looking for
	Relationship []string        `json:"relationship"` // type of relationship
	Education    []LifeRecord    `json:"education"`
	Location     Location        `json:"location"`
	Weight       float32         `json:"weight"`
	Height       float32         `json:"heigth"`
	Religion     string          `json:"religion"`
	Work         []LifeRecord    `json:"work"`
	AboutMe      string          `json:"aboutMe"` // is the self description of the user
	Likes        map[string]bool `json:"likes"`   // topics that the user likes/dislikes
	Spectrum     interface{}     `json:"spectrum"`
}
