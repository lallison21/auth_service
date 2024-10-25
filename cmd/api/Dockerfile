FROM golang:alpine AS builder

WORKDIR /build

COPY ["./go.mod", "./go.sum", "./"]

RUN go mod download

COPY ../.. .

RUN go build -o ./bin/todo ./cmd/api/api.go && \
        chmod +x bin/todo

FROM alpine:latest AS runner

COPY --from=builder /build/bin/todo ./

CMD ["./todo"]

EXPOSE 60000/tcp
