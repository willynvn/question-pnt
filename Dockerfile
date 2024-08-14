FROM golang:1.16-alpine

WORKDIR /app

COPY . .

RUN go build -o main .

FROM alpine:latest

COPY --from=builder /app/main /main

ENTRYPOINT ["/main"]

EXPOSE 8080
