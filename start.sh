#!/bin/bash
docker build -t weather-app --no-cache --build-arg APP_NAME=weather-app .
#docker-compose up