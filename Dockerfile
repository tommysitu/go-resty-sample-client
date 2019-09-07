FROM golang:1.13-alpine
WORKDIR /usr/local/accountapi
COPY . /usr/local/accountapi
CMD CGO_ENABLED=0 GOOS=linux go test -v ./...
