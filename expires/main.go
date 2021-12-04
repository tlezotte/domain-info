package expires

import (
	"time"

	whoisparser "github.com/likexian/whois-parser"
)

const WarningDays = 90
const ErrorDays = 45

func FormatDate(result whoisparser.WhoisInfo, forDatabase bool) string {
	expires, _ := time.Parse(time.RFC3339, result.Domain.ExpirationDate)

	if forDatabase {
		return expires.Format("2006-01-01")
	} else {
		return expires.Format("Jan 1, 2006")
	}
}

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
