# Build Stage
FROM golang:1.20-bullseye as build

RUN apt-get update -y && apt-get upgrade -y

WORKDIR /app

COPY go.mod /app/
COPY go.sum /app/

RUN go mod download
RUN go mod tidy

COPY . /app/

RUN make build

# Execute Stage
FROM ubuntu:22.04

RUN apt-get update -y && apt-get upgrade -y
RUN apt-get install -y ca-certificates

WORKDIR /app
COPY --from=build /app/etc /app/etc
COPY --from=build /app/build/main /app/main

EXPOSE 8080

CMD ["/app/main"]