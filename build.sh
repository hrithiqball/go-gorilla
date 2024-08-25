#!/bin/bash

if [ ! -f version.txt ]; then
  echo "Error: version.txt file not found!"
  exit 1
fi

VERSION=$(cat version.txt | tr -d '[:space:]')

# Check if version is empty
if [ -z "$VERSION" ]; then
  echo "Error: version.txt is empty!"
  exit 1
fi

IMAGE_NAME="go-local-my"

docker build -t "$IMAGE_NAME:$VERSION" -f Dockerfile .

if [ $? -eq 0 ]; then
  echo "Successfully built Docker image: $IMAGE_NAME:$VERSION"
else
  echo "Error: Failed to build Docker image."
  exit 1
fi