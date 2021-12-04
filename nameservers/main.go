// Package is used to get name server date from whois record
package nameservers

import (
	"strings"

	whoisparser "github.com/likexian/whois-parser"
)

// GetName returns name a server
func GetName(result whoisparser.WhoisInfo) string {
	return result.Domain.NameServers[0]
}

// Are nameservers at VUMC?
func IsVUMC(result whoisparser.WhoisInfo) bool {
	return strings.Contains(result.Domain.NameServers[0], "vumc.org")
}
