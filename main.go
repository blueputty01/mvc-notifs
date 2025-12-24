package main

import (
	"flag"
	"fmt"
	"mvc-notif/mvc"
	"mvc-notif/notif"
	"mvc-notif/utils"
	"os"
	"strings"
	"time"
)

const inputTimeFormat = "01/02/2006"

var urlFlag = flag.String("url", "https://telegov.njportal.com/njmvc/AppointmentWizard/12", "The URL of the portal")
var startDateFlag = flag.String("start", "", "The start date to filter appointments (MM/DD/YYYY)")
var endDateFlag = flag.String("end", "", "The end date to filter appointments (MM/DD/YYYY)")
var destinationNumberFlag = flag.String("to", "", "The destination phone number for notifications")
var ignoredLocationsFlag = flag.String("ignore", "", "Comma-separated list of location names to ignore")
var maxDistanceFlag = flag.Float64("max-distance", 100.0, "Maximum distance in miles from the center point")
var centerLatFlag = flag.Float64("lat", 40, "Center latitude for distance filtering")
var centerLonFlag = flag.Float64("lon", 74, "Center longitude for distance filtering")

func main() {
	flag.Parse()

	ignoreLocationSubstrs := strings.Split(*ignoredLocationsFlag, ",")

	startDate := time.Now()
	if *startDateFlag != "" {
		var err error
		startDate, err = time.Parse(inputTimeFormat, *startDateFlag)
		if err != nil {
			panic(fmt.Errorf("invalid start date: %w", err))
		}
	}

	endDate := time.Date(9999, time.December, 31, 23, 59, 59, 999999999, time.UTC)
	if *endDateFlag != "" {
		var err error
		endDate, err = time.Parse(inputTimeFormat, *endDateFlag)
		if err != nil {
			panic(fmt.Errorf("invalid end date: %w", err))
		}
	}

	mvcClient := mvc.NewClient(*urlFlag)

	twilioAccountSid := os.Getenv("TWILIO_ACCOUNT_SID")
	if twilioAccountSid == "" {
		panic("TWILIO_ACCOUNT_SID environment variable not set")
	}
	twilioAuthToken := os.Getenv("TWILIO_AUTH_TOKEN")
	if twilioAuthToken == "" {
		panic("TWILIO_AUTH_TOKEN environment variable not set")
	}
	twilioOriginNumber := os.Getenv("TWILIO_ORIGIN_NUMBER")
	if twilioOriginNumber == "" {
		panic("TWILIO_ORIGIN_NUMBER environment variable not set")
	}
	notifClient := notif.NewClient(
		twilioAccountSid,
		twilioAuthToken,
		twilioOriginNumber,
	)

	locations, err := mvcClient.GetNextAvailable()
	if err != nil {
		panic(err)
	}

	// filter by distance from lat long
	center := utils.Point{Lat: *centerLatFlag, Lon: *centerLonFlag}

	filtered := make([]mvc.Appointment, 0)
	for _, loc := range locations {
		distance := utils.Haversine(
			center,
			utils.Point{Lat: loc.Data.Lat, Lon: loc.Data.Long},
		)
		// fmt.Printf("Location: %s Point: %+v Distance: %f\n", loc.Data.Name, point, distance)
		if distance <= *maxDistanceFlag {
			ignore := false
			for _, ignoreSubstr := range ignoreLocationSubstrs {
				if ignoreSubstr != "" && strings.Contains(loc.Data.Name, ignoreSubstr) {
					fmt.Printf("Ignoring location: %s\n", loc.Data.Name)
					ignore = true
					break
				}
			}
			if !ignore {
				filtered = append(filtered, loc)
			}
		} else {
			fmt.Printf("Skipping location: %s Distance: %f\n", loc.Data.Name, distance)
		}
	}

	withinRange := make([]mvc.Appointment, 0)

	for _, loc := range filtered {
		if !loc.NextAvailable.IsZero() {
			if loc.NextAvailable.After(startDate) && loc.NextAvailable.Before(endDate) {
				withinRange = append(withinRange, loc)
			}
		}
	}

	if len(withinRange) == 0 {
		fmt.Println("No available appointments found within criteria")
		return
	}

	message := fmt.Sprintf("Next available appointment at %s is %s", withinRange[0].Data.Name, withinRange[0].NextAvailable.String())
	fmt.Printf("Sending %s to %s\n", message, *destinationNumberFlag)

	err = notifClient.SendNotification(
		*destinationNumberFlag,
		message,
	)

	if err != nil {
		panic(err)
	}
}
