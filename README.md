# echo

A small Go HTTP service that returns the request's `headers`, query `params`,
raw `body`, and `path` as JSON on any route.

See [SETUP.md](SETUP.md) to build and deploy it locally with Docker, Kind, and
Pulumi.

## Run locally

Requires Go 1.27.0 or newer. From the repository root:

```sh
go run ./cmd/echo
```

The default port is `8080`; set `PORT` to override it. In another terminal:

```sh
curl -sS 'http://localhost:8080/hello?name=world'
```

## Tests

```sh
go test ./...
go vet ./...
```

## Project layout

```text
cmd/echo/               Application entrypoint
internal/echo/          Handler, configuration, and unit tests
infra/                  Pulumi infrastructure (separate Go module)
scripts/ci.sh           Test, build, and load echo-service:local into Kind
.github/workflows/      GitHub Actions workflow using the same CI script
Dockerfile              Container build and runtime
```

CI runs on every push. Deployment with Pulumi is a separate step.
This service is intended for local use: it has no authentication, TLS, or request
body size limit.
