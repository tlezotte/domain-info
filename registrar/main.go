// Package is used to get registrar's name from whois record
package registrar

import (
	"strings"

	whoisparser "github.com/likexian/whois-parser"
)

// Get the registrar's name
func GetName(result whoisparser.WhoisInfo) string {
	return result.Registrar.Name
}

// Are we registered at VUMC?
func IsVUMC(result whoisparser.WhoisInfo) bool {
	return strings.Contains(result.Registrar.Name, "Network Solutions")
}
