# DMV New Appointment Notifications

Scheduling DMV appointments can be hard, especially if you didn't plan weeks in advance. This tool helps by automatically checking for available appointments at New Jersey Motor Vehicle Commission (MVC) locations and notifying you via a phone call when an appointment becomes available that meets your criteria.

## Features

- Fetches next available appointments from the NJ MVC portal (TeleGov)
- Filters appointments by:
  - Distance from a center point (lat/lon)
  - Location names (ignore list)
  - Date range
- Sends voice call via Twilio when appointments are found

## Usage

```bash
go run main.go \
    -url "https://telegov.njportal.com/njmvc/AppointmentWizard/12" \
    -start "01/15/2024" \
    -end "02/01/2024" \
    -to "+1234567890" \
    -ignore "Newark,Trenton" \
    -max-distance 50 \
    -lat 40.216 \
    -lon -74.815
```

## Environment Variables

- `TWILIO_ACCOUNT_SID`
- `TWILIO_AUTH_TOKEN`
- `TWILIO_ORIGIN_NUMBER`

## Architecture

Uses a modular design with scraper and notification clients to facilitate testing and extensibility.
