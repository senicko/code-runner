FROM golang:latest AS builder
WORKDIR /build
COPY . .
RUN CGO_ENABLED=0 GOOS=linux go build main.go

FROM golang:latest
WORKDIR /app
COPY --from=builder /build .
RUN mkdir files
CMD ["./main"]