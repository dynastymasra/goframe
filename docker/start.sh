#!/bin/bash -e

case $1 in
  "build")
    go build -ldflags="-s -w" -tags=main -o goframe
    ;;
  "run")
    if [ ! -x "goframe" ]; then
      $0 build goframe
    fi
    shift && ./goframe "$@"
    ;;
  "test")
    ./test.sh
    ;;
  *)
    echo "usage: $0 [build|run|test]"
    exit 1
    ;;
esac