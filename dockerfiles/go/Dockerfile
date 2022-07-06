FROM golang:1.17.6-alpine3.14 as build
WORKDIR /build
COPY ./runner .
RUN go build main.go

FROM alpine:latest
COPY --from=build /build .
COPY ./dockerfiles/go/runner.config.json ./
RUN mkdir files && apk add --no-cache musl-dev go
CMD ["./main"]