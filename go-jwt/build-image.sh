#!/usr/bin/env sh
# build image
docker build --network host -t go-jwt-demo .
yes|docker image prune
