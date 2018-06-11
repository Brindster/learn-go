FROM golang:1.10-stretch as builder

# Prepare the app directory
WORKDIR /go/src/chrisbrindley.co.uk
ADD . .

# Ensure dependancies are installed
RUN go get github.com/golang/dep/cmd/dep
RUN dep ensure --vendor-only

# Prepare build
RUN go build -o app

FROM debian:stretch-slim
WORKDIR /root
ADD . .
COPY --from=builder /go/src/chrisbrindley.co.uk/app .
CMD ["./app"]