package models

import "time"

// representation of the payment of a user to access premium features
type Bill struct {
	ID             string    // unique id of bill
	Userid         string    // id of user
	CreationDate   time.Time // when the user has payed, the coverage of the service starts, ends 30 days later
	ExpirationDate time.Time // when the coverage of the service ends
	HasNext        bool      // if the system has to automatically renew the subscription
	Total          float64   // total of the bill
	Method         string    // payment method chosed
}
