env GOOS=linux GOARCH=amd64 go build -o bin/linux/sessionServer cmd/session.go
env GOOS=linux GOARCH=amd64 go build -o bin/linux/redirectServer cmd/redirect.go
env GOOS=linux GOARCH=amd64 go build -o bin/linux/infra2Server cmd/infra2.go

