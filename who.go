package main

import (
	"fmt"
	"os"
	"time"

	"github.com/gookit/color"
	"github.com/likexian/whois"
	whoisparser "github.com/likexian/whois-parser"
)

var red = color.FgRed.Render
//var green = color.FgGreen.Render
//var yellow = color.FgYellow.Render

func main() {
	if len(os.Args) >= 1 {
		whois_raw := whoisRaw(os.Args[1])

		result, err := whoisparser.Parse(whois_raw)
		if err == nil {
			// Print the domain expiration date
			expireDate := formatTime(result.Domain.ExpirationDate, true)
			fmt.Printf("This domain expires on %s\n", red(expireDate))

			// Print the registrar name
			fmt.Println(result.Registrar.Name)

			// Print the Name Servers
			fmt.Println(result.Domain.NameServers)
		} else {
			fmt.Println(err)
		}
	} else {
		fmt.Println("Usage: whois <domain>")
	}
}

func whoisRaw(domain string) string {
	whois_raw, err := whois.Whois(domain)
	if err == nil {
		return whois_raw
	} else {
		return err.Error()
	}
}

func formatTime(expirationDate string, forDatabase bool) string {
	t1, _ := time.Parse(time.RFC3339, expirationDate)

	if forDatabase {
		return t1.Format("2006-01-02")
	} else {
		return t1.Format("Jan 2, 2006")	
	}
}
