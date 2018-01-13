package views

import (
	"log"

	"../models"
)

// Data is the top level structure that views expect data
// to come in.
type Data struct {
	Alert *Alert
	User  *models.User
	Yield interface{}
}

// Alert is used to render Bootstrap Alert messages in templates
type Alert struct {
	Level   string
	Message string
}

// PublicError displays errors fit for public consumption
type PublicError interface {
	error
	Public() string
}

// Constants for AlertLevels for Alert types
const (
	AlertLvlError   = "danger"
	AlertLvlWarning = "warning"
	AlertLvlInfo    = "info"
	AlertLvlSuccess = "success"

	// AlertMsgGeneric is displayed when any random error
	// is encountered by our backend.
	AlertMsgGeneric = "Something went wrong. Pleast try " +
		"again, and contact us if the problem persists."
)

// SetAlert returns a type asserted Error
func (d *Data) SetAlert(err error) {
	var msg string
	if pErr, ok := err.(PublicError); ok {
		msg = pErr.Public()
	} else {
		log.Println(err)
		msg = AlertMsgGeneric
	}
	d.Alert = &Alert{
		Level:   AlertLvlError,
		Message: msg,
	}
}

// AlertError constructs an Alert for a Data receiever
func (d *Data) AlertError(msg string) {
	d.Alert = &Alert{
		Level:   AlertLvlError,
		Message: msg,
	}
}
