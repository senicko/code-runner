# Build 
FROM golang:1.17.6-alpine3.14 as build
WORKDIR /build
COPY . .
RUN go build main.go

# Cpp sandbox
FROM alpine:latest
COPY --from=build /build .
RUN mkdir files && apk add --no-cache g++ musl-dev
CMD ["./main"]