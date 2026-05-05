#!/bin/bash
set -e

URL="http://localhost:8081/health"
TIMEOUT=5

response=$(curl -sf --max-time $TIMEOUT $URL 2>&1)
curl_exit=$?

if [ $curl_exit -ne 0 ]; then
    echo "FAIL: Cannot connect to $URL"
    exit 1
fi

if echo "$response" | grep -q '"status":"ok"'; then
    echo "OK: Server is healthy"
    exit 0
else
    echo "FAIL: Unexpected response: $response"
    exit 1
fi