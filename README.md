# gio-frontend-ms
Microservice that exposes a web User Interface (UI) allowing users to interact with the Giò Plant platform.

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

- /rooms: Lists all available rooms

- /rooms/{roomId}: Gets information about a room

- /rooms/{roomId}/devices: Lists all available devices in a room

- /rooms/{roomId}/devices/{deviceId}: Gets information about a device in a room