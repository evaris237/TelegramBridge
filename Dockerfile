FROM golang:1.13 as builder

WORKDIR /app

COPY . /app

RUN go get github.com/bwmarrin/discordgo

RUN CGO_ENABLED=0 go build -v -o DiscordBridge .

FROM alpine:latest

COPY --from=builder /app/DiscordBridge /DiscordBridge

ENTRYPOINT ["/DiscordBridge"]


