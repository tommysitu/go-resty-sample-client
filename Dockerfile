FROM golang:1.13-alpine
WORKDIR /accountapi
COPY go.mod .
COPY go.sum .
RUN go mod download
COPY . .
CMD CGO_ENABLED=0 GOOS=linux go test -v ./...
