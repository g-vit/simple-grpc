
FROM golang:1.12-alpine
RUN apk add --no-cache curl git

WORKDIR $GOPATH/src/github.com/g-vit/simple-grpc

COPY . .

# Server
WORKDIR $GOPATH/src/github.com/g-vit/simple-grpc/server
RUN CGO_ENABLED=0 GOOS=linux GO111MODULE=on go build -a -o /server .

# Client
WORKDIR $GOPATH/src/github.com/g-vit/simple-grpc/client
RUN CGO_ENABLED=0 GOOS=linux GO111MODULE=on go build -a -o /client .

FROM alpine:3.8
RUN apk add --no-cache ca-certificates
COPY --from=0 /server /client /usr/bin/
WORKDIR /
