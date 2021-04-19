#!/bin/bash
# Send echo in every second

while true; do
  grpcurl -d '{"content": "echo"}' -proto echo-grpc/api/echo.proto -plaintext -v echo-grpc:8081 api.Echo/Echo
  sleep 1
done
