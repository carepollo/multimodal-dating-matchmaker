package models

import "time"

// a log created after a user views a profile
type StatsLog struct {
	ID         string        // id of log
	Userid     string        // id of the user whose profile is being seen
	Visitantid string        // id of the user that is viewing the profile
	VisitTime  time.Duration // time the user is in the profile, stored in seconds
	VisitDate  time.Time     // date at the moment of the visit
	Visitor    User          // data of the user that is visiting (should exclude name, id and profile picture)
	Viewed     struct {
		SectionName string        // section of the user that the visitor is watching at the moment (subject to client logic)
		VisitTime   time.Duration // time the visitor spent watching such section
	}
	DidInteract bool // if user triggered any interaction with user after visit, match/message/any/etc
}
