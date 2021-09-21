dotnet publish -c Release -o published && \
docker build -t exo-webapp-dotnet:1.00 .