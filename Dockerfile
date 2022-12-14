# -----------------------
# Build Application
# -----------------------
FROM golang:1.17-alpine as builder

RUN apk update && apk add --no-cache git make

RUN mkdir /source
WORKDIR /source

COPY go.mod .
COPY go.sum .

RUN go mod download
RUN go mod verify

COPY /config ./config
COPY /console ./console
COPY /docs ./docs
COPY /routes ./routes
COPY /app ./app
COPY .env .env

RUN go build -o main-app ./console/main.go

# -----------------------
# Setup Application Runner
# -----------------------
FROM alpine:latest

RUN mkdir /app
WORKDIR /app

RUN apk add --no-cache ca-certificates

COPY --from=builder /source/main-app .
COPY --from=builder /source/docs ./docs
COPY --from=builder /source/.env .

CMD [ "./main-app" ]
