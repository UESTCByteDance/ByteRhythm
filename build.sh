#!/bin/bash
go build -o gateway app/gateway/cmd/main.go
go build -o user app/user/cmd/main.go
go build -o video app/video/cmd/main.go
go build -o favorite app/favorite/cmd/main.go
go build -o comment app/comment/cmd/main.go
go build -o relation app/relation/cmd/main.go
go build -o message app/message/cmd/main.go
