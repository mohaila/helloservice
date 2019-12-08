#!/bin/sh
echo "Building hellosvc Docker image"
if [[ -z "$HELLOSVC_VERSION" ]]; then
    VERSION="0.1"
else
    VERSION="$HELLOSVC_VERSION"
fi
docker build -t hellosvc:$VERSION  -f deployments/docker/Dockerfile .