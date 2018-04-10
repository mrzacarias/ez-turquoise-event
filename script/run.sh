#!/bin/sh

set -e

cd "$(dirname "$0")/.."

script/setup_and_build.sh

bin/ez-turquoise-event $*
