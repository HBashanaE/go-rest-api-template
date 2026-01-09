# Stage 1: Modules caching
FROM golang:1.23-alpine as modules
COPY go.mod go.sum /modules/
WORKDIR /modules
RUN go mod download

# Stage 2: Builder
FROM golang:1.23-alpine as builder
COPY --from=modules /go/pkg /go/pkg
COPY . /app
WORKDIR /app
# Build the binary statically
RUN CGO_ENABLED=0 GOOS=linux go build -o /bin/api ./cmd/api

# Stage 3: Final (lean image)
FROM alpine:latest
COPY --from=builder /bin/api /api
COPY .env / # Optional: remove if using secret managers in prod

# Expose port and run
EXPOSE 8080
CMD ["/api"]