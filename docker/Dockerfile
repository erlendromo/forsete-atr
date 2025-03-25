# Application
FROM golang:1.23.6-alpine3.21 AS builder

WORKDIR /app

# Copy Go module files and download dependencies
COPY go.mod go.sum ./
RUN go mod download

# Copy application source code
COPY . .

# Build Go binary
RUN CGO_ENABLED=0 GOOS=linux go build -a -ldflags '-extldflags "-static"' -o app .



# Final stage
FROM nvidia/cuda:12.8.1-base-ubuntu22.04 AS final

WORKDIR /

# Install system dependencies
RUN apt-get update && apt-get install -y \
    curl \
    git \
    gcc \
    libc-dev \
    libgl1 \
    libglib2.0-0 \
    python3 \
    python3-venv \
    bash \
    && rm -rf /var/lib/apt/lists/*

# Create venv
RUN python3 -m venv /htrflow/venv
ENV PATH="/htrflow/venv/bin:$PATH"

# Upgrade and install dependencies
COPY docker/requirements.txt .
RUN /htrflow/venv/bin/pip install --no-cache-dir --upgrade pip setuptools wheel && \
    /htrflow/venv/bin/pip install --no-cache-dir -r requirements.txt

# Copy from builder stage
COPY --from=builder /app/app /app
COPY --from=builder /app/scripts /scripts
COPY --from=builder /app/docs /docs

# Expose API port
EXPOSE ${API_PORT}

# Run application
CMD ["/app"]
