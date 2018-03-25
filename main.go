package main

import (
	"fmt"
	"net/http"
	//	"net/url"
	"bytes"
	"crypto/hmac"
	"crypto/sha512"
	"encoding/hex"
	"io/ioutil"
	"strconv"
	"time"
)

func sendQuery(url string, query string, apiKey string, apiSecretKey string) {

	t := time.Now().Unix()
	tString := strconv.FormatInt(t, 10)

	var buffer bytes.Buffer
	buffer.WriteString(tString)
	buffer.WriteString(query)

	key := []byte(apiSecretKey)
	h := hmac.New(sha512.New, key)
	h.Write([]byte(buffer.String()))
	sig := hex.EncodeToString(h.Sum(nil))

	req, err := http.NewRequest("POST", url, bytes.NewBuffer([]byte(query)))
	if err != nil {
		return
	}

	client := &http.Client{}
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("x-orionx-timestamp", tString)
	req.Header.Set("x-orionx-apikey", apiKey)
	req.Header.Set("x-orionx-signature", sig)
	req.Header.Set("Content-Length", strconv.Itoa(len(query)))

	fmt.Println(req)
	resp, _ := client.Do(req)
	fmt.Println("")
	fmt.Println(resp)
	fmt.Println("")
	b, _ := ioutil.ReadAll(resp.Body)
	fmt.Println(string(b))
	return
}

func main() {

	apiKey := "BHsHEDtGys2vcRfooz7XxZfjEyK6yWBLnd"

	secKey := ""
	url := "https://api2.orionx.io/graphql"
	query := "{\"query\":\"{me{_id email}}\",\"variables\":null,\"operationName\":null}"

	sendQuery(url, query, apiKey, secKey)

}
