package main

import (
	"io/ioutil"
	"net/http"
	"os"
)

func main() {
	resp, err := http.Get("http://127.0.0.1:80/healtcheck")
	if err == nil {
		body, err := ioutil.ReadAll(resp.Body)
		if err == nil {
			if string(body) == "ok" {
				os.Exit(0)
				return
			}
		}
	}
	os.Exit(1)
}
