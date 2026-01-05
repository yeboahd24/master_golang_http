package main

import (
	"fmt"
	"time"

	"github.com/beevik/ntp"
)

func GetNetworkTime() (time.Time, error) {
	ntpTime, err := ntp.Time("pool.ntp.org")
	if err != nil {
		return time.Time{}, err
	}
	return ntpTime, nil
}

// func main() {
// 	currentTime, err := GetNetworkTime()
// 	if err != nil {
// 		log.Fatalf("Error getting network time: %v", err)
// 	}
// 	fmt.Println("Current Network Time:", currentTime.UTC())
// }

// Server side
func main() {
	// Assume subscription expiry time is stored in UTC
	expiryTimeUTC := time.Date(2024, 12, 31, 23, 59, 59, 0, time.UTC)

	// Get current server time in UTC
	currentTimeUTC := time.Now().UTC()

	if currentTimeUTC.After(expiryTimeUTC) {
		fmt.Println("Subscription has expired.")
	} else {
		fmt.Println("Subscription is still active.")
	}
}
