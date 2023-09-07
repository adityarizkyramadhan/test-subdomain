package main

import (
	"io"
	"log"
	"net/http"
	"os"
	"sync"
)

func main() {
	subdomains := ReadFile("subdomainlist.csv")
	wg := new(sync.WaitGroup)
	wg.Add(2)
	clientHttp := new(http.Client)
	go MakeRequestHttp(subdomains, clientHttp, wg)
	clientHttps := new(http.Client)
	go MakeRequestHttps(subdomains, clientHttps, wg)
	wg.Wait()
}

func MakeRequestHttp(subdomains []string, client *http.Client, wg *sync.WaitGroup) {
	logFile, err := os.Create("httprequest.log")
	if err != nil {
		log.Fatal(err)
	}
	defer func(logFile *os.File) {
		err := logFile.Close()
		panicIfError(err)
	}(logFile)
	logger := log.New(io.MultiWriter(os.Stdout, logFile), "", log.Ldate|log.Ltime)
	for _, subdomain := range subdomains {
		statusCodeHttp, errHttp, isWordPressHttp, url := RequestHttp(subdomain, client)
		if errHttp != nil {
			logger.Printf("ERROR = %s\n", errHttp.Error())
			continue
		}
		logger.Printf("SUCCESS = URL : %s; status code : %d; wordpress : %t\n", url, statusCodeHttp, isWordPressHttp)
	}
	wg.Done()
}
func MakeRequestHttps(subdomains []string, client *http.Client, wg *sync.WaitGroup) {
	logFile, err := os.Create("httpsrequest.log")
	if err != nil {
		log.Fatal(err)
	}
	defer func(logFile *os.File) {
		err := logFile.Close()
		panicIfError(err)
	}(logFile)
	logger := log.New(io.MultiWriter(os.Stdout, logFile), "", log.Ldate|log.Ltime)
	for _, subdomain := range subdomains {
		statusCodeHttp, errHttp, isWordPressHttp, url := RequestHttps(subdomain, client)
		if errHttp != nil {
			logger.Printf("ERROR = %s\n", errHttp.Error())
			continue
		}
		logger.Printf("SUCCESS = URL : %s; status code : %d; wordpress : %t\n", url, statusCodeHttp, isWordPressHttp)
	}
	wg.Done()
}
