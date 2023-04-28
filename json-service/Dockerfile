FROM golang:1.19-buster as builder

WORKDIR /json
COPY go.* ./
RUN go mod download
COPY . .
RUN GOOS=linux GOARCH=amd64 CGO_ENABLED=0 go build -v -o application

FROM alpine:3.15.4
WORKDIR /json
COPY --from=builder /json/application /json/application
COPY *.yaml ./
CMD ["/json/application"]