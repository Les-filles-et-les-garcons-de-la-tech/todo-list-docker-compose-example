FROM mcr.microsoft.com/dotnet/aspnet:5.0 AS runtime

WORKDIR /app
COPY published ./
EXPOSE  5001
ENTRYPOINT ["dotnet", "WebApp.dll"]