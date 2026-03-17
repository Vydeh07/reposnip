# Stage 1 — Build
FROM golang:1.26-alpine AS builder

WORKDIR /app

COPY go.mod .
RUN go mod download

COPY *.go .

RUN CGO_ENABLED=0 go build -o reposnip .

# Stage 2 — Run
FROM alpine:latest

WORKDIR /app

COPY --from=builder /app/reposnip .

ENTRYPOINT ["./reposnip"]
