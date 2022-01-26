FROM golang:1.17-alpine AS builder
WORKDIR /app

COPY . ./
RUN go mod download
RUN go build -o /iqs-api

FROM alpine:latest AS final
WORKDIR /app

COPY --from=builder /iqs-api /

ENV GIN_MODE=release
ENV PORT=8080
EXPOSE 8080
ENTRYPOINT ["/iqs-api"]