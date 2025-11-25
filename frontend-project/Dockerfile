# Stage 1 : build Angular app
FROM docker.io/library/node:22-alpine3.22 AS build

WORKDIR /app

# Copy package.json / package-lock.json
COPY frontend/package*.json ./

# Install deps
RUN npm install

# Copy the rest of the Angular project
COPY frontend/. .

# Build Angular in production mode
RUN npm run build -- --configuration production

# Stage 2 : Nginx to serve the static files
FROM docker.io/library/nginx:1.29-alpine3.22

# Copy Angular built files (Angular 17+ outputs into dist/<name>/browser)
COPY --from=build /app/dist/*/browser /usr/share/nginx/html

EXPOSE 80

CMD ["nginx", "-g", "daemon off;"]
