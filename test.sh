#!/bin/bash
export BACKEND_PORT=3000
export MIGRATION=true
export DB_NAME=test.db
export RUN_ENV=test

cd backend && go run cmd/main.go
