package http_check

import (
	"net/http"
	"crypto/tls"
	"time"
	"log"
)

func TestHttpConnection(service string, timeout int) (status bool) {
	client := http.Client{
		Timeout: time.Duration(timeout) * time.Second,
	}
	log.Println("HTTP Connect to : ", service)
	resp, err := client.Get("http://" + service)
	if err != nil {
		status=false
		return status
	}
	if (resp.StatusCode >= 200) && (resp.StatusCode <= 399) {
		status = true
	} else {
		status = false
	}
	return status
}

func TestHttpsConnection(service string, timeout int) (status bool) {
	http.DefaultTransport.(*http.Transport).TLSClientConfig = &tls.Config{InsecureSkipVerify: true}
	client := http.Client{
		Timeout: time.Duration(timeout) * time.Second,
	}
	log.Println("HTTPS Connect to : ", service)
	resp, err := client.Get("https://" + service)
	if err != nil {
		status=false
		return status
	}
	if (resp.StatusCode >= 200) && (resp.StatusCode <= 399) {
		status = true
	} else {
		status = false
	}
	return status
}
