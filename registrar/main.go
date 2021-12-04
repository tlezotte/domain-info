package registrar

import (
	"strings"

	whoisparser "github.com/likexian/whois-parser"
)

func GetName(result whoisparser.WhoisInfo) string{
	return result.Registrar.Name
}

func IsVUMC(result whoisparser.WhoisInfo) bool {
	return strings.Contains(result.Registrar.Name, "Network Solutions")
}
