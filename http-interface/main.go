package main

import (
	"fmt"
	"io"
	"net/http"
	"os"
)

func main() {
	resp, err := http.Get("https://google.com")
	if err != nil {
		fmt.Println("Error from Get request", err)
		os.Exit(1)
	}

	io.Copy(os.Stdout, resp.Body)
}
