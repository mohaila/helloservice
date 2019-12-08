#!/bin/sh
echo "Deploying hellosv with Docker Compose"
docker-compose -f deployments/compose/docker-compose.yml up -d