package persistence

import (
	"bufio"
	"encoding/json"
	"fmt"
	"io"
	"os"
)

type Client struct {
	storagePath string
	data        NotificationsByLocation
}

func NewClient(storagePath string) (*Client, error) {
	file, err := os.OpenFile(storagePath, os.O_CREATE, 0644)
	if err != nil {
		return nil, fmt.Errorf("Error creating persistence file: %v", err)
	}
	defer file.Close()
	// unbuffered read; don't expect large files but could be issue in future if saving lots of appointments
	data, err := io.ReadAll(file)
	if err != nil {
		return nil, fmt.Errorf("Error reading persistence file: %v", err)
	}

	var notifications NotificationsByLocation
	if len(data) == 0 {
		bufio.NewWriter(file).Write([]byte("{}"))
	} else {
		err = json.Unmarshal(data, &notifications)
		if err != nil {
			return nil, fmt.Errorf("Error unmarshaling persistence data: %v", err)
		}
	}

	return &Client{storagePath: storagePath, data: notifications}, nil
}

func (c *Client) saveData() error {
	file, err := os.Create(c.storagePath)
	if err != nil {
		return fmt.Errorf("Error opening persistence file: %v", err)
	}
	defer file.Close()

	encoder := json.NewEncoder(file)

	err = encoder.Encode(c.data)
	if err != nil {
		return fmt.Errorf("Error marshaling persistence data: %v", err)
	}

	return nil
}

func (c *Client) DiffAndSaveData(proposedChanges NotificationsByLocation) (NotificationsByLocation, error) {
	additions := make(NotificationsByLocation)
	for location, newNotifications := range proposedChanges {
		existingNotifications, exists := c.data[location]
		if !exists {
			additions[location] = newNotifications
			continue
		}
		for _, newNotification := range newNotifications {
			found := false
			for _, existingNotification := range existingNotifications {
				if existingNotification.LocationTime.Equal(newNotification.LocationTime) {
					found = true
					break
				}
			}
			if !found {
				additions[location] = append(additions[location], newNotification)
			}
		}
	}

	c.data = proposedChanges

	err := c.saveData()
	if err != nil {
		return nil, fmt.Errorf("Error saving persistence data: %v", err)
	}

	return additions, nil
}
