dotnet publish -c Release -o published && \
docker build -t web-app . && \
docker image prune -f