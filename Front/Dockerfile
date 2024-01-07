FROM nginx:1.25.3-alpine
COPY nginx.conf /etc/nginx/nginx.conf
COPY ./frontend/dist/frontend /usr/share/nginx/html

#Commande un peu barbare qui permet l'utilisation de variables d'environnement
CMD ["/bin/sh",  "-c",  "envsubst < /usr/share/nginx/html/assets/env.template.js > /usr/share/nginx/html/assets/env.js && exec nginx -g 'daemon off;'"]