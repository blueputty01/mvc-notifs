package mvc

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"mvc-notif/mvc/models"
	"net/http"
	"strings"
	"time"
)

type Client struct {
	BaseURL string
}

func getJSVar(body []byte, varName string) (string, error) {
	prefix := fmt.Sprintf("var %s", varName)
	start := bytes.Index(body, []byte(prefix))
	if start == -1 {
		return "", fmt.Errorf("variable %s not found", varName)
	}
	start += len(prefix)
	end := bytes.IndexAny(body[start:], "\n;")
	if end == -1 {
		return "", fmt.Errorf("end of variable %s not found", varName)
	}
	raw := string(body[start : start+end])
	cleaned := strings.Trim(raw, " =")
	return cleaned, nil
}

func NewClient(baseURL string) *Client {
	return &Client{BaseURL: baseURL}
}

type Appointment struct {
	Data          models.LocationData
	NextAvailable time.Time
}

func (c *Client) GetNextAvailable() ([]Appointment, error) {
	resp, err := http.Get(c.BaseURL)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("unexpected status code: %d", resp.StatusCode)
	}

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, fmt.Errorf("error reading response body: %w", err)
	}

	locationData, err := c.GetLocationData(body)
	if err != nil {
		return nil, fmt.Errorf("error getting location data: %w", err)
	}

	timeData, err := c.GetTimeData(body)
	if err != nil {
		return nil, fmt.Errorf("error getting time data: %w", err)
	}

	result := make([]Appointment, len(locationData))

	for idx, t := range locationData {
		locTime, ok := timeData[t.Id]
		if !ok {
			continue
		}
		result[idx] = Appointment{
			Data:          t,
			NextAvailable: locTime.FirstOpenSlot.Time,
		}
	}

	return result, nil
}

type TimeLookup map[int]models.TimeData

func (c *Client) GetTimeData(body []byte) (TimeLookup, error) {
	timeString, err := getJSVar(body, "timeData")
	if err != nil {
		return nil, fmt.Errorf("Error occurred while getting time data: %w", err)
	}

	var times []models.TimeData
	err = json.Unmarshal([]byte(timeString), &times)
	if err != nil {
		return nil, fmt.Errorf("Error occurred while unmarshaling time data: %w", err)
	}

	result := make(TimeLookup)
	for _, t := range times {
		result[t.LocationId] = t
	}

	return result, nil
}

func (c *Client) GetLocationData(body []byte) ([]models.LocationData, error) {
	locationString, err := getJSVar(body, "locationData")
	if err != nil {
		return nil, err
	}

	var locations []models.LocationData
	err = json.Unmarshal([]byte(locationString), &locations)
	if err != nil {
		return nil, err
	}

	// transform data because lat/long are stored in undetermined format
	scale := 736.45607073
	x := 39.3698364
	y := 73.6242132
	for i := range locations {
		locations[i].Lat = (float64(locations[i].Lat)/scale + x)
		locations[i].Long = (float64(locations[i].Long)/scale + y)
	}

	return locations, nil
}
