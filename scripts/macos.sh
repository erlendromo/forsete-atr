#!/bin/bash

# Exit on errors
set -e

echo "Installing dependencies for macOS..."

# Install Homebrew if missing
if ! command -v brew &> /dev/null; then
    echo "Installing Homebrew..."
    /bin/bash -c "$(curl -fsSL https://raw.githubusercontent.com/Homebrew/install/HEAD/install.sh)"
fi

# Install required packages
brew install git git-lfs make go

# Install Docker (if not installed)
if ! command -v docker &> /dev/null; then
    echo "Installing Docker..."
    brew install --cask docker
    brew install docker-compose
fi

# Start Docker Desktop
open -a Docker
echo "Waiting for Docker to start..."
sleep 15

# Enable Git LFS
git lfs install

# Create model directories
echo "Creating model directories..."
mkdir -p models/linesegmentation models/regionsegmentation models/textrecognition

# Clone Hugging Face models
echo "Cloning models..."
git lfs clone https://huggingface.co/Riksarkivet/yolov9-lines-within-regions-1 models/linesegmentation/yolov9-lines-within-regions-1
git lfs clone https://huggingface.co/Sprakbanken/TrOCR-norhand-v3 models/textrecognition/TrOCR-norhand-v3

# Create .env file
echo "Creating .env file..."
cp example.env .env

echo "Setup complete!"
