FROM golang:1.14 AS builder
ADD . /src
RUN cd /src && go build -o shortd /src/cmd/shortd/main.go

FROM ubuntu:latest
WORKDIR /app
COPY --from=builder /src/shortd /app/
ENTRYPOINT [ "./shortd" ]
