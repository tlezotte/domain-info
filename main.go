package main

import (
	"fmt"
	"os"
	"path/filepath"

	"github.com/gookit/color"
	"github.com/likexian/whois"
	whoisparser "github.com/likexian/whois-parser"

	"github.com/tlezotte/domain-info/expires"
	"github.com/tlezotte/domain-info/nameservers"
	"github.com/tlezotte/domain-info/registrar"
)


var programName = filepath.Base(os.Args[0])
var bold = color.Bold.Render
var error = color.FgRed.Render
var warning = color.FgYellow.Render
//var bold = color.New(color.OpBold)
//var warning = color.New(color.FgYellow, color.OpBold)
//var error = color.New(color.FgRed, color.OpBold)

type WhoisData struct {
	Domain string
	Expires string
	DiffDays int
	Highlight string
	Registrar string
	VumcRegistrar bool
	NameServers string
	VumcNS bool
}

func main() {
	if len(os.Args) == 2 {
		domainName := os.Args[1]

		whoisRaw := getWhois(os.Args[1])
		whoisParsed, errParsed := whoisparser.Parse(whoisRaw)
		if errParsed == nil {
			diffDays, highlight := expires.DiffExpiration(whoisParsed)

			whoisResult := WhoisData{
				Domain: domainName,
				Expires: expires.FormatDate(whoisParsed, false),
				DiffDays: diffDays,
				Highlight: highlight,
				Registrar: registrar.GetName(whoisParsed),
				VumcRegistrar: registrar.IsVUMC(whoisParsed),
				NameServers: nameservers.GetName(whoisParsed),
				VumcNS: nameservers.IsVUMC(whoisParsed),
			}

			outputData(whoisResult)
		} else {
			fmt.Println(errParsed)
		}
	} else {
		printUsage()
	}
}


func getWhois(domain string) string {
	whoisRaw, err := whois.Whois(domain)
	if err == nil {
		return whoisRaw
	} else {
		return err.Error()
	}
}

func outputData(w WhoisData) {
	fmt.Printf("Domain: %s\n", bold(w.Domain))
	fmt.Printf("Expires: %s (%d)\n", bold(w.Expires), w.DiffDays)
	fmt.Printf("Status: %s\n", warning(w.Highlight))
	fmt.Printf("Registrar: %s\n", bold(w.Registrar))
	fmt.Printf("isVUMC: %s\n", bold(w.VumcRegistrar))
	fmt.Printf("Name Servers: %s\n", bold(w.NameServers))
	fmt.Printf("isVUMC: %s\n", bold(w.VumcNS))
	fmt.Println()
}

func printUsage() {
	usage := programName + " <domain>"
	fmt.Printf("Usage: %s\n\n", error(usage))
}
