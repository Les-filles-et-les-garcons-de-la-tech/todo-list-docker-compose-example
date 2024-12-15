#!/bin/bash

# REPOSITORY=fgtech
IMAGE=backend
# TAG=2024.12

dotnet publish -c Release -o published && \
# docker build -t $REPOSITORY/$IMAGE:$TAG .
docker build -t $IMAGE .
# docker push $REPOSITORY/$IMAGE:$TAG