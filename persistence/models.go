package persistence

import "time"

type Notification struct {
	// TimeNotified time.Time `json:"time_notified"`
	LocationTime time.Time `json:"location_time"`
}

// NotificationsByLocation is a map of location strings to Notification structs
type NotificationsByLocation map[string][]Notification
