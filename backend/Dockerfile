# Step 1: Build stage
FROM golang:1.16.3-alpine AS builder

RUN apk update && apk add --no-cache git

WORKDIR /build

COPY . .
RUN CGO_ENABLED=0 go build -o /build/app

# Step 2: Build a small image
FROM scratch

COPY --from=builder /build/app /

ENTRYPOINT ["/app"]
