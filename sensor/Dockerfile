# builder image
FROM golang:latest as builder
ENV CGO_ENABLED=0
WORKDIR /app
COPY go.* cmd/sensor/main.go cmd/sensor/model.yaml ./
COPY internal/pkg internal/pkg
RUN go mod download
RUN go build -o .

# generate clean, final image for end users
FROM debian:latest
WORKDIR /app
COPY --from=builder /app/sensor /app/model.yaml /app/

# executable
ENTRYPOINT ["/app/sensor"]

