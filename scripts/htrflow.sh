#!/bin/bash

# Ensure we are in the correct virtual environment
source /htrflow/venv/bin/activate

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
htrflow pipeline "$YAML_FILE" "$IMAGE_FILE"

exit 0
