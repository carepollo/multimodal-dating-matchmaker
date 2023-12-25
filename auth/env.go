package auth

import "time"

type Environment struct {
	Datasources struct {
		MongoDB struct {
			Uri string `mapstructure:"uri"`
		} `mapstructure:"mongodb"`
		Redis struct {
			Uri      string `mapstructure:"uri"`
			Password string `mapstructure:"password"`
		} `mapstructure:"redis"`
		Neo4j struct {
			Uri      string `mapstructure:"uri"`
			Username string `mapstructure:"username"`
			Password string `mapstructure:"password"`
		} `mapstructure:"neo4j"`
	} `mapstructure:"datasources"`
}

// Gender represents the Gender message.
type Gender struct {
	Id   string
	Name string
}

// Relationship represents the Relationship message.
type Relationship struct {
	Id   string
	Name string
}

// Education represents the Education message.
type Education struct {
	Id          string
	Site        string
	ProgramType string
	ProgramName string
}

// Work represents the Work message.
type Work struct {
	Id   string
	Site string
	Role string
}

// Location represents the Location message.
type Location struct {
	Id        string
	City      string
	Country   string
	Longitude float32
	Latitude  float32
}

// Thing represents the Thing message.
type Thing struct {
	Id       string
	WhatItIs string
	NameOfIt string
}

// User represents the User message.
type User struct {
	Id             string
	Name           string
	Lastname       string
	PhoneNumber    string
	Status         string // 'pending', 'verified', 'deactivated'
	Email          string
	Password       string
	AboutMe        string
	Religion       string
	Score          int64
	Age            int32
	Height         float32
	Weight         float32
	Pictures       []string
	LastConnection time.Time
	CreationDate   time.Time
	Gender         Gender
	LastLocation   Location
	Relationship   []Relationship // what the user is looking for
	Work           []Work         // optinal logs for formal work experience like linkedin
	Education      []Education    // optional logs for formal education like linkedin
	Tags           []Thing        // concepts that the user associate to itself
	Likes          []Thing
	Dislikes       []Thing
}
