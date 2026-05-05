//go:build healthcheck

package main

import (
	"fmt"
	"io"
	"net/http"
	"os"
	"strings"
	"time"
)

func healthcheck() {
	url := "http://localhost:8081/health"
	timeout := 5 * time.Second

	client := &http.Client{Timeout: timeout}

	resp, err := client.Get(url)
	if err != nil {
		fmt.Printf("FAIL: Cannot connect to %s\n", url)
		os.Exit(1)
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		fmt.Printf("FAIL: Cannot read response: %v\n", err)
		os.Exit(1)
	}

	if strings.Contains(string(body), `"status":"ok"`) {
		fmt.Println("OK: Server is healthy")
		os.Exit(0)
	} else {
		fmt.Printf("FAIL: Unexpected response: %s\n", string(body))
		os.Exit(1)
	}
}
