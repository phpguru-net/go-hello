package timer

import (
	"fmt"
	"strings"
	"time"
)

//const UTCFormat = "2019-07-25 13:45:00 +0000 UTC"
// 7/25/2019 13:45:00
// Schedule returns a time.Time from a string containing a date
func Schedule(date string) time.Time {
	t, _ := time.Parse("1/2/2006 15:04:05", date)
	return t
}

// HasPassed returns whether a date has passed
func HasPassed(date string) bool {
	// July 25, 2019 13:45:00
	appointmentDate, _ := time.Parse("January 2, 2006 15:04:05", date)
	fmt.Println(appointmentDate)
	return time.Now().After(appointmentDate)
}

// IsAfternoonAppointment returns whether a time is in the afternoon
func IsAfternoonAppointment(date string) bool {
	// Thursday, May 13, 2010 20:32:00
	appointmentDate, _ := time.Parse("Monday, January 2, 2006 15:04:05", date)
	fmt.Println(appointmentDate)
	return appointmentDate.Hour() >= 12 && appointmentDate.Hour() < 18
}

// Description returns a formatted string of the appointment time
func Description(date string) string {
	appointmentDate := Schedule(date)
	return appointmentDate.Format("You have an appointment on Monday, January 2, 2006, at 15:04.")
}

// AnniversaryDate returns a Time with this year's anniversary
func AnniversaryDate() time.Time {
	fmt.Println(time.Now().Location())
	date := time.Date(time.Now().UTC().Year(), 9, 15, 0, 0, 0, 0, time.Now().Location())
	return date
}

func UTCToLocal(utcString string) time.Time {
	utcDate, _ := time.Parse("02/01/2006 15:04:05", utcString)
	loc, _ := time.LoadLocation("Asia/Ho_Chi_Minh")
	fmt.Println(utcDate)
	return utcDate.In(loc)
}

func DayLightSaving() {
	// Clock Changes in Washington DC, USA
	// Start DST:	Sunday, 14 March 2021 — 1 hour forward
	// End DST:	Sunday, 7 November 2021 — 1 hour backward

	loc, _ := time.LoadLocation("America/New_York")
	startDST := time.Date(2021, 3, 15, 0, 0, 0, 0, loc)
	endDST := time.Date(2021, 11, 8, 0, 0, 0, 0, loc)
	fmt.Println(strings.Repeat("=", 50))
	fmt.Printf("Timezone: %s\n", "America/New_York")
	fmt.Println(startDST, " ==> ", startDST.UTC())
	fmt.Println(startDST.AddDate(0, 0, -2), " ==> ", startDST.AddDate(0, 0, -1).UTC())
	fmt.Println(endDST, " ==> ", endDST.UTC())
	fmt.Println(endDST.AddDate(0, 0, -2), " ==> ", endDST.AddDate(0, 0, -2).UTC())
	// next year : 13 Mar 2022, 02:00
	fmt.Println(strings.Repeat("=", 50))
	nextYearDST := time.Date(2022, 3, 13, 2, 0, 0, 0, loc)
	fmt.Println(nextYearDST, " ==> ", nextYearDST.UTC())
	fmt.Println(nextYearDST.Add(time.Hour*1), " ==> ", nextYearDST.Add(time.Hour*1).UTC())
}
