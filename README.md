# AOZUMI Go basic fixture

A small public Go web app for exercising AOZUMI repository detection, build, health, and HTTP routing.

```sh
go test ./...
go run .
```

The app reads `PORT` (default `8080`) and `NAME` (default `World`), serves `GET /`, and reports `GET /healthz` with HTTP 204. Unknown routes return 404.
