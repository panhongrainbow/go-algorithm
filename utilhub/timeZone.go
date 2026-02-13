package utilhub

import (
	"io/ioutil"
	"path/filepath"
	"time"
)

// ListTimezones returns all available timezones in the zoneinfo directory.
func ListTimezones(timezoneDir string) ([]string, error) {
	var timezones []string

	files, err := ioutil.ReadDir(timezoneDir)
	if err != nil {
		return nil, err
	}

	for _, file := range files {
		if file.IsDir() {
			zones, err := filepath.Glob(timezoneDir + file.Name() + "/*")
			if err != nil {
				return nil, err
			}
			for _, zone := range zones {
				timezones = append(timezones, file.Name()+"/"+filepath.Base(zone))
			}
		} else {
			timezones = append(timezones, file.Name())
		}
	}
	return timezones, nil
}

// GetTimeString to get formatted time or date in a specified time zone.

// Here are 5 commonly used time format :
//   year-month-day : 2006-01-02
//   YYYY-MM-DD HH:MM:SS : 2006-01-02 15:04:05

// Here are 5 commonly used time zones :
//   UTC (Coordinated Universal Time), Identifier: "UTC"
//   Asia/Shanghai (China Standard Time - CST), Identifier: "Asia/Shanghai"
//   America/New_York (Eastern Time - ET, US), Identifier: "America/New_York"
//   Europe/London (Greenwich Mean Time - GMT / British Summer Time - BST), Identifier: "Europe/London"
//   Asia/Tokyo (Japan Standard Time - JST), Identifier: "Asia/Tokyo"

func GetNowTimeString(format string, timeZone string) (string, error) {
	// Load the specified time zone
	location, err := time.LoadLocation(timeZone)
	if err != nil {
		return "", err
	}

	// Get the current time and convert it to the specified time zone
	currentTime := time.Now().In(location)

	// Format the date string
	dateStr := currentTime.Format(format)
	return dateStr, nil
}
