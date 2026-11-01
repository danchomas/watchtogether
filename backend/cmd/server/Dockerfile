FROM golang:1.27-alpine AS builder

WORKDIR /src
COPY go.mod go.sum ./
RUN go mod download
COPY cmd ./cmd
COPY internal ./internal
RUN CGO_ENABLED=0 GOOS=linux go build -trimpath -ldflags="-s -w" -o /server ./cmd/server

FROM alpine:latest
RUN addgroup -S app && adduser -S -G app app
COPY --from=builder /server /server
USER app
EXPOSE 8080
CMD ["/server"]
