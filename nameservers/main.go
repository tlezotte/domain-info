package nameservers

import (
	"strings"

	whoisparser "github.com/likexian/whois-parser"
)

func GetName(result whoisparser.WhoisInfo) string{
	return result.Domain.NameServers[0]
}

func IsVUMC(result whoisparser.WhoisInfo) bool {
	return strings.Contains(result.Domain.NameServers[0], "vumc.org")
}
