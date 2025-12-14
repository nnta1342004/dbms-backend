FROM golang:latest as builder

RUN mkdir /app

ADD . /app/
WORKDIR /app
RUN go mod tidy
RUN CGO_ENABLED=0 GOOS=linux go build -o app

FROM alpine:latest
WORKDIR /app/
COPY --from=builder /app .
CMD ["/app/app"]
