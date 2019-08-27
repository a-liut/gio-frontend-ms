# gio-frontend-ms
Microservice that exposes a web interface for interacting with the Giò Plant platform.

## Run

### Docker

The simplest way to execute the server is to use Docker.
Follow these steps to run the container:

```bash
docker build -t gio-frontend-ms:latest .
docker run -it -p 8080:8080 gio-frontend-ms:latest
```

### Standalone

gio-api-frontend-ms is developed as a Go module.

```bash
go build -o frontend cmd/frontend/main.go

./frontend
```

## Endpoints

- /rooms: List all available rooms

- /rooms/{id}: Get information about a room

- /rooms/{id}/devices: List all available devices in a room

- /rooms/{id}/devices/{id}: Get information about a device in a room