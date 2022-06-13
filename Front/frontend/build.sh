#!/bin/bash

#####################
# Version Classique #
#####################
npm install
npm run build && \
docker build -t frontend:1.0 ../


#####################
# Version dynamique #
#####################
# npm install

# for i in {10..15}; do
#     TAG=stu$i
#     npm run build -- --base-href /"$TAG"/ && \
#     docker build -t fgtech/frontend-todolist:"$TAG" ../

#     docker push fgtech/frontend-todolist:"$TAG"
 
# done