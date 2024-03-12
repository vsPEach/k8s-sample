#!/bin/sh

IMAGE_NAME=vspeach/otus-k8s-deploy
IMAGE_VERSION=latest

docker build -t $IMAGE_NAME:$IMAGE_VERSION ./app/
docker push $IMAGE_NAME:$IMAGE_VERSION
