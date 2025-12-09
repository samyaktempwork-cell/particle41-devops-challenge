# SimpleTime Service

A lightweight Go application that returns the current time and IP address in JSON format.

## Features
- JSON response: `{"timestamp": "...", "ip": "..."}`
- **Secure**: Runs as a non-root user.
- **Lightweight**: Built on `scratch` (approx. 8MB).

## How to Run
```bash
docker run -p 8080:8080 yourusername/simpletime:latest
```

## Local Development
1. Build the image: ```docker build -t simpletime .```
2. Run locally: ```docker run -p 8080:8080 simpletime```