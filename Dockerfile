FROM alpine:latest as deps

# getting certs
RUN apk update && apk upgrade && \
    apk add --no-cache ca-certificates curl

FROM golang:latest as build

# set work dir
WORKDIR /app

# copy the source files
COPY . .

# disable crosscompiling 
ENV CGO_ENABLED=0

# compile linux only
ENV GOOS=linux

# build the binary with debug information removed
RUN go build -ldflags '-w -s' -a -installsuffix cgo -o server

FROM alpine:latest

# copy our static linked library
COPY --from=build /app/server .

# copy certs
COPY --from=deps /etc/ssl/certs/ca-certificates.crt /etc/ssl/certs/

# tell we are exposing our service on port 8081 50051
EXPOSE 8081 50051

# run it!
CMD ["./server", "serve"]