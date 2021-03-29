FROM golang:1.14 AS builder
ADD . /src
RUN cd /src && go build -o shorty /src/cmd/shorty/main.go

FROM ubuntu:latest
WORKDIR /app
COPY --from=builder /src/shorty /app/
ENTRYPOINT [ "./shorty" ]
