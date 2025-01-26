ARG GOLANG_VERSION=1.21-bookworm
ARG DEBIAN_VERSION=bookworm-slim

FROM golang:$GOLANG_VERSION as builder

WORKDIR /app

COPY go.* ./
RUN go mod download

COPY . ./

RUN go build -v -o server

FROM debian:$DEBIAN_VERSION

RUN set -x && apt update && DEBIAN_FRONTEND=noninteractive apt-get install -y \
    ca-certificates && \
    rm -rf /var/lib/apt/lists/*

COPY --from=builder /app/server /app/server

EXPOSE 3000

CMD ["/app/server"]