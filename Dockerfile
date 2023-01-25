FROM golang:1.18 AS build

ENV CGO_ENABLED 0

WORKDIR /usr/src/app

COPY go.mod go.sum ./
RUN go mod download && go mod verify

COPY . .

RUN go build -a -installsuffix cgo -o ares .

FROM alpine:latest AS runtime

COPY --from=build /usr/src/app/ares /usr/local/bin/ares

CMD ["ares"]
