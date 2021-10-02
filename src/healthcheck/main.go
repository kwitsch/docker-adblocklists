package main

import (
	"fmt"
	"net/http"
	"os"
)

func main() {
	_, err := http.Get("http://127.0.0.1:80/healtcheck")
	if err != nil {
		fmt.Println("Healthchek fail")
		os.Exit(1)
	} else {
		fmt.Println("Healthchek ok")
		os.Exit(0)
	}
}
