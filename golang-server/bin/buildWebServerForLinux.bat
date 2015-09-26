@echo off
set GOOS=linux
set ARCH=amd64
set CGO_ENABLE=0

go build -a -ldflags "-s -w" ../src/webserver.go
pause