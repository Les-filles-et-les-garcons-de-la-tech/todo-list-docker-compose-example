npm install
npm run build && \
docker build -t yannisbourree/exo-front-angular:1.00 ../ && \
docker push yannisbourree/exo-front-angular:1.00