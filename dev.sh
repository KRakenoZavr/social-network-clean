#!/bin/bash
export BACKEND_PORT=3000
export MIGRATION=false
export DB_NAME=social.db
export RUN_ENV=dev

cd backend && echo $BACKEND_PORT && go run cmd/main.go
