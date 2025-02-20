# Build stage
FROM golang:1.24-alpine AS builder
WORKDIR /app
# Add build tools and SQLite dependencies
RUN apk add --no-cache gcc musl-dev
COPY go.mod go.sum ./
RUN go mod download
COPY . .
# Enable CGO and build
ENV CGO_ENABLED=1
RUN go build -o main .

# Development stage
FROM golang:1.24-alpine AS development
WORKDIR /app
# Add build tools and SQLite dependencies
RUN apk add --no-cache gcc musl-dev sqlite-dev
# Enable CGO
ENV CGO_ENABLED=1
COPY go.mod go.sum ./
RUN go mod download
COPY . .
EXPOSE 3000
CMD ["go", "run", "main.go"]

# Production stage
FROM alpine:latest AS production
WORKDIR /app
# Add SQLite runtime dependencies
RUN apk add --no-cache sqlite-libs
COPY --from=builder /app/main .
EXPOSE 3000
CMD ["./main"] 