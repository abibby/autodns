FROM golang:1.20 as go-build

WORKDIR /go/src/github.com/abibby/autodns

COPY go.mod go.sum ./
RUN go mod download

COPY . .

RUN GOOS=linux GOARCH=amd64 CGO_ENABLED=0 go build

# Now copy it into our base image.
FROM alpine

RUN apk update && \
    apk add ca-certificates && \
    update-ca-certificates

COPY --from=go-build /go/src/github.com/abibby/autodns/autodns /autodns

CMD ["/autodns"]
