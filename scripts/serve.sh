#!/bin/bash 

go run cmd/server/server.go &

echo "Monitoring Changes"
while inotifywait -q -q -e modify,create,move,delete -r ./; do
    ecshift_api_client_contextho "Changes Detected" && \
    echo "Restarting Server" && \
    pkill go
    go run cmd/server/server.go &
done
