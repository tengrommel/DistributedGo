#!/usr/bin/env bash
GOOS=linux GOARCH=amd64 go build -o main
zip main.zip main