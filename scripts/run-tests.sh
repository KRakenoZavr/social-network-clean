#!/bin/bash

export BACKEND_PORT=3003
export MIGRATION=true
export DB_NAME=test.db
export RUN_ENV=test

(trap 'kill 0' SIGINT; 
cd backend && go run cmd/main.go & 
cd test && npm i && npm run test && kill -9 `lsof -t -i:3003` && rm ../backend/test.db)
