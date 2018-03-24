package main

import (
	"fmt"
	"net/http"
	"net/url"
	"strconv"
	"strings"
)

func sendQuery(url string, query string, apiKey string, apiSecretKey string) {

	req, err := http.NewRequest("POST", url, strings.NewReader(query))
	if err != nil {
		return
	}

	client := &http.Client{}

	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("X-ORIONX-TIMESTAMP", "application/json")
	req.Header.Set("X-ORIONX-APIKEY", apiKey)
	req.Header.Set("X-ORIONX-SIGNATURE", "application/json")
	req.Header.Set("Content-Length", strconv.Itoa(len(query)))

	resp, _ := client.Do(req)
	fmt.Println(resp)

	return
}

func main() {

}
