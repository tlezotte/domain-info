package main

import (
	"fmt"
	"os"
	"path/filepath"
	"time"

	"github.com/gookit/color"
	"github.com/likexian/whois"
	whoisparser "github.com/likexian/whois-parser"
	"github.com/tlezotte/domain-info/expires"
	"github.com/tlezotte/domain-info/nameservers"
	"github.com/tlezotte/domain-info/registrar"
)


var program_name = filepath.Base(os.Args[0])
var warning_days = 90
var error_days = 45
var bold = color.Bold.Render
var error = color.FgRed.Render
var warning = color.FgYellow.Render
//var bold = color.New(color.OpBold)
//var warning = color.New(color.FgYellow, color.OpBold)
//var error = color.New(color.FgRed, color.OpBold)

func main() {
	fmt.Println(nameservers.Hello("tom"))
	fmt.Println(expires.Hello())
	fmt.Println(registrar.Hello())
	
	if len(os.Args) == 2 {
		domain := os.Args[1]
		whois_raw := whoisRaw(os.Args[1])

		result, err := whoisparser.Parse(whois_raw)
		if err == nil {
			expireDate := formatDate(result.Domain.ExpirationDate, false)
			registrar := result.Registrar.Name
			ns := result.Domain.NameServers
			diffDays, highlight := diffExpiration(result.Domain.ExpirationDate)

			fmt.Printf("Domain: %s\n", bold(domain))
			fmt.Printf("Expires: %s (%d)\n", bold(expireDate), diffDays)
			fmt.Printf("Status: %s\n", warning(highlight))
			fmt.Printf("Registrar: %s\n", bold(registrar))
			fmt.Printf("Name Servers: %s\n", bold(ns))
			fmt.Println()
		} else {
			fmt.Println(err)
		}
	} else {
		print_usage()
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

func formatDate(expirationDate string, forDatabase bool) string {
	expires, _ := time.Parse(time.RFC3339, expirationDate)

	if forDatabase {
		return expires.Format("2006-01-01")
	} else {
		return expires.Format("Jan 1, 2006")
	}
}

func diffExpiration(expirationDate string) (int, string) {
	highlight := "bold"
	current_date := time.Now()
	expires, _ := time.Parse(time.RFC3339, expirationDate)
	
	diff := expires.Sub(current_date)
	
	days := int(diff.Hours() / 24)

	switch {
	case days <= warning_days:
		highlight = "warning"
	case days <= error_days:
		highlight = "error"
	}

	return days, highlight
}

func print_usage() {
	usage := program_name + " <domain>"
	fmt.Printf("Usage: %s\n\n", error(usage))
}
