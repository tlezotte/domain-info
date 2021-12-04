// Package is used to get expiration date from whois record
package expires

import (
	"time"

	whoisparser "github.com/likexian/whois-parser"
)

const (
	// Days before expiration to start warning
	WarningDays = 90
	// Days before expiration to start error
	ErrorDays = 45
)

// How to display the expiration date
func FormatDate(result whoisparser.WhoisInfo, forDatabase bool) string {
	expires, _ := time.Parse(time.RFC3339, result.Domain.ExpirationDate)

	if forDatabase {
		return expires.Format("2006-01-01")
	} else {
		return expires.Format("Jan 1, 2006")
	}
}

// Days left before expiration
func DiffExpiration(result whoisparser.WhoisInfo) (int, string) {
	highlight := "bold"
	current_date := time.Now()
	expires, _ := time.Parse(time.RFC3339, result.Domain.ExpirationDate)

	diff := expires.Sub(current_date)

	days := int(diff.Hours() / 24)

	switch {
	case days <= WarningDays:
		highlight = "warning"
	case days <= ErrorDays:
		highlight = "error"
	}

	return days, highlight
}
