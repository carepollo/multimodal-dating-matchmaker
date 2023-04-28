package models

// this is the represetantion of user within the system
type User struct {
	ID       string `json:"id"`
	Name     string `json:"name"`
	TypeName string `json:"typename"`
	Location struct {
		X float64 `json:"x"`
		Y float64 `json:"y"`
	} `json:"location"`
	Age          int8        `json:"age"`
	Tags         interface{} `json:"tags"`
	Email        string      `json:"email"`
	Password     string      `json:"password"`
	InterestedIn []string    `json:"interestedIn"`
	LookingFor   []string    `json:"lookingFor"`
}
