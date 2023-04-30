FROM golang:1.19-buster as builder

WORKDIR /avro
COPY go.* ./
RUN go mod download
COPY . .
RUN GOOS=linux GOARCH=amd64 CGO_ENABLED=0 go build -v -o application

FROM alpine:3.15.4
WORKDIR /avro
COPY --from=builder /avro/application /avro/application
COPY *.yaml ./
CMD ["/avro/application"]