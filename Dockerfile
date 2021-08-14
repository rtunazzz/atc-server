FROM golang:1.16-alpine

# Go env variables
ENV GOFLAGS=-mod=vendor

ENV PORT=9000
ENV SERVER_HOME /go/src/github.com/bonzayio/go-atc-server

# Set the current user to 'server' and set the current working directory
WORKDIR $SERVER_HOME

COPY go.mod .
COPY go.sum .

RUN go mod download
COPY . .

# Build
RUN go build -o atc-server

# Start the server
CMD $SERVER_HOME/atc-server

EXPOSE 9000
