package utils

import (
	"strings"
	"time"
)

type DateFromat string

const (
	// Date

	DateISO         DateFromat = "2006-01-02"              // 2026-07-22
	DateShort       DateFromat = "02 Jan 2006"             // 22 Jul 2026
	DateLong        DateFromat = "January 2, 2006"         // July 22, 2026
	DateLongOrdinal DateFromat = "Monday, January 2, 2006" // Wednesday, July 22, 2026

	// Time

	Time24        DateFromat = "15:04"      // 13:45
	Time24Seconds DateFromat = "15:04:05"   // 13:45:30
	Time12        DateFromat = "3:04 PM"    // 1:45 PM
	Time12Seconds DateFromat = "3:04:05 PM" // 1:45:30 PM

	// Date & Time

	DateTimeISO             DateFromat = "2006-01-02 15:04:05"
	DateTimeISOWithMillis   DateFromat = "2006-01-02 15:04:05.000"
	DateTimeISOWithTimezone DateFromat = "2006-01-02T15:04:05Z07:00"

	DateTimeShort24 DateFromat = "02 Jan 2006 15:04"
	DateTimeShort12 DateFromat = "02 Jan 2006 3:04 PM"

	DateTimeLong24 DateFromat = "Monday, January 2, 2006 at 15:04"
	DateTimeLong12 DateFromat = "Monday, January 2, 2006 at 3:04 PM"

	DateTimeLong24Seconds DateFromat = "Monday, January 2, 2006 at 15:04:05"
	DateTimeLong12Seconds DateFromat = "Monday, January 2, 2006 at 3:04:05 PM"

	// File names

	FileDate      DateFromat = "2006-01-02"
	FileDateTime  DateFromat = "2006-01-02_15-04-05"
	FileTimestamp DateFromat = "20060102_150405"

	// Logs

	LogTimestamp DateFromat = "2006-01-02 15:04:05.000"

	// Built-in RFC formats

	RFC3339     DateFromat = "2006-01-02T15:04:05Z07:00"
	RFC3339Nano DateFromat = "2006-01-02T15:04:05.999999999Z07:00"
	RFC822      DateFromat = "02 Jan 06 15:04 MST"
	RFC850      DateFromat = "Monday, 02-Jan-06 15:04:05 MST"
	RFC1123     DateFromat = "Mon, 02 Jan 2006 15:04:05 MST"
	ANSIC       DateFromat = "Mon Jan _2 15:04:05 2006"
	UnixDate    DateFromat = "Mon Jan _2 15:04:05 MST 2006"
)

func HumanTime(t time.Time, format ...DateFromat) string {
	f := DateTimeLong12
	if len(format) > 0 {
		f = format[0]
	}
	return t.Format(string(f))
}

// ParseTimestamp parses the common string forms that come back from
// PostgreSQL drivers (RFC3339 with/without nanos and zone, ISO date-time
// with a space separator) into a time.Time. Returns the zero time on
// empty input.
func ParseTimestamp(s string) (time.Time, error) {
	s = strings.TrimSpace(s)
	if s == "" {
		return time.Time{}, nil
	}
	layouts := []string{
		time.RFC3339Nano,
		time.RFC3339,
		string(RFC3339),
		"2006-01-02 15:04:05.999999",
		"2006-01-02 15:04:05.999999999",
		"2006-01-02 15:04:05",
		"2006-01-02T15:04:05",
		string(DateISO),
	}
	for _, layout := range layouts {
		if t, err := time.Parse(layout, s); err == nil {
			return t, nil
		}
	}
	return time.Time{}, &time.ParseError{Value: s, Message: "unrecognised timestamp format"}
}
