package models

import "time"

// Gender represents the Gender message.
type Gender struct {
	ID   string
	Name string
}

// Relationship represents the Relationship message.
type Relationship struct {
	ID   string
	Name string
}

// Education represents the Education message.
type Education struct {
	ID          string
	Site        string
	ProgramType string
	ProgramName string
}

// Work represents the Work message.
type Work struct {
	ID   string
	Site string
	Role string
}

// Location represents the Location message.
type Location struct {
	ID        string
	City      string
	Country   string
	Longitude float32
	Latitude  float32
}

// Thing represents the Thing message.
type Thing struct {
	ID       string
	WhatItIs string
	NameOfIt string
}

// User represents the User message.
type User struct {
	ID             string
	Name           string
	Lastname       string
	PhoneNumber    string
	Status         string
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
