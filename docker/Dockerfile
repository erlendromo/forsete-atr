FROM golang:1.23.6-alpine3.21 AS builder

WORKDIR /app

COPY go.mod go.sum ./
RUN go mod download

COPY . .

RUN CGO_ENABLED=0 GOOS=linux go build -a -ldflags '-extldflags "-static"' -o executable .





FROM nvidia/cuda:12.8.1-base-ubuntu22.04 AS final

WORKDIR /

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

RUN python3 -m venv /htrflow/venv
ENV PATH="/htrflow/venv/bin:$PATH"

COPY docker/requirements.txt .
RUN /htrflow/venv/bin/pip install --no-cache-dir --upgrade pip setuptools wheel && \
    /htrflow/venv/bin/pip install --no-cache-dir -r requirements.txt

COPY --from=builder /app/assets/migrations /assets/migrations
COPY --from=builder /app/assets/scripts /assets/scripts
COPY --from=builder /app/docs /docs
COPY --from=builder /app/executable /executable

EXPOSE ${API_PORT}

CMD ["/executable"]
