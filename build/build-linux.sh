env GOOS=linux GOARCH=amd64 go build -o bin/linux/sessionServer session/cmd/session.go
env GOOS=linux GOARCH=amd64 go build -o bin/linux/redirectServer redirect/cmd/redirect.go
env GOOS=linux GOARCH=amd64 go build -o bin/linux/infra2Server infra/cmd/infra2.go

