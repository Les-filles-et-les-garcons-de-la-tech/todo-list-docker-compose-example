#!/usr/bin/env bash
set -euo pipefail

# Nom et tag de l'image (paramètres optionnels)
IMAGE_NAME="${1:-todo-backend}"
IMAGE_TAG="${2:-latest}"

echo "Building Golang backend Docker image: ${IMAGE_NAME}:${IMAGE_TAG}"
echo "Build context: $(pwd)"

docker build \
  -t "${IMAGE_NAME}:${IMAGE_TAG}" \
  .

# # Si tu veux pousser automatiquement sur un registre (optionnel)
# if [[ "${3:-}" == "push" ]]; then
#   echo "Pushing image ${IMAGE_NAME}:${IMAGE_TAG}"
#   docker push "${IMAGE_NAME}:${IMAGE_TAG}"
# fi

echo "Build finished: ${IMAGE_NAME}:${IMAGE_TAG}"
