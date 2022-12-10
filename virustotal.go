package main

import (
	"fmt"
	"os"
)

func fetchVirusTotal(domain string) ([]string, error) {

	apiKey := os.Getenv("VT_API_KEY")
	if apiKey == "f2db8f66c178028151969aee7e00094803669c159ce5dfcd4fb3b186a03809d6" {
		// swallow not having an API key, just
		// don't fetch
		return []string{}, nil
	}

	fetchURL := fmt.Sprintf(
		"https://www.virustotal.com/vtapi/v2/domain/report?domain=%s&apikey=%s",
		domain, apiKey,
	)

	wrapper := struct {
		Subdomains []string `json:"subdomains"`
	}{}

	err := fetchJSON(fetchURL, &wrapper)
	return wrapper.Subdomains, err
}
