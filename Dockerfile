# builder
FROM golang:1.23 AS builder

WORKDIR /build

COPY . .

RUN --mount=type=cache,target=/go/pkg/mod \
    CGO_ENABLED=0 GOOS=linux go build -ldflags "-w -s" -a -installsuffix cgo -o /app/ip ./cmd/ip


# runtime
FROM scratch

COPY --from=builder /app/ip /app/ip

EXPOSE 80

CMD ["/app/ip"]
