FROM golang:alpine3.19 AS builder
WORKDIR /build
ADD go.mod .
COPY . .
RUN go build -o stranger-bot main.go


FROM alpine:3.19
WORKDIR /build
COPY --from=builder /build/stranger-bot /build/stranger-bot
CMD ["/build/stranger-bot"]