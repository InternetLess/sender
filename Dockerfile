FROM golang:1.15 as build
WORKDIR /build
COPY go.mod .
RUN go mod download
COPY . .
RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -ldflags="-w -s" -o ./app *.go

FROM alpine:3.12

COPY --from=build /build/app /usr/local/bin/app
CMD ["/usr/local/bin/app"]
