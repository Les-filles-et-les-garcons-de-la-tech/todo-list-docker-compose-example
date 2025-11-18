# Stage 1 : build Angular app
FROM node:22-alpine AS build

WORKDIR /app

# Copy package.json / package-lock.json
COPY frontend/package*.json ./

# Install deps
RUN npm ci

# Copy the rest of the Angular project
COPY frontend/. .

# Build Angular in production mode
RUN npm run build -- --configuration production

# Stage 2 : Nginx to serve the static files
FROM nginx:1.27-alpine

# Copy Angular built files (Angular 17+ outputs into dist/<name>/browser)
COPY --from=build /app/dist/*/browser /usr/share/nginx/html

EXPOSE 80

CMD ["nginx", "-g", "daemon off;"]
