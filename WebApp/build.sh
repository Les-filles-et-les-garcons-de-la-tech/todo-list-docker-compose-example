dotnet publish -c Release -o published && \
docker build -t yannisbourree/exo-webapp-dotnet:1.00 .
docker push yannisbourree/exo-webapp-dotnet:1.00