#!/bin/bash

echo 'Begin'
CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o bin/centmonit

CONTAINER=test_monit
docker cp ./bin $CONTAINER:/home/
docker cp html $CONTAINER:/home/

echo 'Done!'
