#!/bin/sh

set -e

cd "$(dirname "$0")/.."

echo "==> Setup app..."

# Nothing to do right now

echo "==> Building app..."
go build -ldflags="-s -w" -o bin/ez-turquoise-event main.go

echo "==> Done!"
