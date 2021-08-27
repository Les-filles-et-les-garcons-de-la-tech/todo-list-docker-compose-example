docker compose down && \
npm run build && \
docker build -t frontend .. && \
docker image prune -f && \
docker compose up -d;