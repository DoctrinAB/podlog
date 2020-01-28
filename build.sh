#!/bin/sh

VERSION=`cat version`

for OS in darwin linux; do
	echo "Start build for $OS"
	GOOS=$OS GOARCH=amd64 go build -ldflags "-X main.Version=$VERSION" -o bin/$OS/podlog .
done
