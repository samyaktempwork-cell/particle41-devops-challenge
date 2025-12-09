FROM golang:1.23-alpine AS builder

WORKDIR /app

# Create a non-root user named 'appuser'
RUN adduser -D -g '' appuser

COPY go.mod ./

RUN go mod download

COPY . .

RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -ldflags="-w -s" -o simpletime main.go

FROM scratch

COPY --from=builder /etc/passwd /etc/passwd
COPY --from=builder /app/simpletime /simpletime

# Use the non-root user
USER appuser

ENTRYPOINT ["/simpletime"]