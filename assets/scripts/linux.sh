#!/bin/bash

# Exit on errors
set -e

# Update package list and install common dependencies
echo "Installing dependencies for Linux..."
sudo apt-get update
sudo apt-get install -y git git-lfs make gnupg ca-certificates curl nvidia-container-toolkit

# Add Docker GPG key and repository
if [ ! -f /etc/apt/keyrings/docker.asc ]; then
  echo "Adding Docker GPG key and repository..."
  curl -fsSL https://download.docker.com/linux/ubuntu/gpg | sudo tee /etc/apt/keyrings/docker.asc > /dev/null
  echo "deb [arch=amd64 signed-by=/etc/apt/keyrings/docker.asc] https://download.docker.com/linux/ubuntu $(lsb_release -cs) stable" | sudo tee /etc/apt/sources.list.d/docker.list > /dev/null
fi

# Install Docker and Docker Compose Plugin
sudo apt-get update
sudo apt-get install -y docker-ce docker-ce-cli containerd.io docker-buildx-plugin docker-compose-plugin

# Setup Docker daemon for NVIDIA GPU support
sudo mkdir -p /etc/docker
sudo tee /etc/docker/daemon.json <<EOF
{
  "runtimes": {
    "nvidia": {
      "path": "/usr/bin/nvidia-container-runtime",
      "runtimeArgs": []
    }
  }
}
EOF

# Create model directories
echo "Creating model directories..."
mkdir -p assets/models/linesegmentation assets/models/regionsegmentation assets/models/textrecognition

# Clone Hugging Face models
echo "Installing Git LFS and cloning models..."
git lfs install
git lfs clone https://huggingface.co/Riksarkivet/yolov9-lines-within-regions-1 assets/models/linesegmentation/yolov9-lines-within-regions-1
git lfs clone https://huggingface.co/Sprakbanken/TrOCR-norhand-v3 assets/models/textrecognition/TrOCR-norhand-v3

# Create .env file
echo "Creating .env file..."
if [ ! -f .env ]; then
  cp example.env .env
else
  echo ".env file already exists. Skipping."
fi

# Add current user to Docker group (for Linux only)
echo "Adding current user to Docker group..."
sudo usermod -aG docker $USER
newgrp docker
sudo systemctl daemon-reload
sudo systemctl restart docker

echo "Setup complete!"
