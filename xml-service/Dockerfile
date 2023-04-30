FROM golang:1.19-buster as builder

WORKDIR /xml
COPY go.* ./
RUN go mod download
COPY . .
RUN GOOS=linux GOARCH=amd64 CGO_ENABLED=0 go build -v -o application

FROM alpine:3.15.4
WORKDIR /xml
COPY --from=builder /xml/application /xml/application
COPY *.yaml ./
CMD ["/xml/application"]