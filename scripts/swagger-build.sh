#! /bin/sh

docker run --rm -u "$(id -u):$(id -g)" -v $(pwd):/code ghcr.io/swaggo/swag:latest init
