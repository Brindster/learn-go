FROM golang:1.10-alpine as dev

# Prepare the app directory
WORKDIR /go/src/chrisbrindley.co.uk

# Ensure dependancies are installed
RUN apk add --update git && rm -rf /var/cache/apk/*
RUN go get github.com/golang/dep/cmd/dep
RUN go get github.com/pilu/fresh

# Prepare dev build
ADD . .
RUN dep ensure --vendor-only
RUN go build -o app
CMD ["fresh"]

FROM alpine:latest as prod
WORKDIR /root
ADD . .
COPY --from=dev /go/src/chrisbrindley.co.uk/app .
CMD ["./app"]