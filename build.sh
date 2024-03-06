#!/bin/sh

IMAGE_NAME=vspeach/otus-k8s-deploy
IMAGE_VERSION=2.0

docker build -t $IMAGE_NAME:$IMAGE_VERSION ./app/
docker push $IMAGE_NAME:$IMAGE_VERSION
