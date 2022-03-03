dotnet publish -c Release -o published && \
docker build -t h3rv3/backend-todolist:1.0 .