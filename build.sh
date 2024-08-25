#!/bin/bash

VERSION=$(cat version.txt | tr -d '[:space:]')

if [ -z "$VERSION" ]; then
  echo "Error: version.txt is empty!"
  exit 1
fi

if [ ! -f .env.local ]; then
  echo "Error: .env.local file not found!"
  exit 1
fi


IMAGE_NAME="local/go-local-my"

docker build -t "$IMAGE_NAME:$VERSION" -f Dockerfile .

if [ $? -eq 0 ]; then
  echo "Successfully built Docker image: $IMAGE_NAME:$VERSION"
else
  echo "Error: Failed to build Docker image."
  exit 1
fi
