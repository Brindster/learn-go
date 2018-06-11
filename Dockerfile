FROM golang:1.10-stretch as builder

# Prepare the app directory
WORKDIR /go/src/chrisbrindley.co.uk
ADD . .

# Ensure dependancies are installed
RUN go get github.com/golang/dep/cmd/dep
RUN dep ensure --vendor-only

CMD ["go", "run", "main.go"]
