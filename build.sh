#!/bin/sh
GOOS=linux GOARCH=386 go build -o builds/gradle-update-notifier_linux_386
GOOS=linux GOARCH=amd64 go build -o builds/gradle-update-notifier_linux_amd64
GOOS=windows GOARCH=386 go build -o builds/gradle-update-notifier_windows_386
GOOS=windows GOARCH=amd64 go build -o builds/gradle-update-notifier_windows_amd64
GOOS=darwin GOARCH=386 go build -o builds/gradle-update-notifier_darwin_386
GOOS=darwin GOARCH=amd64 go build -o builds/gradle-update-notifier_darwin_386_amd64
