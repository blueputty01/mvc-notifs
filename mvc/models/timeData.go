package models

import (
	"fmt"
	"strings"
	"time"
)

type TimeData struct {
	LocationId    int                `json:"LocationId"`
	FirstOpenSlot AppointmentSummary `json:"FirstOpenSlot"`
}

type AppointmentSummary struct {
	Time time.Time
}

const ctLayout = "01/02/2006 03:04 PM"

func (ct *AppointmentSummary) UnmarshalJSON(b []byte) (err error) {
	s := strings.Trim(string(b), "\"")
	s = strings.Split(s, "Next Available: ")[1]
	if s == "null" {
		ct.Time = time.Time{}
		return
	}
	ct.Time, err = time.Parse(ctLayout, s)
	return
}

func (ct *AppointmentSummary) MarshalJSON() ([]byte, error) {
	if ct.Time.IsZero() {
		return []byte("null"), nil
	}
	return []byte(fmt.Sprintf("\"%s\"", ct.Time.Format(ctLayout))), nil
}

var nilTime = (time.Time{}).UnixNano()

func (ct *AppointmentSummary) IsSet() bool {
	return !ct.Time.IsZero()
}
