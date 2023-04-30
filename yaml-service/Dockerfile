FROM golang:1.19-buster as builder

WORKDIR /yaml
COPY go.* ./
RUN go mod download
COPY . .
RUN GOOS=linux GOARCH=amd64 CGO_ENABLED=0 go build -v -o application

FROM alpine:3.15.4
WORKDIR /yaml
COPY --from=builder /yaml/application /yaml/application
COPY *.yaml ./
CMD ["/yaml/application"]