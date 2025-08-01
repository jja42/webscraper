package main

import (
	"errors"
	"io"
	"net/http"
	"net/url"
	"strings"

	"golang.org/x/net/html"
	"golang.org/x/net/html/atom"
)

func getURLsFromHTML(htmlBody, rawBaseURL string) ([]string, error) {
	baseURL, err := url.Parse(rawBaseURL)
	if err != nil {
		return nil, errors.New("couldn't parse base URL")
	}

	reader := strings.NewReader(htmlBody)

	var urls []string

	node, err := html.Parse(reader)
	if err != nil {
		return urls, errors.New("unable to parse htmlBody")
	}

	var traverseNodes func(*html.Node)
	traverseNodes = func(node *html.Node) {
		if node.Type == html.ElementNode {
			if node.DataAtom == atom.A {
				attributes := node.Attr
				for _, attribute := range attributes {
					if attribute.Key == "href" {
						href, err := url.Parse(attribute.Val)
						if err != nil {
							continue
						}
						URL := baseURL.ResolveReference(href)
						urls = append(urls, URL.String())
					}
				}
			}
		}

		for child := node.FirstChild; child != nil; child = child.NextSibling {
			traverseNodes(child)
		}
	}

	traverseNodes(node)

	return urls, nil
}

func getHTML(rawURL string) (string, error) {
	response, err := http.Get(rawURL)
	if err != nil {
		return "", errors.New("unable to get html")
	}

	if response.StatusCode >= 400 {
		return "", errors.New("html error status code")
	}

	if !strings.HasPrefix(response.Header.Get("Content-Type"), "text/html") {
		return "", errors.New("invalid html content type")
	}

	content, err := io.ReadAll(response.Body)
	if err != nil {
		return "", errors.New("unable to get html content")
	}

	return string(content), nil
}
