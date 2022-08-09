# -----------------------
# Build Application
# -----------------------
FROM golang:alpine as builder

RUN apk update && apk add --no-cache git make

RUN mkdir /source
WORKDIR /source

COPY go.mod .
COPY go.sum .

RUN go mod download
RUN go mod verify

COPY /config ./config
COPY /console ./console
COPY /app ./app
COPY .env .env

RUN ls -lah

RUN go build -o main-app ./console/main.go

RUN ls -lah

# -----------------------
# Setup Application Runner
# -----------------------
FROM alpine:latest

RUN mkdir /app
WORKDIR /app

RUN apk add --no-cache ca-certificates

COPY --from=builder /source/main-app .
COPY --from=builder /source/.env .

EXPOSE 8081

RUN ls -lah

CMD [ "./main-app" ]
