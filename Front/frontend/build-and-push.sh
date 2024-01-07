#!/bin/bash

#####################
# Version Classique #
#####################
REPOSITORY=fgtech
IMAGE=
TAG=1.0
npm install
npm run build && \
docker build -t "$REPOSITORY"/"$IMAGE":"$TAG" ../
docker push "$REPOSITORY"/"$IMAGE":"$TAG"

##########################
# Version dynamique avec #
# base-href par étudiant #
##########################
# npm install

# for i in {1..9}; do
#     TAG=stu0$i
#     npm run build -- --base-href /"$TAG"/ && \
#     docker build -t "$REPOSITORY"/"$IMAGE":"$TAG" ../

#     docker push "$REPOSITORY"/"$IMAGE":"$TAG"
 
# done

# for i in {10..15}; do
#     TAG=stu$i
#     npm run build -- --base-href /"$TAG"/ && \
#     docker build -t "$REPOSITORY"/"$IMAGE":"$TAG" ../

#     docker push "$REPOSITORY"/"$IMAGE":"$TAG"
 
# done