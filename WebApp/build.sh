dotnet publish -c Release -o published && \
docker build -t webapp:1.0 .