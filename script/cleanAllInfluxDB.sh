#!/usr/bin/env bash
curl --request POST http://ip/api/v2/delete\?org\=qz\&bucket\=qz \
  --header 'Authorization: Token token' \
  --header 'Content-Type: application/json' \
  --data '{
    "start": "2022-03-01T00:00:00Z",
    "stop": "2022-11-14T00:00:00Z"
  }'