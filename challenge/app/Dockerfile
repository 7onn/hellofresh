FROM golang:1.14-alpine as builder
WORKDIR /app
COPY . .
RUN go install && \
    go build .

FROM golang:1.14-alpine
WORKDIR /app
COPY --from=builder /app/hellofresh ./hellofresh

CMD ["./hellofresh"]
