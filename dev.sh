#!/bin/bash

RED='\033[0;31m'
NC='\033[0m' # No Color

build() {
  printf "${RED}Building binary...${NC}\n"
  CGO_ENABLED=0 GOOS=linux go build -o bin/centmonit-linux
  CGO_ENABLED=0 GOOS=solaris go build -o bin/centmonit-solaris
  CGO_ENABLED=0 GOOS=openbsd go build -o bin/centmonit-openbsd
  CGO_ENABLED=0 GOOS=netbsd go build -o bin/centmonit-netbsd
  CGO_ENABLED=0 GOOS=freebsd go build -o bin/centmonit-freebsd
}

docker_deployment() {
  printf "${RED}Deploy to docker test...${NC}\n"
  CONTAINER=test_monit
  docker cp bin $CONTAINER:/home/CentMonit/
  docker cp html $CONTAINER:/home/CentMonit/
  docker cp template $CONTAINER:/home/CentMonit/
  docker cp config.yml $CONTAINER:/home/CentMonit/
  docker cp control.sh $CONTAINER:/home/CentMonit/
}

make_release() {
  printf "${RED}Make release...${NC}\n"
  VERSION=$(cat ./VERSION)
  mkdir -p releases
  tar cvfj releases/CentMonit-$VERSION.tar.bz2 -C $PWD/../ \
    CentMonit/bin \
    CentMonit/html \
    CentMonit/template \
    CentMonit/config.yml \
    CentMonit/control.sh \
    CentMonit/MANUAL.pdf
}

case "$1" in
  build)
    build
    ;;
  docker)
    docker_deployment
    ;;
  release)
    make_release
    ;;
  *)
    echo "--------------------------------------------------------"
    echo "Usage: ./dev.sh <build> | <docker> | <release>"
    echo "--------------------------------------------------------"
    exit 1
    ;;
esac
exit $?
