#!/bin/bash

REPOSITORY=fgtech
IMAGE=backend-todolist
TAG=1.0

dotnet publish -c Release -o published && \
docker build -t $REPOSITORY/$IMAGE:$TAG .
docker push fgtech/$IMAGE:$TAG