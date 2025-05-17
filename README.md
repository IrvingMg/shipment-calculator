# Shipment Calculator

A simple web application to calculate optimal shipment pack combinations.

## Getting Started

### Prerequisites

- [Go 1.24](https://golang.org/doc/install)
- [Docker](https://www.docker.com/)
- [Docker Compose](https://docs.docker.com/compose/)

## How to Run

To run the full application (both server and UI), the easiest method is via Docker Compose:

```bash
make docker-run
```

Once running:

- The UI will be available at: [http://localhost:8080](http://localhost:8080)
- The API (server) will be available at: [http://localhost:8081](http://localhost:8081)

### Running the Server Only

If you only need to run the server, you have a few options:

> The server will run on **port 8080 by default**.

1. **Run from source:**

```bash
make go-run-server
```

2. **Build and run the compiled binary:**

```bash
make run-server
```

3. **Run in Docker:**

```bash
make docker-run-server
```

## API

The backend exposes a single endpoint:

**POST** `/shipment-packs`

### Request

```json
{
  "pack_sizes": [250, 500, 1000],
  "amount": 1500
}
```

### Response

```json
{
  "packs": {
    "1000": 1,
    "500": 1
  }
}
```

### Example using `curl`

```bash
curl -X POST http://localhost:8080/shipment-packs \
  -H "Content-Type: application/json" \
  -d '{"pack_sizes": [250, 500, 1000], "amount": 1500}'
```
