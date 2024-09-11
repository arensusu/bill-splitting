#!/bin/bash

git pull origin main
docker compose build
docker compose up -d
docker system prune -af
