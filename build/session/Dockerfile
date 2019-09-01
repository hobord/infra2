
# build stage
FROM golang:1.12-alpine AS build-env
RUN apk add --no-cache git mercurial

WORKDIR /go/app
COPY . .

RUN go get -d -v ./...
RUN go build -o build/session/sessionServer session/cmd/session.go

# final stage
FROM alpine
WORKDIR /app/
COPY --from=build-env /go/app/build/session/sessionServer /app/

EXPOSE 50051
ENTRYPOINT /app/sessionServer
