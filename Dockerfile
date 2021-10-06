# Builder
FROM golang:1.17-alpine AS builder
RUN mkdir /app
ADD . /app
WORKDIR /app
RUN go mod tidy

# Runner
RUN go build -o main
FROM alpine:3.14
WORKDIR /app
COPY --from=builder /app/main .
EXPOSE 8080
CMD ["./main"]