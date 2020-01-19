package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"net/url"
	"regexp"
	"time"
)

var client = &http.Client{}
var headers = http.Header{"Connection": {"keep-alive"}}

func elapsed(what string) func() {
	start := time.Now()
	return func() {
		log.Printf("%s took %v\n", what, time.Since(start))
	}
}

func isPaddingErr(cipher []byte) (bool, error) {

	// encode the cipher
	cipherEncoded := encode(cipher)

	url, err := url.Parse(fmt.Sprintf("http://localhost:5000/decrypt?cipher=%s", url.QueryEscape(cipherEncoded)))
	if err != nil {
		log.Fatal(err)
	}

	req := &http.Request{
		URL: url,
	}

	// send request
	resp, err := client.Do(req)
	if err != nil {
		return false, err
	}
	defer resp.Body.Close()

	// parse the answer
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return false, err
	}

	matched, err := regexp.Match("IncorrectPadding", body)
	if err != nil {
		return false, err
	}

	return matched, nil
}

func isPaddingError(cipher []byte) (bool, error) {
	defer elapsed("http request")
	return isPaddingErr(cipher)

	// encode the cipher
	cipherEncoded := encode(cipher)

	// build the url
	url := fmt.Sprintf("http://localhost:5000/decrypt?cipher=%s", url.QueryEscape(cipherEncoded))

	// send request
	resp, err := http.Get(url)
	if err != nil {
		return false, err
	}
	defer resp.Body.Close()

	// parse the answer
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return false, err
	}

	matched, err := regexp.Match("IncorrectPadding", body)
	if err != nil {
		return false, err
	}

	return matched, nil
}