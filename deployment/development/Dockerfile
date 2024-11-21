FROM golang:alpine AS builder

WORKDIR /build

RUN apk --no-cache add bash make
COPY ["./go.mod", "./go.sum", "./"]

RUN go mod download

COPY ../.. .

RUN make build

FROM alpine:latest AS runner

COPY --from=builder /build/bin/auth_service ./

CMD ["./auth_service"]

EXPOSE 50000/tcp
