FROM golang:alpine AS builder

WORKDIR /frontend

# Install git for fetching dependencies
RUN apk update && apk add --no-cache git

COPY go.mod .

RUN go mod download

COPY . .

# Build the binary.
RUN go build -o /go/bin/frontend cmd/frontend/main.go

## Build lighter image
FROM alpine:latest
LABEL Name=gio-frontend-ms Version=1.0.0

# Copy our static executable.
COPY --from=builder /go/bin/frontend /frontend

EXPOSE 8080

# Run the binary.
ENTRYPOINT /frontend