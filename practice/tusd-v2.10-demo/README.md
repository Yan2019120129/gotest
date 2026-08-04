# tusd v2.10.0 Go client/server demo

This project embeds `github.com/tus/tusd/v2 v2.10.0` in a Go HTTP server and implements a small Go tus 1.0 client directly with `net/http`.

## Included features

### Server

- Local `filestore` plus `filelocker`.
- Tus creation, resumable upload, HEAD status, GET download, DELETE termination, and concatenation support exposed by tusd.
- Maximum upload size.
- Bearer-token middleware.
- Filename/extension metadata validation and server-owned metadata.
- Created/progress/completed/terminated notifications.
- CORS configuration, reverse-proxy forwarded headers, health endpoint, and graceful shutdown.

### Client

- `OPTIONS` capability discovery.
- Upload creation with `Upload-Length` and Base64 `Upload-Metadata`.
- Chunked PATCH uploads.
- Automatic resume based on server `HEAD` offset.
- Atomic JSON checkpoints.
- Retry with exponential delay and offset resynchronization after network failures.
- Upload status, cancellation, and download commands.
- SIGINT/SIGTERM cancellation.

## Requirements

`tusd v2.10.0` declares Go `1.25.8` in its module. With a recent Go installation and `GOTOOLCHAIN=auto`, Go can select/download the required toolchain automatically.

## Install dependencies

```bash
go mod download
```

## Run server

```bash
go run ./cmd/server \
  -addr :8080 \
  -upload-dir ./data/uploads \
  -token demo-token \
  -max-size 2147483648
```

Health check:

```bash
curl http://localhost:8080/healthz
```

Discover tus capabilities:

```bash
go run ./cmd/client options -endpoint http://localhost:8080/files/
```

## Upload

```bash
go run ./cmd/client upload \
  -file ./vendor.zip \
  -endpoint http://localhost:8080/files/ \
  -token demo-token \
  -chunk-size 8388608
```

Press `Ctrl+C` during upload. Run the same command again; the client loads `.tus-checkpoints`, sends `HEAD`, obtains the authoritative `Upload-Offset`, and continues from that byte.

The `resume` command is an alias:

```bash
go run ./cmd/client resume -file ./vendor.zip
```

Create a completely new upload instead of using the local checkpoint:

```bash
go run ./cmd/client upload -file ./vendor.zip -restart
```

## Inspect, download, and terminate

```bash
UPLOAD_URL='http://localhost:8080/files/<upload-id>'

go run ./cmd/client info -url "$UPLOAD_URL" -token demo-token

go run ./cmd/client download \
  -url "$UPLOAD_URL" \
  -out ./downloaded-vendor.zip \
  -token demo-token

go run ./cmd/client cancel -url "$UPLOAD_URL" -token demo-token
```

## Storage layout

The local tusd file store writes the raw upload using its generated upload ID and stores metadata in a sibling `<id>.info` JSON file. The original filename is metadata; it is not used directly as the storage path.

## Important implementation details

1. `POST /files/` creates an upload resource and returns `Location`.
2. `HEAD <Location>` returns the current `Upload-Offset`.
3. `PATCH <Location>` sends bytes with `Content-Type: application/offset+octet-stream` and the matching `Upload-Offset`.
4. On retry, the client sends another `HEAD` because a disconnected PATCH may have committed some bytes.
5. Completion removes the local checkpoint. The server retains the uploaded object until DELETE or an external cleanup policy removes it.

See `TODO.md` and inline `TODO(...)` comments for production work that is deliberately not implemented.
