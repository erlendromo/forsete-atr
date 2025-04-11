#!/bin/bash

# Exit on error and catch errors in piped commands
set -eo pipefail

echo "Activating Virtual Environment..."
source /htrflow/venv/bin/activate || { echo "FAILED TO ACTIVATE VENV"; exit 1; }

echo "Checking HTRflow:"
which htrflow || { echo "HTRFLOW NOT FOUND"; exit 1; }

# Check if two arguments are provided
if [ "$#" -ne 2 ]; then
    echo "Usage: $0 <path-to-yaml-file> <path-to-image-file>"
    exit 1
fi

# Assign arguments to variables
YAML_FILE="$1"
IMAGE_FILE="$2"

# Check if the YAML file exists
if [ ! -f "$YAML_FILE" ]; then
    echo "Error: YAML file '$YAML_FILE' not found!"
    exit 1
fi

# Check if the image file exists
if [ ! -f "$IMAGE_FILE" ]; then
    echo "Error: Image file '$IMAGE_FILE' not found!"
    exit 1
fi

echo "Running htrflow with YAML File: $YAML_FILE and Image File: $IMAGE_FILE"
htrflow pipeline "$YAML_FILE" "$IMAGE_FILE" || { echo "HTRFlow execution failed!"; exit 1; }

exit 0
