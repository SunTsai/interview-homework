#!/bin/bash

if [ -z "$1" ]; then
  echo "Please specify a problem number (1, 2 or 3)"
  exit 1
fi

case $1 in
  1)
    DIRECTORY="math_questions"
    ;;
  2)
    DIRECTORY="quorum_election"
    ;;
  *)
    echo "Invalid problem number. Use 1 for math_questions, 2 for quorum_election."
    exit 1
    ;;
esac

if [ ! -d "$DIRECTORY" ]; then
  echo "Directory $DIRECTORY does not exist"
  exit 1
fi

cd "$DIRECTORY" || exit

GO_FILE="./cmd/main.go"
if [ ! -f "$GO_FILE" ]; then
  echo "$GO_FILE does not exist"
  exit 1
fi

if [ "$1" = 2 ]; then
  go run "$GO_FILE" $2
else
  go run "$GO_FILE"
fi