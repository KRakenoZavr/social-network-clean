#!/bin/bash

GIT_DIR=$(git rev-parse --git-dir)

echo "Installing hooks..."
# this command creates symlink to our pre-commit script
ln -s ../../scripts/pre-commit.sh $GIT_DIR/hooks/pre-commit

echo "Installing node modules..."
# install node modules
cd test && npm i && cd ../

echo "Installing go modules..."
# tidy modules
cd backend && go mod tidy

echo "Done"!
