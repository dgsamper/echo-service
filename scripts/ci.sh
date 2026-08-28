#!/usr/bin/env bash
set -euo pipefail

cd "$(dirname "${BASH_SOURCE[0]}")/.."

KIND_CLUSTER_NAME="${KIND_CLUSTER_NAME:-kind}"

go test -v ./internal/echo
go vet ./...

(
  cd infra
  go build -o /dev/null .
  go vet ./...
)

docker build -t echo-service:local .

if ! kind get clusters | grep -Fx -- "$KIND_CLUSTER_NAME" >/dev/null; then
  echo "Kind cluster '$KIND_CLUSTER_NAME' not found. Create it and rerun this script." >&2
  exit 1
fi

kind load docker-image echo-service:local --name "$KIND_CLUSTER_NAME"
