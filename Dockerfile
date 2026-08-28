FROM golang:1.27-alpine AS build

WORKDIR /src

COPY go.mod ./
COPY cmd/ ./cmd/
COPY internal/ ./internal/

RUN CGO_ENABLED=0 go build -o /echo ./cmd/echo

FROM scratch

COPY --from=build /echo /echo

USER 65532:65532

EXPOSE 8080

ENTRYPOINT ["/echo"]