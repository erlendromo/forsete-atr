# Base image with Python 3.10
FROM python:3.10-slim AS base

# Install system dependencies
RUN apt-get update && apt-get install -y \
    curl \
    git \
    gcc \
    libc-dev \
    libgl1 \
    libglib2.0-0 \
    && rm -rf /var/lib/apt/lists/*

# Create and activate Python virtual environment
RUN python3 -m venv /htrflow/venv
ENV PATH="/htrflow/venv/bin:$PATH"

# Upgrade pip and install dependencies
RUN /htrflow/venv/bin/pip install --no-cache-dir --upgrade pip setuptools wheel

# Copy and install Python requirements
COPY requirements.txt .
RUN /htrflow/venv/bin/pip install --no-cache-dir -r requirements.txt

# Verify htrflow installation
RUN /htrflow/venv/bin/htrflow --help || echo "HTRflow is not found"

# Set working directory for Go application
WORKDIR /app

# Install Go and build the application
FROM golang:1.23.6-alpine3.21 AS builder

WORKDIR /app

# Copy Go module files and download dependencies
COPY go.mod go.sum ./
RUN go mod download

# Copy application source code
COPY . .

# Build Go binary
RUN CGO_ENABLED=0 GOOS=linux go build -a -ldflags '-extldflags "-static"' -o app .

# Final stage: Merge everything into a single container
FROM base

WORKDIR /

# Copy the Go application from the builder stage
COPY --from=builder /app/app /app
COPY --from=builder /app/scripts /scripts

# Make sure bash is installed
RUN apt-get update && apt-get install -y bash

# Expose API port
EXPOSE 8080

# Run the Go application (while `htrflow` is available in the environment)
CMD ["/app/app"]
