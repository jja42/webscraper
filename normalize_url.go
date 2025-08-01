package main

import (
	"errors"
	"net/url"
	"strings"
)

func normalizeURL(rawURL string) (string, error) {
	url_struct, err := url.Parse(rawURL)
	if err != nil {
		return "", errors.New("unable to parse url")
	}

	domain := url_struct.Hostname()
	path := url_struct.Path

	normalizedURL := domain + path

	normalizedURL = strings.ToLower(normalizedURL)

	normalizedURL = strings.TrimSuffix(normalizedURL, "/")

	return normalizedURL, nil
}
