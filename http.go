package main

import (
	"context"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"net/url"
	"regexp"
)

var client *http.Client
var headers http.Header

func initHTTP() {
	// parse proxy URL
	proxyURL, err := url.Parse(*config.proxyURL)
	if err != nil {
		log.Fatal(err)
	}

	// http client
	client = &http.Client{
		Transport: &http.Transport{
			MaxConnsPerHost: *config.parallel,
			Proxy:           http.ProxyURL(proxyURL),
		},
	}

	// headers
	headers = http.Header{"Connection": {"keep-alive"}}
}

func isPaddingError(cipher []byte, ctx *context.Context) (bool, error) {
	// encode the cipher
	cipherEncoded := config.encoder.encode(cipher)
	url, err := url.Parse(fmt.Sprintf(*config.URL, url.QueryEscape(cipherEncoded)))
	if err != nil {
		log.Fatal(err)
	}

	// create request
	req := &http.Request{
		URL:    url,
		Header: headers,
	}

	// add context if passed
	if ctx != nil {
		req = req.WithContext(*ctx)
	}

	// send request
	resp, err := client.Do(req)
	if err != nil {
		return false, err
	}
	defer resp.Body.Close()

	// report about made request
	if currentStatus != nil {
		currentStatus.chanReq <- 1
	}

	// parse the answer
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return false, err
	}

	matched, err := regexp.Match(*config.paddingError, body)
	if err != nil {
		return false, err
	}
	return matched, nil
}
