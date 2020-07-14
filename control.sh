#!/bin/bash

RED='\033[0;31m'
NC='\033[0m' # No Color

start() {
  printf "${RED}Starting CentMonit...${NC}\n"
  nohup ./bin/centmonit >/dev/null 2>&1 &
}

stop() {
  printf "${RED}Shutdown CentMonit...${NC}\n"
  kill -9 `ps -e -o pid,command | grep 'cent' | grep 'monit' | awk '{print $1}'`
}

status() {
  printf "${RED}Checking CentMonit...${NC}\n"
  ps -e -o pid,command | grep 'cent' | grep 'monit'
}

case "$1" in
  start)
    start
    ;;
  stop)
    stop
    ;;
  status)
    status
    ;;
  *)
    echo "--------------------------------------------------------"
    echo "Usage: ./control.sh <start> | <stop> | <status>"
    echo "--------------------------------------------------------"
    exit 1
    ;;
esac
exit $?
