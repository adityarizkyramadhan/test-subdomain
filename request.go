package main

import (
	"context"
	"net/http"
	"time"
)

func RequestHttp(domain string, client *http.Client) (int, error, bool, string) {
	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()
	req, err := http.NewRequest("GET", "http://"+domain, nil)
	if err != nil {
		return 0, err, false, ""
	}
	req = req.WithContext(ctx)
	resp, err := client.Do(req)
	if err != nil {
		return 0, err, false, ""
	}
	isWordPress := CheckIsWordpress(req.URL.String(), client)
	if isWordPress {
		return resp.StatusCode, nil, true, req.URL.String()
	}
	return resp.StatusCode, nil, false, req.URL.String()
}

func RequestHttps(domain string, client *http.Client) (int, error, bool, string) {
	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()
	req, err := http.NewRequest("GET", "https://"+domain, nil)
	if err != nil {
		return 0, err, false, ""
	}
	req = req.WithContext(ctx)
	resp, err := client.Do(req)
	if err != nil {
		return 0, err, false, ""
	}
	isWordPress := CheckIsWordpress(req.URL.String(), client)
	if isWordPress {
		return resp.StatusCode, nil, true, req.URL.String()
	}
	return resp.StatusCode, nil, false, req.URL.String()
}
func CheckIsWordpress(link string, client *http.Client) bool {
	req, err := http.NewRequest("GET", link+"/wp-login.php", nil)
	// defer error only return bool
	if err != nil {
		return false
	}
	resp, err := client.Do(req)
	if err != nil {
		return false
	}
	if resp.StatusCode == 200 {
		return true
	} else {
		return false
	}
}
