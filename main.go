package main

import (
	"fmt"
	"net/http"
	"os"
	"time"
)

var startTime = time.Now()
var version = "1.0.0"

func main() {
	hostname, _ := os.Hostname()

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "text/html; charset=utf-8")
		fmt.Fprintf(w, `<!DOCTYPE html>
<html>
<head>
    <meta charset="UTF-8">
    <title>Go Demo - Jenkins Learning</title>
    <style>
        body {
            font-family: -apple-system, BlinkMacSystemFont, 'Segoe UI', Roboto, sans-serif;
            background: linear-gradient(135deg, #1a1a2e 0%%, #16213e 100%%);
            color: #fff;
            display: flex;
            justify-content: center;
            align-items: center;
            min-height: 100vh;
            margin: 0;
        }
        .container {
            text-align: center;
            padding: 2rem;
            background: rgba(255,255,255,0.05);
            border-radius: 16px;
            backdrop-filter: blur(10px);
            box-shadow: 0 8px 32px rgba(0,0,0,0.3);
        }
        h1 { color: #00d4aa; margin-bottom: 1rem; }
        .info { 
            margin: 0.5rem 0; 
            color: #a0a0a0;
            font-size: 0.95rem;
        }
        .label { color: #888; }
        .value { color: #fff; font-weight: 500; }
        .version { 
            margin-top: 1.5rem; 
            padding-top: 1rem; 
            border-top: 1px solid rgba(255,255,255,0.1);
            font-size: 0.85rem;
            color: #666;
        }
    </style>
</head>
<body>
    <div class="container">
        <h1>Go Demo - Jenkins Learning</h1>
        <div class="info">
            <span class="label">Hostname:</span> 
            <span class="value">%s</span>
        </div>
        <div class="info">
            <span class="label">Server Started:</span> 
            <span class="value">%s</span>
        </div>
        <div class="info">
            <span class="label">Uptime:</span> 
            <span class="value">%s</span>
        </div>
        <div class="version">Version %s</div>
    </div>
</body>
</html>`, hostname, startTime.Format("2006-01-02 15:04:05"), time.Since(startTime).Round(time.Second), version)
	})

	http.HandleFunc("/health", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		fmt.Fprint(w, `{"status":"ok"}`)
	})

	port := ":8081"
	fmt.Printf("Server starting on port %s\n", port)
	if err := http.ListenAndServe(port, nil); err != nil {
		fmt.Fprintf(os.Stderr, "Server error: %v\n", err)
		os.Exit(1)
	}
}
